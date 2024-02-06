/*
# Name: authentication - auth helper functions
# Author: Aaron Saikovski - asaikovski@outlook.com
*/

package authentication

import (
	"errors"
	"net/http"
	"strings"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
)

// SetHeaders - Set the login headers for the SEMS API login
func setHeaders(r *http.Request) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Token", "{\"version\":\"v2.1.0\",\"client\":\"ios\",\"language\":\"en\"}")
}

// CheckUserLoginInfo - Check user login struct is valid/not null
func checkUserLoginInfo(UserLogin *types.LoginCredentials) error {
	if (*UserLogin == types.LoginCredentials{}) {
		return errors.New("**Error: User Login details are empty or invalid..**")
	} else {
		return nil
	}
}

// CheckUserLoginResponse - check for successful login return value..return a login error
func checkUserLoginResponse(loginResponse string) error {
	if strings.Compare(loginResponse, "Successful") != 0 {
		return errors.New("**API Login Error: " + loginResponse)
	} else {
		return nil
	}
}
