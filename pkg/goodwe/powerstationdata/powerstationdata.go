/*
# Name: powerstationdata - gets data from the goodwe API - "v2/PowerStation/GetMonitorDetailByPowerstationId"
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package powerstationdata

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/constants"
	"github.com/AaronSaikovski/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/utils"
)

// setHeaders - Set the headers for the SEMS Data API
func setHeaders(r *http.Request, tokenstring []byte) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Token", string(tokenstring))
}

// FetchData - Fetches Data from the specified PowerstationID via tht SEMs API
func FetchData(SemsResponseData *types.SemsResponseData,
	UserLogin *types.SemsLoginCreds,
	PowerstationOutputData *types.StationResponseData) error {

	// get the Token header data
	tokenMapJSONData, _ := DataTokenJSON(SemsResponseData)

	// get the Powerstation ID header data
	powerStationMapJSONData, _ := PowerStationIDJSON(UserLogin)

	//Get the url from the Auth API and append the data url part
	url := SemsResponseData.API + constants.PowerStationURL

	// Create a new http request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(powerStationMapJSONData))
	if err != nil {
		utils.HandleError(err)
	}

	//Add headers pass in the pointer to set the headers on the request object
	setHeaders(req, tokenMapJSONData)

	//make the API Call
	client := &http.Client{Timeout: constants.HTTPTimeout * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		utils.HandleError(err)
		return err
	}

	//cleanup
	defer resp.Body.Close()

	// Get the response body
	respBody, _ := utils.FetchResponseBody(resp.Body)

	//marshall response to SemsRespInfo struct
	dataerr := utils.UnmarshalDataToStruct(respBody, &PowerstationOutputData)
	if dataerr != nil {
		return dataerr
	}

	return nil

}
