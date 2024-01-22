/*
# Name: data - fetches data from the goodwe API - and processes it to pass back to caller
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package inverter

import (
	"errors"
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/authentication"
	"github.com/AaronSaikovski/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/utils"
	"github.com/logrusorgru/aurora"
)

// apiLogin -  Login to the API
func apiLogin(SemsUserLogin *types.LoginCredentials, SemsResponseData *types.LoginResponse) error {

	// Do the login - update the pointer to the struct SemsResponseData
	autherr := authentication.DoLogin(SemsResponseData, SemsUserLogin)
	if autherr != nil {
		utils.HandleError(autherr)
		return autherr
	} else {
		return nil
	}
}

// FetchData - Main API fetch function
func FetchData(Account string, Password string, PowerStationID string) error {

	// Data types
	var SemsResponseData types.LoginResponse
	var PowerstationData types.InverterData

	// Create a new SemsLoginCreds object via a struct literal
	var SemsUserLogin = types.LoginCredentials{
		Account:        Account,
		Password:       Password,
		PowerStationID: PowerStationID,
	}

	// Do the login..check for errors
	err := apiLogin(&SemsUserLogin, &SemsResponseData)
	if err == nil {

		// Fetch the data
		dataerr := fetchInverterData(&SemsResponseData, &SemsUserLogin, &PowerstationData)
		if dataerr != nil {
			utils.HandleError(errors.New("error: fetching powerstation data, check powerstationid is correct"))
			return dataerr
		} else {
			// Get output
			dataOutput, jsonerr := GetDataJSON(&PowerstationData)
			if jsonerr != nil {
				utils.HandleError(errors.New("error: converting powerstation data"))
				return jsonerr

			} else {
				//Display output
				fmt.Println(aurora.BrightYellow(string(dataOutput)))
			}
		}

	} else {
		utils.HandleError(err)
		return err
	}

	return nil
}
