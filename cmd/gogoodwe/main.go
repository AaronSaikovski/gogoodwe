/*
# Name: main package - Authenticates to and queries the SEMS Solar inverter API
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package main

// Main package - This is the main program entry point
import (
	"os"

	"github.com/AaronSaikovski/gogoodwe/internal/pkg/fetchdata"
	"github.com/AaronSaikovski/gogoodwe/internal/pkg/utils"
	"github.com/alexflint/go-arg"
)

// run - main program runner
func run() error {
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
	return fetchdata.GetData(args.Account, args.Password, args.PowerStationID)
}

// main - program main
func main() {
	if err := run(); err != nil {
		utils.HandleError(err)
		os.Exit(1)
	}
}
