package auth

import (
	"errors"
	"fmt"
	"net/http"
)

const tokenHeaderValue = `{"version":"v2.1.0","client":"ios","language":"en"}`

// setHeaders sets the headers for the SEMS API login.
//
// It takes a pointer to an http.Request as a parameter and adds the following headers:
// - "Content-Type": "application/json"
// - "Token": "{\"version\":\"v2.1.0\",\"client\":\"ios\",\"language\":\"en\"}"
func setHeaders(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Token", tokenHeaderValue)
}

// checkUserLoginResponse checks for successful login return value and returns a login error if unsuccessful.
//
// Parameter:
// - loginResponse: the response string to check.
// Return type: error
func checkUserLoginResponse(loginResponse string) error {
	if loginResponse != "Successful" {
		return fmt.Errorf("**API Login Error: %s", loginResponse)
	}
	return nil
}

// checkUserLoginInfo checks if the provided user login credentials are valid or not.
//
// It takes a pointer to an ApiLoginCredentials struct as a parameter and returns an error if the credentials are empty or invalid.
// The function returns nil if the credentials are valid.
func checkUserLoginInfo(userLogin *SemsLoginCredentials) error {
	// Check individual required fields instead of comparing entire struct
	if userLogin == nil || userLogin.Account == "" || userLogin.Password == "" || userLogin.PowerStationID == "" {
		return errors.New("**Error: User Login details are empty or invalid**")
	}
	return nil
}

// setPowerPlantHeaders sets the headers for the Power Plant API.
// func setPowerPlantHeaders(r *http.Request) {
// 	r.Header.Add("Content-Type", "application/json")
// 	r.Header.Add("Token", "{\"version\":\"v2.1.0\",\"client\":\"ios\",\"language\":\"en\"}")
// 	r.Header.Add("data", "{\"id\":\"GUID\",\"date\":\"2024-06-21\"}")
// }
