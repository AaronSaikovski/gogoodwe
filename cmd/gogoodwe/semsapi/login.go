package semsapi

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

// Login -  Login to the SEMS API passing in a LoginCredentials struct and returning a LoginResponse struct.
func Login(LoginCredentials *types.LoginCredentials) (*types.LoginResponse, error) {

	// API Response struct
	loginApiResponse := types.LoginResponse{}

	//check if the UserLogin struct is empty
	if err := checkUserLoginInfo(LoginCredentials); err != nil {
		return nil, err
	}

	// User login struct to be converted to JSON
	loginData, err := utils.MarshalStructToJSON(LoginCredentials)
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
