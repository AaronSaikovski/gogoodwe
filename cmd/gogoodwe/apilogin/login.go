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

package apilogin

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

// Login -  Login to the SEMS API passing in a LoginCredentials struct and returning a LoginResponse struct.
func (loginCredentials *ApiLoginCredentials) APILogin() (*ApiLoginResponse, error) {

	// API Response struct
	loginApiResponse := ApiLoginResponse{}

	//check if the UserLogin struct is empty
	if err := checkUserLoginInfo(loginCredentials); err != nil {
		return nil, err
	}

	// User login struct to be converted to JSON
	loginData, err := utils.MarshalStructToJSON(loginCredentials)
	if err != nil {
		return nil, err
	}

	// Create a new http request
	req, err := http.NewRequest(http.MethodPost, AuthLoginURL, bytes.NewBuffer(loginData))
	if err != nil {
		return nil, err
	}

	//Add headers pass in the pointer to set the headers on the request object
	setHeaders(req)

	//make the API Call
	client := &http.Client{Timeout: time.Duration(HTTPTimeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Get the response body
	respBody, respErr := utils.FetchResponseBody(resp.Body)
	if respErr != nil {
		return nil, respErr
	}

	//marshall response to loginresponse struct
	dataErr := utils.UnmarshalDataToStruct(respBody, &loginApiResponse)
	if dataErr != nil {
		return nil, dataErr
	}

	// check for successful login return value..return a login error
	loginErr := checkUserLoginResponse(loginApiResponse.Msg)
	if loginErr != nil {
		return nil, loginErr
	}

	return &loginApiResponse, nil
}
