/*
# Name: inverter - gets data from the goodwe API - "v2/PowerStation/GetMonitorDetailByPowerstationId"
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package inverter

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/constants"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

// fetchInverterData - Fetches Data from the Inverter via the specified PowerstationID using theSEMs API
func fetchInverterData(SemsResponseData *types.LoginResponse, UserLogin *types.LoginCredentials, PowerstationOutputData *types.InverterData) error {

	// get the Token header data
	tokenMapJSONData, err := utils.DataTokenJSON(SemsResponseData)
	if err != nil {
		return err
	}

	// get the Powerstation ID header data
	powerStationMapJSONData, err := utils.PowerStationIDJSON(UserLogin)
	if err != nil {
		return err
	}

	//Get the url from the Auth API and append the data url part
	url := SemsResponseData.API + constants.PowerStationURL

	// Create a new http request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(powerStationMapJSONData))
	if err != nil {
		return err
	}

	//Add headers pass in the pointer to set the headers on the request object
	utils.SetHeaders(req, tokenMapJSONData)

	//make the API Call
	client := &http.Client{Timeout: constants.HTTPTimeout * time.Second}
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
