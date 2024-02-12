/*
# Name: powerstationhelper - helper functions to get the Powerstation Data from the API
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package powerstation

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

// setHeaders - Set the headers for the SEMS Data API
func setHeaders(r *http.Request, tokenstring []byte) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Token", string(tokenstring))
}

// PowerStationIdJSON - Makes a map for the powerStationId to be passed to the Data API header and returns a JSON string
func powerStationIdJSON(UserLogin *types.LoginCredentials) ([]byte, error) {
	powerStationMap := make(map[string]string)
	powerStationMap["powerStationId"] = UserLogin.PowerStationID

	// convert to byte[]
	jsonStr, err := json.Marshal(powerStationMap)
	return jsonStr, err
}

func dataTokenJSON(SemsResponseData *types.LoginResponse) ([]byte, error) {
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

func getDataJSON(PowerstationOutputData *types.InverterData) ([]byte, error) {

	// Get the response and return any errors
	resp, err := utils.MarshalStructToJSON(&PowerstationOutputData)
	return resp, err
}
