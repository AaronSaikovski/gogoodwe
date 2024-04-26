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
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/apilogin"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

// setHeaders - Set the headers for the SEMS Data API
func setHeaders(r *http.Request, tokenstring []byte) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Token", string(tokenstring))
}

// PowerStationIdJSON - Makes a map for the powerStationId to be passed to the Data API header and returns a JSON string
func powerStationIdJSON(UserLogin *apilogin.ApiLoginCredentials) ([]byte, error) {
	powerStationMap := make(map[string]string)
	powerStationMap["powerStationId"] = UserLogin.PowerStationID

	// convert to byte[]
	jsonStr, err := json.Marshal(powerStationMap)
	return jsonStr, err
}

func dataTokenJSON(SemsResponseData *apilogin.ApiLoginResponse) ([]byte, error) {
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

// parse json data
func getDataJSON[T ISemsDataConstraint](data T) ([]byte, error) {

	// Get the response and return any errors
	resp, err := utils.MarshalStructToJSON(&data)
	return resp, err
}
