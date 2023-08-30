/*
# Name: authentication - authenticates to the goodwe API - https://www.semsportal.com/api/v2/Common/CrossLogin
# Author: Aaron Saikovski - asaikovski@outlook.com
*/

package authentication

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/internal/pkg/constants"
	"github.com/AaronSaikovski/gogoodwe/internal/pkg/entities"
	"github.com/AaronSaikovski/gogoodwe/internal/pkg/utils"
)

// DoLogin - Main public login function
// Logs into the SEMs API
func DoLogin(SemsResponseData *entities.SemsResponseData, UserLogin *entities.SemsLoginCreds) error {

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
	req, err := http.NewRequest(http.MethodPost, constants.AuthLoginUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	//Add headers pass in the pointer to set the headers on the request object
	setHeaders(req)

	//make the API Call
	client := &http.Client{Timeout: constants.HTTPTimeout * time.Second}
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
