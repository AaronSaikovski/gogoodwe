/*
# Name: apiLogin - Logs in to the SEMS API
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package semsapi

import (
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/authentication"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

// ApiLogin -  Login to the SEMS API
func ApiLogin(UserLoginFlow *types.LoginDataFlow) error {

	// Do the login - update the pointer to the struct SemsResponseData
	err := authentication.DoLogin(UserLoginFlow)
	if err != nil {
		utils.HandleError(err)
	}

	return err
}
