/*
# Name: authentication - authenticates to the goodwe API - https://www.semsportal.com/api/v2/Common/CrossLogin
# Author: Aaron Saikovski - asaikovski@outlook.com
*/

package authentication

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/constants"
	"github.com/AaronSaikovski/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/utils"
)

// SetHeaders - Set the login headers for the SEMS API login
func SetHeaders(r *http.Request) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Token", "{\"version\":\"v2.1.0\",\"client\":\"ios\",\"language\":\"en\"}")
}

// DoLogin - Main public login function
// Logs into the SEMs API
func DoLogin(SemsResponseData *types.SemsResponseData, UserLogin *types.SemsLoginCreds) error {

	//check if the UserLogin struct is empty
	usererr := CheckUserLoginInfo(UserLogin)
	if usererr != nil {
		return usererr
	}

	// User login struct to be converted to JSON
	jsonData, _ := utils.MarshalStructToJSON(UserLogin)

	// Create a new http request
	req, err := http.NewRequest(http.MethodPost, constants.AuthLoginUrL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	//Add headers pass in the pointer to set the headers on the request object
	SetHeaders(req)

	//make the API Call
	client := &http.Client{Timeout: constants.HTTPTimeout * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	//cleanup
	defer resp.Body.Close()

	// Get the response body
	respBody, _ := utils.FetchResponseBody(resp.Body)

	//marshall response to SemsRespInfo struct
	dataErr := utils.UnmarshalDataToStruct(respBody, &SemsResponseData)
	if dataErr != nil {
		return dataErr
	}

	// check for successful login return value..return a login error
	loginErr := CheckUserLoginResponse(SemsResponseData.Msg)

	// Return the loginerror
	if loginErr != nil {
		return loginErr
	}

	return nil

}
