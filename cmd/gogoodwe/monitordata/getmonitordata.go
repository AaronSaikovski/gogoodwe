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

// getMonitorData retrieves data from the API using the provided login credentials and login API response.
//
// LoginCredentials: The login credentials for the API.
// LoginApiResponse: The login API response.
// InverterOutput: The output struct to store the retrieved data.
// error: An error if any occurred during the retrieval process.
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

// getMonitorDataOutput retrieves monitor data for a given LoginCredentials and LoginApiResponse,
// and then processes the data to generate an output. It takes in three parameters:
// - LoginCredentials: a pointer to an ApiLoginCredentials struct representing the login credentials.
// - LoginApiResponse: a pointer to an ApiLoginResponse struct representing the API login response.
// - InverterOutput: a pointer to a generic type T that implements the ISemsDataConstraint interface.
// The function returns nothing. It first calls the getMonitorData function to retrieve the monitor data,
// then calls the getDataJSON function to convert the data to JSON format, and finally calls the parseOutput
// function to parse the JSON data and print the output. If any errors occur during the process, the
// utils.HandleError function is called to handle the error.
func getMonitorDataOutput[T ISemsDataConstraint](LoginCredentials *apilogin.ApiLoginCredentials, LoginApiResponse *apilogin.ApiLoginResponse, InverterOutput *T) error {

	var powerstationData T
	err := getMonitorData(LoginCredentials, LoginApiResponse, &powerstationData)
	if err != nil {
		return err
	}

	dataOutput, err := getDataJSON(powerstationData)
	if err != nil {
		return err
	}

	output, err := parseOutput(dataOutput)
	if err != nil {
		utils.HandleError(err)
		return err
	}
	printOutput(output)

	return nil
}

// GetMonitorDetailByPowerstationId retrieves the monitor details for a specific power station ID.
//
// LoginCredentials: The login credentials for the API.
// LoginApiResponse: The login API response.
//
// No return value.
func getMonitorDetailByPowerstationId(LoginCredentials *apilogin.ApiLoginCredentials, LoginApiResponse *apilogin.ApiLoginResponse) error {
	var powerstationData InverterData
	err := getMonitorDataOutput(LoginCredentials, LoginApiResponse, &powerstationData)

	if err != nil {
		utils.HandleError(errors.New("error: fetching powerstation data"))
		return err
	}

	return nil
}

// GetMonitorSummaryByPowerstationId retrieves the monitor summary data for a specific power station ID.
//
// LoginCredentials: The login credentials for the API.
// LoginApiResponse: The login API response.
//
// No return value.
func getMonitorSummaryByPowerstationId(LoginCredentials *apilogin.ApiLoginCredentials, LoginApiResponse *apilogin.ApiLoginResponse) error {

	var powerstationData DailySummaryData
	err := getMonitorDataOutput(LoginCredentials, LoginApiResponse, &powerstationData)

	if err != nil {
		utils.HandleError(errors.New("error: fetching powerstation summary data"))
		return err
	}

	return nil

}

// GetData retrieves data based on the provided login credentials and login API response.
//
// LoginCredentials: A pointer to an ApiLoginCredentials struct representing the login credentials.
// LoginApiResponse: A pointer to an ApiLoginResponse struct representing the login API response.
// isDailySummary: A boolean indicating whether to retrieve daily summary data.
func GetData(LoginCredentials *apilogin.ApiLoginCredentials, LoginApiResponse *apilogin.ApiLoginResponse, isDailySummary bool) error {

	if isDailySummary {
		return getMonitorSummaryByPowerstationId(LoginCredentials, LoginApiResponse)

	} else {
		return getMonitorDetailByPowerstationId(LoginCredentials, LoginApiResponse)
	}
}
