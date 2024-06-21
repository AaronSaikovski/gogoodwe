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
package monitordetail

import (
	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
	"github.com/AaronSaikovski/gogoodwe/pkg/helpers"
)

const (

	// Powerstation API Url
	powerStationURL string = "v3/PowerStation/GetMonitorDetailByPowerstationId"

	// Default timeout value
	HTTPTimeout int = 20
)

// GetMonitorData retrieves monitor data using login credentials and response, storing it in inverterOutput.
//
// Parameters:
// - authLoginInfo: pointer to the LoginInfo struct containing the login credentials and API response
// - inverterOutput: pointer to the data output
// Return type: error
func (summaryData *InverterData) GetMonitorData(authLoginInfo *auth.LoginInfo, inverterOutput interface{}) error {
	// // Get the Token header data
	// apiResponseJSONData, err := helpers.DataTokenJSON(authLoginInfo.SemsLoginResponse)
	// if err != nil {
	// 	return err
	// }

	// // Get the Powerstation ID header data
	// powerStationIDJSONData, err := helpers.PowerStationIdJSON(authLoginInfo.SemsLoginCredentials)
	// if err != nil {
	// 	return err
	// }

	// // Create URL from the Auth API and append the data URL part
	// url := authLoginInfo.SemsLoginResponse.API + powerStationURL

	// // Create a new HTTP request
	// req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(powerStationIDJSONData))
	// if err != nil {
	// 	return err
	// }

	// // Add headers
	// helpers.SetHeaders(req, apiResponseJSONData)

	// // Make the API call
	// client := &http.Client{Timeout: time.Duration(HTTPTimeout) * time.Second}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()

	// // Get the response body
	// respBody, err := utils.FetchResponseBody(resp.Body)
	// if err != nil {
	// 	return err
	// }

	// // Unmarshal response to struct pointer
	// if err := utils.UnmarshalDataToStruct(respBody, inverterOutput); err != nil {
	// 	return err
	// }

	// return nil

	return helpers.FetchMonitorData(authLoginInfo, powerStationURL, HTTPTimeout, inverterOutput)
}

// GetPowerData retrieves the power data for a detailed inverter using the provided authentication information.
//
// Parameters:
// - authLoginInfo: a pointer to the auth.LoginInfo struct containing the login credentials and API response
//
// Returns:
// - error: an error if there was a problem retrieving the power data
func (detailData *InverterData) GetPowerData(authLoginInfo *auth.LoginInfo) error {

	// Get monitor data
	if err := detailData.GetMonitorData(authLoginInfo, detailData); err != nil {
		return err
	}

	return helpers.ProcesData(detailData)

	// // Get data JSON
	// dataOutput, err := utils.MarshalStructToJSON(detailData)
	// if err != nil {
	// 	return err
	// }

	// // Parse output
	// output, err := utils.ParseOutput(dataOutput)
	// if err != nil {
	// 	return err
	// }

	// // Print output
	// utils.PrintOutput(output)

	// return nil
}
