package auth

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/AaronSaikovski/gogoodwe/internal/shared/utils"
)

const (
	AuthLoginURL = "https://www.semsportal.com/api/v2/Common/CrossLogin"
)

var httpClient = utils.SharedHTTPClient

// SemsLogin is a method on the SemsLoginCredentials struct that performs a Sems login.
//
// It takes a context for cancellation and returns a pointer to a SemsLoginResponse struct and an error.
func (loginCredentials *SemsLoginCredentials) SemsLogin(ctx context.Context) (*SemsLoginResponse, error) {
	// Check if the UserLogin struct is empty
	if err := checkUserLoginInfo(loginCredentials); err != nil {
		return nil, err
	}

	// Convert User login struct to JSON
	loginData, err := utils.MarshalStructToJSON(loginCredentials)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal login credentials: %w", err)
	}

	// Create a new http request with pre-sized buffer
	req, err := http.NewRequest(http.MethodPost, AuthLoginURL, bytes.NewReader(loginData))
	if err != nil {
		return nil, fmt.Errorf("failed to create login request: %w", err)
	}

	req = req.WithContext(ctx)

	// Add headers
	setHeaders(req)

	// Make the API call with reusable client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("login request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("login request returned HTTP %d: %s", resp.StatusCode, body)
	}

	// Get the response body
	respBody, err := utils.FetchResponseBody(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read login response: %w", err)
	}

	// Unmarshal response to loginresponse struct
	var loginApiResponse SemsLoginResponse
	if err := utils.UnmarshalDataToStruct(respBody, &loginApiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal login response: %w", err)
	}

	// Check for successful login
	if err := checkUserLoginResponse(loginApiResponse.Msg); err != nil {
		return nil, err
	}

	return &loginApiResponse, nil
}
