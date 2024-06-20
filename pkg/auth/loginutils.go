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
