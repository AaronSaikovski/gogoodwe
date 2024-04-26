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
package monitordata

import (
	"bytes"
	"errors"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/apilogin"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

// Generic function to retrieve data from the API via an ISemsDataConstraint Interface of defined structs
func getMonitorData[T ISemsDataConstraint](LoginCredentials *apilogin.ApiLoginCredentials, LoginApiResponse *apilogin.ApiLoginResponse, InverterOutput *T) error {

	// get the Token header data
	apiResponseJsonData, err := dataTokenJSON(LoginApiResponse)
	if err != nil {
		return err
	}

	// get the Powerstation ID header data
	powerStationIdJsonData, err := powerStationIdJSON(LoginCredentials)
	if err != nil {
		return err
	}

	//Get the url from the Auth API and append the data url part
	url := (LoginApiResponse.API + PowerStationURL)

	// Create a new http request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(powerStationIdJsonData))
	if err != nil {
		return err
	}

	//Add headers pass in the pointer to set the headers on the request object
	setHeaders(req, apiResponseJsonData)

	//make the API Call
	client := &http.Client{Timeout: time.Duration(HTTPTimeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	//cleanup
	defer resp.Body.Close()

	// Get the response body
	respBody, err := utils.FetchResponseBody(resp.Body)
	if err != nil {
		return err
	}

	//marshall response to struct pointer
	inverterDataerr := utils.UnmarshalDataToStruct(respBody, &InverterOutput)
	if inverterDataerr != nil {
		return inverterDataerr
	}

	return nil

}

// Get Monitor Detailed data
func GetMonitorDetailByPowerstationId(LoginCredentials *apilogin.ApiLoginCredentials, LoginApiResponse *apilogin.ApiLoginResponse) {
	var powerstationData InverterData

	err := getMonitorData(LoginCredentials, LoginApiResponse, &powerstationData)
	if err != nil {
		utils.HandleError(errors.New("error: fetching powerstation data"))
	}

	dataOutput, err := getDataJSON(powerstationData)
	if err != nil {
		utils.HandleError(errors.New("error: converting powerstation data"))
	}

	output, err := parseOutput(dataOutput)
	if err != nil {
		utils.HandleError(err)
	}
	printOutput(output)

}

// Get Monitor summary data
func GetMonitorSummaryByPowerstationId(LoginCredentials *apilogin.ApiLoginCredentials, LoginApiResponse *apilogin.ApiLoginResponse) {

	var powerstationData DailySummaryData
	err := getMonitorData(LoginCredentials, LoginApiResponse, &powerstationData)
	if err != nil {
		utils.HandleError(errors.New("error: fetching powerstation summary data"))
	}

	dataOutput, err := getDataJSON(powerstationData)
	if err != nil {
		utils.HandleError(errors.New("error: converting powerstation summary data"))
	}

	output, err := parseOutput(dataOutput)
	if err != nil {
		utils.HandleError(err)
	}
	printOutput(output)

}
