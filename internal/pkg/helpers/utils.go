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
package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
	"github.com/AaronSaikovski/gogoodwe/internal/apilogin"
	"github.com/AaronSaikovski/gogoodwe/internal/pkg/interfaces"
)

// setHeaders sets the headers for the SEMS Data API.
//
// It takes an http.Request pointer 'r' and a byte slice 'tokenstring' as parameters.
func SetHeaders(r *http.Request, tokenstring []byte) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Token", string(tokenstring))
}

// powerStationIdJSON generates a JSON representation of the power station ID.
//
// It takes an ApiLoginCredentials pointer 'userLogin' as a parameter.
// Returns a byte slice and an error.
func PowerStationIdJSON(userLogin *apilogin.ApiLoginCredentials) ([]byte, error) {
	powerStationMap := map[string]string{"powerStationId": userLogin.PowerStationID}
	return json.Marshal(powerStationMap)
}

// dataTokenJSON generates a JSON representation of the data token.
//
// It takes a pointer to an ApiLoginResponse struct 'semsResponseData' as a parameter.
// Returns a byte slice and an error.
func DataTokenJSON(semsResponseData *apilogin.ApiLoginResponse) ([]byte, error) {
	tokenMap := map[string]interface{}{
		"version":   "v2.1.0",
		"client":    "ios",
		"language":  "en",
		"timestamp": semsResponseData.Data.Timestamp,
		"uid":       semsResponseData.Data.UID,
		"token":     semsResponseData.Data.Token,
	}
	return json.Marshal(tokenMap)
}

// getDataJSON generates a JSON representation of the given data.
//
// The function takes a parameter 'data' of type T, which must satisfy the ISemsDataConstraint interface.
// It returns a byte slice containing the JSON representation of the data, and an error if any occurred.
func GetDataJSON[T interfaces.SemsDataConstraint](data T) ([]byte, error) {

	// Get the response and return any errors
	return utils.MarshalStructToJSON(&data)
}
