package auth

import (
	"errors"
	"fmt"
	"net/http"
)

// setHeaders sets the headers for the SEMS API login.
//
// It takes a pointer to an http.Request as a parameter and adds the following headers:
// - "Content-Type": "application/json"
// - "Token": "{\"version\":\"v2.1.0\",\"client\":\"ios\",\"language\":\"en\"}"
func setHeaders(r *http.Request) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Token", "{\"version\":\"v2.1.0\",\"client\":\"ios\",\"language\":\"en\"}")
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
	if *userLogin == (SemsLoginCredentials{}) {
		return errors.New("**Error: User Login details are empty or invalid**")
	}
	return nil
}
