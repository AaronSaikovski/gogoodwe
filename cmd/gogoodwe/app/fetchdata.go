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
	"context"
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
	"github.com/AaronSaikovski/gogoodwe/pkg/interfaces"
	"github.com/AaronSaikovski/gogoodwe/pkg/monitordetail"
	// 	"github.com/AaronSaikovski/gogoodwe/pkg/monitordetail"
	// "github.com/AaronSaikovski/gogoodwe/pkg/monitorsummary"
)

// fetchData fetches data using the provided account credentials and power station ID.
//
// Account: the email account associated with the user.
// Password: the password associated with the user's account.
// PowerStationID: the ID of the power station.
// DailySummary: a boolean indicating whether to retrieve a daily summary.
// error: an error if there was a problem logging in or fetching data.
func fetchData(context context.Context, Account, Password, PowerStationID string, isDailySummary bool) error {

	// User account struct
	apiLoginCreds := auth.NewSemsLoginCredentials(Account, Password, PowerStationID, "au")

	// Assign the login interface
	var loginService interfaces.SemsLogin = apiLoginCreds

	// Do the login
	loginApiResponse, err := loginService.SemsLogin()
	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	//Populate the loginInfo struct
	loginInfo := &auth.LoginInfo{
		SemsLoginCredentials: apiLoginCreds,
		SemsLoginResponse:    loginApiResponse,
	}

	// Assign the login interface
	//var dataService interfaces.PowerData = monitorsummary.NewDailySummaryData()
	var dataService interfaces.PowerData = monitordetail.NewMonitorData()

	if err := dataService.GetPowerData(loginInfo); err != nil {
		return fmt.Errorf("data retrieval failed: %w", err)
	}

	if err := context.Err(); err != nil {
		return fmt.Errorf("context error: %w", err)
	}

	return nil
}
