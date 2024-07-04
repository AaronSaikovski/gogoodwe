/*
MIT License

# Copyright (c) 2024 Aaron Saikovski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
