/*
# Name: fetchdata - fetches data from the goodwe API - and processes it to pass back to caller
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package fetchdata

import (
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/pkg/goodwe/authentication"
	"github.com/AaronSaikovski/gogoodwe/pkg/goodwe/powerstationdata"
	"github.com/AaronSaikovski/gogoodwe/pkg/goodwe/types"
	"github.com/logrusorgru/aurora"
)

// doLogin -  Login to the API
func doLogin(SemsUserLogin *types.SemsLoginCreds, SemsResponseData *types.SemsResponseData) error {

	// Do the login - update the pointer to the struct SemsResponseData
	autherr := authentication.DoLogin(SemsResponseData, SemsUserLogin)
	if autherr != nil {
		return autherr
	} else {
		return nil
	}
}

// GetData - Main process data function
func GetData(SemsUserLogin *types.SemsLoginCreds) error {

	// Data types
	var SemsResponseData types.SemsResponseData
	var PowerstationData types.StationResponseData

	// Do the login..check for errors
	err := doLogin(SemsUserLogin, &SemsResponseData)
	if err == nil {

		// Fetch the data
		dataerr := powerstationdata.FetchData(&SemsResponseData, SemsUserLogin, &PowerstationData)
		if dataerr != nil {
			return dataerr
		} else {
			// Get output
			dataOutput, jsonerr := powerstationdata.GetDataJSON(&PowerstationData)
			if jsonerr != nil {
				return jsonerr

			} else {
				//Display output
				fmt.Println(aurora.BrightYellow(string(dataOutput)))
			}
		}

	} else {
		return err
	}

	return nil
}
