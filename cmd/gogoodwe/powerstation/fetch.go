/*
# Name: data - fetches data from the goodwe API - and processes it to pass back to caller
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package powerstation

import (
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/semsapi"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

func FetchData(Account string, Password string, PowerStationID string, DailySummary bool) error {

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

	//fetch data based on
	if DailySummary {
		getMonitorSummaryByPowerstationId(creds, loginApiResponse)

	} else {
		//powerstationData = types.InverterData
		getMonitorDetailByPowerstationId(creds, loginApiResponse)
	}

	return nil
}
