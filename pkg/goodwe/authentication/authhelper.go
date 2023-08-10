/*
# Name: authhelper - auth helper functions
# Author: Aaron Saikovski - asaikovski@outlook.com
*/

package authentication

import (
	"errors"
	"strings"

	"github.com/AaronSaikovski/gogoodwe/constants"
	"github.com/AaronSaikovski/gogoodwe/types"
)

// CheckUserLoginInfo - Check user login struct is valid/not null
func CheckUserLoginInfo(UserLogin *types.SemsLoginCreds) error {
	//check if the UserLogin struct is empty
	if (*UserLogin == types.SemsLoginCreds{}) {
		return errors.New("**Error: User Login details are empty or invalid..**")
	} else {
		return nil
	}
}

// CheckUserLoginResponse - check for successful login return value..return a login error
func CheckUserLoginResponse(loginResponse string) error {
	if strings.Compare(loginResponse, constants.SemsLoginSuccessResponse) != 0 {
		return errors.New("API Login Error: " + loginResponse)
	} else {
		return nil
	}
}
