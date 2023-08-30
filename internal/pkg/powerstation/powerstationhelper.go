/*
# Name: powerstationhelper - helper functions to get the Powerstation Data from the API
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package powerstation

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AaronSaikovski/gogoodwe/internal/pkg/entities"
	"github.com/AaronSaikovski/gogoodwe/internal/pkg/utils"
)

// setHeaders - Set the headers for the SEMS Data API
func setHeaders(r *http.Request, tokenstring []byte) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Token", string(tokenstring))
}

// DataTokenJSON - Makes a map for the token to be passed to the Data API header and returns a JSON string
func dataTokenJSON(SemsResponseData *entities.SemsResponseData) ([]byte, error) {

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
func powerStationIDJSON(UserLogin *entities.SemsLoginCreds) ([]byte, error) {
	powerStationMap := make(map[string]string)
	powerStationMap["powerStationId"] = UserLogin.PowerStationID

	// convert to byte[]
	jsonStr, err := json.Marshal(powerStationMap)
	return jsonStr, err
}

// GetDataJSON - Returns the PowerstationOutputData as JSON
func GetDataJSON(PowerstationOutputData *entities.StationResponseData) ([]byte, error) {

	// Get the response and return any errors
	resp, err := utils.MarshalStructToJSON(&PowerstationOutputData)
	return resp, err
}
