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
package app

// Main package - This is the main program entry point
import (
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/apilogin"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/monitordata"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
	"github.com/alexflint/go-arg"
)

// Run - main program runner
func Run() error {

	//Get the args input data
	var args utils.Args
	p := arg.MustParse(&args)

	//check for valid email address input
	if !utils.CheckValidEmail(args.Account) {
		p.Fail("Invalid Email address format - should be: 'user@somedomain.com'.")
	}

	//check for valid powerstation Id
	if !utils.CheckValidPowerstationID(args.PowerStationID) {
		p.Fail("Invalid Powerstation ID format: - should be: 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX'.")
	}

	// Get the data from the API, return any errors. Pass in args as string
	//return powerstation.FetchData(args.Account, args.Password, args.PowerStationID, args.DailySummary)
	return fetchData(args.Account, args.Password, args.PowerStationID, args.DailySummary)

}

// fetchData - Fetches data based on user account credentials and power station ID, and can retrieve daily summary if specified.
func fetchData(Account string, Password string, PowerStationID string, DailySummary bool) error {

	// User account struct
	loginCreds := &apilogin.ApiLoginCredentials{
		Account:        Account,
		Password:       Password,
		PowerStationID: PowerStationID,
	}

	// Do the login..check for errors
	loginApiResponse, err := loginCreds.APILogin()
	if err != nil {
		utils.HandleError(err)
		return err
	}

	//fetch data based on
	if DailySummary {
		monitordata.GetMonitorSummaryByPowerstationId(loginCreds, loginApiResponse)

	} else {
		monitordata.GetMonitorDetailByPowerstationId(loginCreds, loginApiResponse)
	}

	return nil
}
