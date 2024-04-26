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

package powerstation

import (
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/semsapi"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

// FetchData fetches data based on user account credentials and power station ID, and can retrieve daily summary if specified.
// Parameters:
//
//	Account string - user account
//	Password string - account password
//	PowerStationID string - ID of the power station
//	DailySummary bool - whether to retrieve daily summary
//
// Return type:
//
//	error
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
