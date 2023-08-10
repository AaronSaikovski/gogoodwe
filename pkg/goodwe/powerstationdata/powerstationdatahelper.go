/*
# Name: powerstationdatahelper - helper functions to get the Powerstation Data from the API
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package powerstationdata

import (
	"encoding/json"
	"strconv"

	"github.com/AaronSaikovski/gogoodwe/pkg/goodwe/types"
	"github.com/AaronSaikovski/gogoodwe/pkg/goodwe/utils"

// DataTokenJSON - Makes a map for the token to be passed to the Data API header and returns a JSON string
func DataTokenJSON(SemsResponseData *types.SemsResponseData) ([]byte, error) {

	tokenMap := make(map[string]string)
	tokenMap["version"] = "v2.1.0"
	tokenMap["client"] = "ios"
	tokenMap["language"] = "en"
	tokenMap["timestamp"] = strconv.Itoa(SemsResponseData.Data.Timestamp)
	tokenMap["uid"] = SemsResponseData.Data.UID
	tokenMap["token"] = SemsResponseData.Data.Token

	// convert to byte[]
	jsonStr, err := json.Marshal(tokenMap)
	return jsonStr, err
}

// PowerStationIDJSON - Makes a map for the powerStationId to be passed to the Data API header and returns a JSON string
func PowerStationIDJSON(UserLogin *types.SemsLoginCreds) ([]byte, error) {
	powerStationMap := make(map[string]string)
	powerStationMap["powerStationId"] = UserLogin.PowerStationID

	// convert to byte[]
	jsonStr, err := json.Marshal(powerStationMap)
	return jsonStr, err
}

// GetDataJSON - Returns the PowerstationOutputData as JSON
func GetDataJSON(PowerstationOutputData *types.StationResponseData) ([]byte, error) {

	// Get the response and return any errors
	resp, err := utils.MarshalStructToJSON(&PowerstationOutputData)
	return resp, err
}
