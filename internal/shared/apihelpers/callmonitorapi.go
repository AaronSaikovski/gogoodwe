package apihelpers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/AaronSaikovski/gogoodwe/internal/shared/auth"
	"github.com/AaronSaikovski/gogoodwe/internal/shared/utils"
)

var httpClient = utils.SharedHTTPClient

// validateAPIURL checks that a server-provided API URL uses HTTPS and points to an allowed domain.
func validateAPIURL(rawURL string) error {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("invalid API URL: %w", err)
	}
	if parsed.Scheme != "https" {
		return fmt.Errorf("API URL must use HTTPS, got: %s", parsed.Scheme)
	}
	host := strings.ToLower(parsed.Hostname())
	if !strings.HasSuffix(host, ".semsportal.com") && host != "semsportal.com" {
		return fmt.Errorf("API URL host not in allowed domain: %s", host)
	}
	return nil
}

// FetchMonitorAPIData fetches data from the Monitor API.
//
// It takes in the context, authentication information, the URL of the power station,
// the HTTP timeout, and a pointer to a struct to store the output.
// It returns the raw JSON bytes and an error if there was a problem with the API call.
func FetchMonitorAPIData(ctx context.Context, authLoginInfo *auth.LoginInfo, powerStationURL string, inverterOutput any) ([]byte, error) {
	// Validate input parameters
	if authLoginInfo == nil || authLoginInfo.SemsLoginResponse == nil || authLoginInfo.SemsLoginCredentials == nil {
		return nil, fmt.Errorf("invalid authentication information")
	}
	if powerStationURL == "" {
		return nil, fmt.Errorf("powerStationURL cannot be empty")
	}

	// Get the Token header data
	apiResponseJSONData, err := auth.DataTokenJSON(authLoginInfo.SemsLoginResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to create token JSON: %w", err)
	}

	// Get the Powerstation ID header data
	powerStationIDJSONData, err := auth.PowerStationIdJSON(authLoginInfo.SemsLoginCredentials)
	if err != nil {
		return nil, fmt.Errorf("failed to create powerstation ID JSON: %w", err)
	}

	// Validate the API base URL from the login response to prevent SSRF
	apiBaseURL := authLoginInfo.SemsLoginResponse.API
	if err := validateAPIURL(apiBaseURL); err != nil {
		return nil, err
	}

	// Create URL from the Auth API and append the data URL part
	url := apiBaseURL + powerStationURL

	// Create a new HTTP request with pre-sized buffer
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(powerStationIDJSONData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req = req.WithContext(ctx)

	// Add headers
	auth.SetHeaders(req, apiResponseJSONData)

	// Make the API call with reusable client
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
	if err := utils.UnmarshalDataToStruct(respBody, inverterOutput); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Return the raw JSON bytes to avoid remarshaling
	return respBody, nil
}
