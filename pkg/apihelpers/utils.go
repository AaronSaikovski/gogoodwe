package apihelpers

import (
	"encoding/json"
	"net/http"
	"unsafe"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
	"github.com/AaronSaikovski/gogoodwe/pkg/utils"
)

// bytesToString converts byte slice to string without allocation
func bytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// setHeaders sets the headers for the SEMS Data API.
//
// It takes an http.Request pointer 'r' and a byte slice 'tokenstring' as parameters.
func SetHeaders(r *http.Request, tokenstring []byte) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Token", bytesToString(tokenstring))
}

// setPowerPlantHeaders sets the headers for the Power Plant API.
func SetPowerPlantHeaders(r *http.Request, tokenstring []byte, powerPlantTokenstring []byte) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Token", bytesToString(tokenstring))
	r.Header.Set("data", bytesToString(powerPlantTokenstring))
}

// powerStationIdJSON generates a JSON representation of the power station ID.
//
// It takes an ApiLoginCredentials pointer 'userLogin' as a parameter.
// Returns a byte slice and an error.
func PowerStationIdJSON(userLogin *auth.SemsLoginCredentials) ([]byte, error) {
	// Use struct for better performance and type safety
	powerStationData := struct {
		PowerStationID string `json:"powerStationId"`
	}{
		PowerStationID: userLogin.PowerStationID,
	}
	return json.Marshal(powerStationData)
}

// dataTokenJSON generates a JSON representation of the data token.
//
// It takes a pointer to an ApiLoginResponse struct 'semsResponseData' as a parameter.
// Returns a byte slice and an error.
func DataTokenJSON(semsResponseData *auth.SemsLoginResponse) ([]byte, error) {
	// Use struct for better performance and type safety
	tokenData := struct {
		Version   string `json:"version"`
		Client    string `json:"client"`
		Language  string `json:"language"`
		Timestamp int64  `json:"timestamp"`
		UID       string `json:"uid"`
		Token     string `json:"token"`
	}{
		Version:   "v2.1.0",
		Client:    "ios",
		Language:  "en",
		Timestamp: semsResponseData.Data.Timestamp,
		UID:       semsResponseData.Data.UID,
		Token:     semsResponseData.Data.Token,
	}
	return json.Marshal(tokenData)
}

// PowerPlantdataTokenJSON generates a JSON representation of the data token.
//
// It takes a pointer to an ApiLoginResponse struct 'semsResponseData' as a parameter.
// Returns a byte slice and an error.
func PowerPlantdataTokenJSON(semsResponseData *auth.SemsLoginResponse) ([]byte, error) {
	// Use struct for better performance and type safety
	tokenData := struct {
		ID   string `json:"id"`
		Date string `json:"date"`
	}{
		ID:   semsResponseData.Data.UID,
		Date: utils.GetDate(),
	}
	return json.Marshal(tokenData)
}
