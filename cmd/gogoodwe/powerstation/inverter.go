/*
# Name: inverter - gets data from the goodwe API - "v2/PowerStation/GetMonitorDetailByPowerstationId"
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package powerstation

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

const (
	// Powerstation API Url
	PowerStationURL string = "v2/PowerStation/GetMonitorDetailByPowerstationId"

	// Default timeout value
	HTTPTimeout int = 20
)

// fetchInverterData - Fetches Data from the Inverter via the specified PowerstationID using the SEMs API
func fetchInverterData(UserLoginFlow *types.LoginDataFlow, PowerstationOutputData *types.InverterData) error {

	// get the Token header data
	tokenMapJSONData, err := utils.DataTokenJSON(UserLoginFlow.LoginResp)
	if err != nil {
		return err
	}

	// get the Powerstation ID header data
	powerStationMapJSONData, err := utils.PowerStationIdJSON(UserLoginFlow.LoginCreds)
	if err != nil {
		return err
	}

	//Get the url from the Auth API and append the data url part
	url := (UserLoginFlow.LoginResp.API + PowerStationURL)

	// Create a new http request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(powerStationMapJSONData))
	if err != nil {
		return err
	}

	//Add headers pass in the pointer to set the headers on the request object
	utils.SetHeaders(req, tokenMapJSONData)

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

	//marshall response to SemsRespInfo struct
	dataStructErr := utils.UnmarshalDataToStruct(respBody, &PowerstationOutputData)
	if dataStructErr != nil {
		return dataStructErr
	}

	return nil

}
