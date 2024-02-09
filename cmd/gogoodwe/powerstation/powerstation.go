/*
# Name: data - fetches data from the goodwe API - and processes it to pass back to caller
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package powerstation

import (
	"errors"
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/semsapi"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
	"github.com/logrusorgru/aurora"
	"github.com/valyala/fastjson"
)

// fetchInverterData - Fetches Data from the Inverter via the specified PowerstationID using the SEMs API
func FetchData(Account string, Password string, PowerStationID string) error {

	// User account struct
	creds := &types.LoginCredentials{
		Account:        Account,
		Password:       Password,
		PowerStationID: PowerStationID,
	}

	// Create a new LoginDataFlow object reference
	loginDataFlow := &types.LoginDataFlow{
		LoginCreds: creds,
		LoginResp:  &types.LoginResponse{},
	}

	// Do the login..check for errors
	err := semsapi.ApiLogin(loginDataFlow)
	if err != nil {
		utils.HandleError(err)
		return err
	}

	// Powerstation Output Data
	powerstationData := types.InverterData{}

	// Fetch the data
	fetchDataerr := fetchInverterData(loginDataFlow, &powerstationData)
	if fetchDataerr != nil {
		utils.HandleError(errors.New("error: fetching powerstation data, check powerstationid is correct"))
		return fetchDataerr
	}

	// Get output
	dataOutput, err := utils.GetDataJSON(&powerstationData)
	if err != nil {
		utils.HandleError(errors.New("error: converting powerstation data"))
		return err
	}

	//parse JSON output
	var parser fastjson.Parser
	output, err := parser.Parse(string(dataOutput))
	if err != nil {
		utils.HandleError(errors.New("error: parsing powerstation data"))
		return err
	}

	fmt.Println(aurora.BrightYellow(output))

	return nil
}
