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

func FetchData(Account string, Password string, PowerStationID string) error {

	// User account struct
	creds := &types.LoginCredentials{
		Account:        Account,
		Password:       Password,
		PowerStationID: PowerStationID,
	}

	// Do the login..check for errors
	loginApiResponse, err := semsapi.Login(creds)
	if err != nil {
		utils.HandleError(err)
		return err
	}

	powerstationData, err := getMonitorDetailByPowerstationId(creds, loginApiResponse)
	if err != nil {
		utils.HandleError(err)
		return err
	}

	dataOutput, err := getDataJSON(powerstationData)
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
