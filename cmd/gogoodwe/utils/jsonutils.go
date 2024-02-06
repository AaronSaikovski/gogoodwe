/*
# Name: jsonutils - helper functions to get the Powerstation Data from the API
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package utils

import (
	"encoding/json"
	"strconv"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
)

// DataTokenJSON - Makes a map for the token to be passed to the Data API header and returns a JSON string
func DataTokenJSON(SemsResponseData *types.LoginResponse) ([]byte, error) {
	tokenMap := make(map[string]string)
	tokenMap["version"] = "v2.1.0"
	tokenMap["client"] = "ios"
	tokenMap["language"] = "en"
	tokenMap["timestamp"] = strconv.FormatInt(SemsResponseData.Data.Timestamp, 10)
	tokenMap["uid"] = SemsResponseData.Data.UID
	tokenMap["token"] = SemsResponseData.Data.Token

	// convert to byte[]
	jsonStr, err := json.Marshal(tokenMap)
	return jsonStr, err
}

// PowerStationIDJSON - Makes a map for the powerStationId to be passed to the Data API header and returns a JSON string
func PowerStationIDJSON(UserLogin *types.LoginCredentials) ([]byte, error) {
	powerStationMap := make(map[string]string)
	powerStationMap["powerStationId"] = UserLogin.PowerStationID

	// convert to byte[]
	jsonStr, err := json.Marshal(powerStationMap)
	return jsonStr, err
}

// GetDataJSON - Returns the PowerstationOutputData as JSON
func GetDataJSON(PowerstationOutputData *types.InverterData) ([]byte, error) {

	// Get the response and return any errors
	resp, err := MarshalStructToJSON(&PowerstationOutputData)
	return resp, err
}

// UnmarshalDataToStruct - Unmarshall http response to target struct
func UnmarshalDataToStruct(respBody []byte, targetStruct interface{}) error {
	resperr := json.Unmarshal(respBody, &targetStruct)
	return resperr
}

// MarshalStructToJSON - Marshall the struct pointer to JSON
func MarshalStructToJSON(targetStruct interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(&targetStruct)
	return jsonData, err
}
