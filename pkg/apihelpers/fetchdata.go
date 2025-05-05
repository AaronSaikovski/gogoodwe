package apihelpers

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
	"github.com/AaronSaikovski/gogoodwe/pkg/utils"
)

// FetchMonitorData fetches data from the Monitor API.
//
// It takes in the authentication information, the URL of the power station,
// the HTTP timeout, and a pointer to a struct to store the output.
// It returns an error if there was a problem with the API call.
func FetchMonitorAPIData(authLoginInfo *auth.LoginInfo, powerStationURL string, HTTPTimeout int, inverterOutput interface{}) error {

	// Get the Token header data
	apiResponseJSONData, err := DataTokenJSON(authLoginInfo.SemsLoginResponse)
	if err != nil {
		return err
	}

	// //for 'https://au.semsportal.com/api/v2/Charts/GetPlantPowerChart' specific data
	// apiplantPowerResponseJSONData, err := PowerPlantdataTokenJSON(authLoginInfo.SemsLoginResponse)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("apiplantPowerResponseJSONData", string(apiplantPowerResponseJSONData))

	// Get the Powerstation ID header data
	powerStationIDJSONData, err := PowerStationIdJSON(authLoginInfo.SemsLoginCredentials)
	if err != nil {
		return err
	}

	// Create URL from the Auth API and append the data URL part
	url := authLoginInfo.SemsLoginResponse.API + powerStationURL

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(powerStationIDJSONData))
	if err != nil {
		return err
	}

	// Add headers
	SetHeaders(req, apiResponseJSONData)
	//SetPowerPlantHeaders(req, apiResponseJSONData, apiplantPowerResponseJSONData)

	// Make the API call
	client := &http.Client{Timeout: time.Duration(HTTPTimeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Get the response body
	respBody, err := utils.FetchResponseBody(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal response to struct pointer
	if err := utils.UnmarshalDataToStruct(respBody, inverterOutput); err != nil {
		return err
	}

	return nil
}
