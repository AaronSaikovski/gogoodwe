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
package auth

import "time"

// ApiLoginCredentials - Struct to hold User login credentials
type SemsLoginCredentials struct {
	Account        string `json:"account"`
	Password       string `json:"pwd"`
	PowerStationID string `json:"powerstationid"`
	ID             string `json:"id"`
	Date           string `json:"date"` // YYYY-MM-DD
}

func NewSemsLoginCredentials(account, password, powerStationID string) *SemsLoginCredentials {
	return &SemsLoginCredentials{
		Account:        account,
		Password:       password,
		PowerStationID: powerStationID,
		ID:             powerStationID,
		Date:           getFormattedData(),
	}

}

func getFormattedData() string {
	// Get the current date and time
	currentTime := time.Now()

	// Format the date as yyyy-mm-dd
	return currentTime.Format("2006-01-02")

}
