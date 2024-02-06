/*
# Name: authentication - authenticates to the goodwe API - https://www.semsportal.com/api/v2/Common/CrossLogin
# Author: Aaron Saikovski - asaikovski@outlook.com
*/

package authentication

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

var (
	// Auth Login Url
	AuthLoginUrl string = "https://www.semsportal.com/api/v2/Common/CrossLogin"

	// Default timeout value
	HTTPTimeout int = 20

	//API login success response message
	SemsLoginSuccessResponse string = "Successful"
)

// DoLogin - Main public login function
// Logs into the SEMs API
func DoLogin(SemsResponseData *types.LoginResponse, UserLogin *types.LoginCredentials) error {

	//check if the UserLogin struct is empty
	if usererr := checkUserLoginInfo(UserLogin); usererr != nil {
		return usererr
	}

	// User login struct to be converted to JSON
	jsonData, jsonErr := utils.MarshalStructToJSON(UserLogin)
	if jsonErr != nil {
		return jsonErr
	}

	// Create a new http request
	req, err := http.NewRequest(http.MethodPost, AuthLoginUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	//Add headers pass in the pointer to set the headers on the request object
	setHeaders(req)

	//make the API Call
	client := &http.Client{Timeout: time.Duration(HTTPTimeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	//cleanup
	defer resp.Body.Close()

	// Get the response body
	respBody, respErr := utils.FetchResponseBody(resp.Body)
	if respErr != nil {
		return respErr
	}

	//marshall response to SemsRespInfo struct
	dataErr := utils.UnmarshalDataToStruct(respBody, &SemsResponseData)
	if dataErr != nil {
		return dataErr
	}

	// check for successful login return value..return a login error
	loginErr := checkUserLoginResponse(SemsResponseData.Msg)
	if loginErr != nil {
		return loginErr
	}

	return nil

}

func DoLoginv2(UserLoginFlow *types.LoginDataFlow) error {
	//check if the UserLogin struct is empty
	if usererr := checkUserLoginInfo(UserLoginFlow.LoginCreds); usererr != nil {
		return usererr
	}

	// User login struct to be converted to JSON
	jsonData, jsonErr := utils.MarshalStructToJSON(UserLoginFlow.LoginCreds)
	if jsonErr != nil {
		return jsonErr
	}

	// Create a new http request
	req, err := http.NewRequest(http.MethodPost, AuthLoginUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	//Add headers pass in the pointer to set the headers on the request object
	setHeaders(req)

	//make the API Call
	client := &http.Client{Timeout: time.Duration(HTTPTimeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	//cleanup
	defer resp.Body.Close()

	// Get the response body
	respBody, respErr := utils.FetchResponseBody(resp.Body)
	if respErr != nil {
		return respErr
	}

	//marshall response to SemsRespInfo struct
	dataErr := utils.UnmarshalDataToStruct(respBody, &UserLoginFlow.LoginResp)
	if dataErr != nil {
		return dataErr
	}

	// check for successful login return value..return a login error
	loginErr := checkUserLoginResponse(UserLoginFlow.LoginResp.Msg)
	if loginErr != nil {
		return loginErr
	}

	return nil
}
