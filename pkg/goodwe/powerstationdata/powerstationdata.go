/*
# Name: powerstationdata - gets data from the goodwe API - "v2/PowerStation/GetMonitorDetailByPowerstationId"
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package powerstationdata

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/pkg/goodwe/constants"
	"github.com/AaronSaikovski/gogoodwe/pkg/goodwe/types"
	"github.com/AaronSaikovski/gogoodwe/pkg/goodwe/utils"
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
	tokenMapJSONData, tokenMapJSONErr := DataTokenJSON(SemsResponseData)
	if tokenMapJSONErr != nil {
		return tokenMapJSONErr
	}

	// get the Powerstation ID header data
	powerStationMapJSONData, powerStationMapJSONErr := PowerStationIDJSON(UserLogin)
	if powerStationMapJSONErr != nil {
		return powerStationMapJSONErr
	}

	//Get the url from the Auth API and append the data url part
	url := SemsResponseData.API + constants.PowerStationURL

	// Create a new http request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(powerStationMapJSONData))
	if err != nil {
		return err
	}

	//Add headers pass in the pointer to set the headers on the request object
	setHeaders(req, tokenMapJSONData)

	//make the API Call
	client := &http.Client{Timeout: constants.HTTPTimeout * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	//cleanup
	defer resp.Body.Close()

	// Get the response body
	respBody, respBodyErr := utils.FetchResponseBody(resp.Body)
	if respBodyErr != nil {
		return respBodyErr
	}

	//marshall response to SemsRespInfo struct
	dataStructErr := utils.UnmarshalDataToStruct(respBody, &PowerstationOutputData)
	if dataStructErr != nil {
		return dataStructErr
	}

	return nil

}
