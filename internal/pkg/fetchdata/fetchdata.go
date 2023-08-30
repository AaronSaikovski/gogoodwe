/*
# Name: fetchdata - fetches data from the goodwe API - and processes it to pass back to caller
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package fetchdata

import (
	"errors"
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/internal/pkg/authentication"
	"github.com/AaronSaikovski/gogoodwe/internal/pkg/entities"
	"github.com/AaronSaikovski/gogoodwe/internal/pkg/powerstation"
	"github.com/AaronSaikovski/gogoodwe/internal/pkg/utils"
	"github.com/logrusorgru/aurora"
)

// apiLogin -  Login to the API
func apiLogin(SemsUserLogin *entities.SemsLoginCreds, SemsResponseData *entities.SemsResponseData) error {

	// Do the login - update the pointer to the struct SemsResponseData
	autherr := authentication.DoLogin(SemsResponseData, SemsUserLogin)
	if autherr != nil {
		utils.HandleError(autherr)
		return autherr
	} else {
		return nil
	}
}

// GetData - Main process data function
func GetData(Account string, Password string, PowerStationID string) error {

	// Data types
	var SemsResponseData entities.SemsResponseData
	var PowerstationData entities.StationResponseData

	// Create a new SemsLoginCreds object via a struct literal
	var SemsUserLogin = entities.SemsLoginCreds{
		Account:        Account,
		Password:       Password,
		PowerStationID: PowerStationID,
	}

	// Do the login..check for errors
	err := apiLogin(&SemsUserLogin, &SemsResponseData)
	if err == nil {

		// Fetch the data
		dataerr := powerstation.FetchData(&SemsResponseData, &SemsUserLogin, &PowerstationData)
		if dataerr != nil {
			utils.HandleError(errors.New("error: fetching powerstation data, check powerstationid is correct"))
			return dataerr
		} else {
			// Get output
			dataOutput, jsonerr := powerstation.GetDataJSON(&PowerstationData)
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
