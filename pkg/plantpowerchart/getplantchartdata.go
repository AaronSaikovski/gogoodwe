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
package plantchartdata

import (
	"github.com/AaronSaikovski/gogoodwe/pkg/apihelpers"
	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
)

const (

	// Powerstation API Url
	powerStationURL string = "/v2/Charts/GetPlantPowerChart"

	// Default timeout value
	HTTPTimeout int = 20
)

// --data '{"id":"d035a6b1-a99a-46cf-84a0-ddd5730d7f5c","date":"2024-06-21"}'

// GetMonitorData retrieves monitor data using login credentials and response, storing it in inverterOutput.
//
// Parameters:
// - authLoginInfo: pointer to the LoginInfo struct containing the login credentials and API response
// - inverterOutput: pointer to the data output
// Return type: error
func (plantChartData *PlantPowerChart) GetMonitorData(authLoginInfo *auth.LoginInfo, inverterOutput interface{}) error { // // Get the Token header data

	return apihelpers.FetchMonitorAPIData(authLoginInfo, powerStationURL, HTTPTimeout, inverterOutput)
}

// GetPowerData retrieves the power data for a detailed inverter using the provided authentication information.
//
// Parameters:
// - authLoginInfo: a pointer to the auth.LoginInfo struct containing the login credentials and API response
//
// Returns:
// - error: an error if there was a problem retrieving the power data
func (plantChartData *PlantPowerChart) GetPowerData(authLoginInfo *auth.LoginInfo) error {

	// Get monitor data
	if err := plantChartData.GetMonitorData(authLoginInfo, plantChartData); err != nil {
		return err
	}

	return apihelpers.ProcesData(plantChartData)
}
