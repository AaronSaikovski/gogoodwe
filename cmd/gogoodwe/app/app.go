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

import (
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
	"github.com/alexflint/go-arg"
)

// Main package - This is the main program entry point

// Run is the main program runner.
//
// No parameters.
// Returns an error.
// func Run(versionString string) error {

// 	//Get the args input data
// 	var args utils.Args
// 	//set the version build info
// 	args.SetVersion(versionString)
// 	p := arg.MustParse(&args)

// 	//check for valid email address input
// 	if !utils.CheckValidEmail(args.Account) {
// 		p.Fail("Invalid Email address format - should be: 'user@somedomain.com'.")
// 	}

// 	//check for valid powerstation Id
// 	if !utils.CheckValidPowerstationID(args.PowerStationID) {
// 		p.Fail("Invalid Powerstation ID format: - should be: 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX'.")
// 	}

// 	// Get the data from the API, return any errors. Pass in args as string
// 	return fetchData(args.Account, args.Password, args.PowerStationID, args.DailySummary)

// }

func Run(versionString string) error {
	// Set version build info

	var args utils.Args
	args.SetVersion(versionString)

	// Parse args
	p := arg.MustParse(&args)

	// Check for valid email address input
	if !utils.CheckValidEmail(args.Account) {
		p.Fail("invalid email address format: should be 'user@somedomain.com'")
	}

	// Check for valid powerstation ID
	if !utils.CheckValidPowerstationID(args.PowerStationID) {
		p.Fail("invalid Powerstation ID format: should be 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX'")
	}

	// Get the data from the API, return any errors
	return fetchData(args.Account, args.Password, args.PowerStationID, args.DailySummary)
}
