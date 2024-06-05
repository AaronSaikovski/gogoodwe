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
	"fmt"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
	"github.com/AaronSaikovski/gogoodwe/internal/pkg/helpers"
	"github.com/AaronSaikovski/gogoodwe/internal/pkg/interfaces"
	"github.com/AaronSaikovski/gogoodwe/internal/pkg/types"
)

const (

	// Powerstation API Url
	PowerStationURL string = "v2/PowerStation/GetMonitorDetailByPowerstationId"

	// Default timeout value
	HTTPTimeout int = 20
)

// getMonitorData retrieves monitor data using login credentials and response, storing it in inverterOutput.
//
// Parameters:
// - monitorDataLoginInfo: pointer to the MonitorDataLoginInfo struct containing the login credentials and API response
// - inverterOutput: pointer to the data output
// Return type: error
func getMonitorData[T interfaces.SemsDataConstraint](monitorDataLoginInfo *MonitorDataLoginInfo, inverterOutput *T) error {
	// Get the Token header data
	apiResponseJSONData, err := helpers.DataTokenJSON(monitorDataLoginInfo.LoginApiResponse)
	if err != nil {
		return err
	}

	// Get the Powerstation ID header data
	powerStationIDJSONData, err := helpers.PowerStationIdJSON(monitorDataLoginInfo.LoginApiCredentials)
	if err != nil {
		return err
	}

	// Create URL from the Auth API and append the data URL part
	url := monitorDataLoginInfo.LoginApiResponse.API + PowerStationURL

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(powerStationIDJSONData))
	if err != nil {
		return err
	}

	// Add headers
	helpers.SetHeaders(req, apiResponseJSONData)

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

// getMonitorDataOutput retrieves the monitor data output for a given MonitorDataLoginInfo and stores it in the provided inverterOutput.
//
// Parameters:
// - monitorDataLoginInfo: a pointer to the MonitorDataLoginInfo struct containing the login information for the monitor data.
// - inverterOutput: a pointer to the variable where the monitor data output will be stored.
//
// Returns:
// - error: an error if any occurred during the retrieval or processing of the monitor data.
func getMonitorDataOutput[T interfaces.SemsDataConstraint](monitorDataLoginInfo *MonitorDataLoginInfo, inverterOutput *T) error {
	// Get monitor data
	var powerstationData T
	if err := getMonitorData(monitorDataLoginInfo, &powerstationData); err != nil {
		return err
	}

	// Get data JSON
	dataOutput, err := helpers.GetDataJSON(powerstationData)
	if err != nil {
		return err
	}

	// Parse output
	output, err := utils.ParseOutput(dataOutput)
	if err != nil {
		return err
	}

	// Print output
	utils.PrintOutput(output)

	return nil
}

// getMonitorDetailByPowerstationId retrieves the monitor detail for a given powerstation ID.
//
// It takes a pointer to a MonitorDataLoginInfo struct as a parameter.
// The function returns an error if there was an issue fetching the powerstation data.
func getMonitorDetailByPowerstationId(monitorDataLoginInfo *MonitorDataLoginInfo) error {
	var powerstationData types.InverterData
	if err := getMonitorDataOutput(monitorDataLoginInfo, &powerstationData); err != nil {
		return fmt.Errorf("error fetching powerstation data: %v", err)
	}
	return nil
}

// getMonitorSummaryByPowerstationId retrieves the monitor summary data for a specific power station ID.
//
// Parameters:
// - monitorDataLoginInfo: a pointer to the MonitorDataLoginInfo struct containing the login credentials and power station ID.
//
// Returns:
// - error: an error if there was an issue fetching the powerstation summary data.
func getMonitorSummaryByPowerstationId(monitorDataLoginInfo *MonitorDataLoginInfo) error {
	var powerstationData types.DailySummaryData
	if err := getMonitorDataOutput(monitorDataLoginInfo, &powerstationData); err != nil {
		return fmt.Errorf("error fetching powerstation summary data: %v", err)
	}
	return nil
}

// GetData retrieves either monitor summary or monitor details based on the specified flag.
//
// Parameters:
// - monitorDataLoginInfo: a pointer to the MonitorDataLoginInfo struct
// - isDailySummary: a flag to determine if daily summary data should be retrieved
//
// Returns an error if there was an issue fetching the data.
func (monitorDataLoginInfo *MonitorDataLoginInfo) GetData(isDailySummary bool) error {
	if isDailySummary {
		return getMonitorSummaryByPowerstationId(monitorDataLoginInfo)
	}
	return getMonitorDetailByPowerstationId(monitorDataLoginInfo)
}
