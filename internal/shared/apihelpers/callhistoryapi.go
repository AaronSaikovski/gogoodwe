package apihelpers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/AaronSaikovski/gogoodwe/internal/shared/auth"
	"github.com/AaronSaikovski/gogoodwe/internal/shared/utils"
)

const (
	historyDataPath = "HistoryData/ExportExcelStationHistoryData"
)

// FetchHistoryExportData fetches historical data from the ExportExcelStationHistoryData API.
//
// Uses the login response API base URL with the history data path.
func FetchHistoryExportData(ctx context.Context, authLoginInfo *auth.LoginInfo, requestBody []byte, output any) ([]byte, error) {
	if authLoginInfo == nil || authLoginInfo.SemsLoginResponse == nil {
		return nil, fmt.Errorf("invalid authentication information")
	}

	// Get the Token header data
	apiResponseJSONData, err := auth.DataTokenJSON(authLoginInfo.SemsLoginResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to create token JSON: %w", err)
	}

	// Build URL from login response API base
	apiURL := authLoginInfo.SemsLoginResponse.API + historyDataPath

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req = req.WithContext(ctx)

	// Add headers
	auth.SetHeaders(req, apiResponseJSONData)
	req.Header.Set("Accept", "application/json, text/plain, */*")

	// Make the API call
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, body)
	}

	// Get the response body
	respBody, err := utils.FetchResponseBody(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal response to struct pointer
	if err := utils.UnmarshalDataToStruct(respBody, output); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return respBody, nil
}
