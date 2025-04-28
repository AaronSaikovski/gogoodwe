package apihelpers

import (
	"encoding/json"
	"net/http"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
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
func PowerStationIdJSON(userLogin *auth.SemsLoginCredentials) ([]byte, error) {
	powerStationMap := map[string]string{"powerStationId": userLogin.PowerStationID}
	return json.Marshal(powerStationMap)
}

// dataTokenJSON generates a JSON representation of the data token.
//
// It takes a pointer to an ApiLoginResponse struct 'semsResponseData' as a parameter.
// Returns a byte slice and an error.
func DataTokenJSON(semsResponseData *auth.SemsLoginResponse) ([]byte, error) {
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
