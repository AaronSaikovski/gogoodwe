package authentication

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

const (
	AuthLoginURL = "https://www.semsportal.com/api/v2/Common/CrossLogin"
	HTTPTimeout  = 20 // seconds
)

// Login - Login to the SEMS API
func DoLogin(UserLoginFlow *types.LoginDataFlow) error {

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
	req, err := http.NewRequest(http.MethodPost, AuthLoginURL, bytes.NewBuffer(jsonData))
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
