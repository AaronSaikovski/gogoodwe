package apihelpers

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
	"github.com/AaronSaikovski/gogoodwe/pkg/utils"
)

var (
	// Reusable HTTP client for better performance - no timeout set here
	httpClient = &http.Client{
		Transport: utils.NewHTTPTransport(),
	}
)

// FetchMonitorAPIData fetches data from the Monitor API.
//
// It takes in the context, authentication information, the URL of the power station,
// the HTTP timeout, and a pointer to a struct to store the output.
// It returns the raw JSON bytes and an error if there was a problem with the API call.
func FetchMonitorAPIData(ctx context.Context, authLoginInfo *auth.LoginInfo, powerStationURL string, HTTPTimeout int, inverterOutput interface{}) ([]byte, error) {
	// Validate input parameters
	if authLoginInfo == nil || authLoginInfo.SemsLoginResponse == nil || authLoginInfo.SemsLoginCredentials == nil {
		return nil, fmt.Errorf("invalid authentication information")
	}
	if powerStationURL == "" {
		return nil, fmt.Errorf("powerStationURL cannot be empty")
	}

	// Get the Token header data
	apiResponseJSONData, err := DataTokenJSON(authLoginInfo.SemsLoginResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to create token JSON: %w", err)
	}

	// Get the Powerstation ID header data
	powerStationIDJSONData, err := PowerStationIdJSON(authLoginInfo.SemsLoginCredentials)
	if err != nil {
		return nil, fmt.Errorf("failed to create powerstation ID JSON: %w", err)
	}

	// Create URL from the Auth API and append the data URL part (simple concatenation is faster for 2 strings)
	url := authLoginInfo.SemsLoginResponse.API + powerStationURL

	// Create a new HTTP request with pre-sized buffer
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(powerStationIDJSONData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Add context with timeout for thread-safe timeout handling
	ctx, cancel := context.WithTimeout(ctx, time.Duration(HTTPTimeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	// Add headers
	SetHeaders(req, apiResponseJSONData)

	// Make the API call with reusable client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	// Get the response body
	respBody, err := utils.FetchResponseBody(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal response to struct pointer
	if err := utils.UnmarshalDataToStruct(respBody, inverterOutput); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Return the raw JSON bytes to avoid remarshaling
	return respBody, nil
}
