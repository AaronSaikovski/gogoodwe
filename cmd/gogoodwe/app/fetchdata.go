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
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/apilogin"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/monitordata"
)

// fetchData fetches data based on user account credentials and power station ID, and can retrieve daily summary if specified.
//
// Parameters:
// - Account: the email account associated with the user.
// - Password: the password associated with the user's account.
// - PowerStationID: the ID of the power station.
// - DailySummary: a boolean indicating whether to retrieve a daily summary.
//
// Returns:
// - error: an error if there was a problem logging in or fetching data.
// func fetchData(Account string, Password string, PowerStationID string, DailySummary bool) error {

// 	// User account struct
// 	loginCreds := &apilogin.ApiLoginCredentials{
// 		Account:        Account,
// 		Password:       Password,
// 		PowerStationID: PowerStationID,
// 	}

// 	// Do the login..check for errors
// 	loginApiResponse, err := loginCreds.APILogin()
// 	if err != nil {
// 		utils.HandleError(err)
// 		return err
// 	}

// 	//fetch data and output
// 	dataErr := monitordata.GetData(loginCreds, loginApiResponse, DailySummary)
// 	if dataErr != nil {
// 		utils.HandleError(dataErr)
// 		return dataErr
// 	}

// 	return nil
// }

func fetchData(Account, Password, PowerStationID string, DailySummary bool) error {
	// User account struct
	loginCreds := &apilogin.ApiLoginCredentials{
		Account:        Account,
		Password:       Password,
		PowerStationID: PowerStationID,
	}

	// Do the login
	loginApiResponse, err := loginCreds.APILogin()
	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	// Fetch data and handle errors
	if err := monitordata.GetData(loginCreds, loginApiResponse, DailySummary); err != nil {
		return fmt.Errorf("data fetching failed: %w", err)
	}

	return nil
}
