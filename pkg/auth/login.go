package auth

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

const (
	AuthLoginURL = "https://www.semsportal.com/api/v2/Common/CrossLogin"
	HTTPTimeout  = 20 // seconds
)

// SemsLogin is a method on the SemsLoginCredentials struct that performs a Sems login.
//
// It takes no parameters and returns a pointer to a SemsLoginResponse struct and an error.
func (loginCredentials *SemsLoginCredentials) SemsLogin() (*SemsLoginResponse, error) {
	// Check if the UserLogin struct is empty
	if err := checkUserLoginInfo(loginCredentials); err != nil {
		return nil, err
	}

	// Convert User login struct to JSON
	loginData, err := utils.MarshalStructToJSON(loginCredentials)
	if err != nil {
		return nil, err
	}

	// Create a new http request
	req, err := http.NewRequest(http.MethodPost, AuthLoginURL, bytes.NewBuffer(loginData))
	if err != nil {
		return nil, err
	}

	// Add headers
	setHeaders(req)

	// Make the API call
	client := &http.Client{Timeout: HTTPTimeout * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Get the response body
	respBody, err := utils.FetchResponseBody(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal response to loginresponse struct
	var loginApiResponse SemsLoginResponse
	if err := utils.UnmarshalDataToStruct(respBody, &loginApiResponse); err != nil {
		return nil, err
	}

	// Check for successful login
	if err := checkUserLoginResponse(loginApiResponse.Msg); err != nil {
		return nil, err
	}

	return &loginApiResponse, nil
}
