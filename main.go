/*
# Name: main package - Authenticates to and queries the SEMS Solar inverter API
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package main

// Main package - This is the main program entry point
import (
	"os"

	"github.com/AaronSaikovski/gogoodwe/pkg/goodwe/constants"
	"github.com/AaronSaikovski/gogoodwe/pkg/goodwe/fetchdata"
	"github.com/AaronSaikovski/gogoodwe/pkg/goodwe/types"
	"github.com/AaronSaikovski/gogoodwe/pkg/goodwe/utils"
	"github.com/alexflint/go-arg"
)

// args - srtruct using go-arg- https://github.com/alexflint/go-arg
type args struct {
	Account        string `arg:"required,-a,--account" help:"SEMS Email Account."`
	Pwd            string `arg:"required,-p,--pwd" help:"SEMS Account password."`
	PowerStationID string `arg:"required,-i,--powerstationid" help:"SEMS Powerstation ID."`
}

// Description - App description
func (args) Description() string {
	return "A command line tool to query the GOODWE SEMS Portal APIs and Solar SEMS API."
}

// Version - Version info
func (args) Version() string {
	return constants.VersionString
}

// run - main program entry point
func run() error {
	//Get the args input data
	var args args
	p := arg.MustParse(&args)

	//check for valid email address input
	if !utils.CheckValidEmail(args.Account) {
		p.Fail("Invalid Email address format - should be: 'user@somedomain.com'.")
	}

	//check for valid powerstation Id
	if !utils.CheckValidPowerstationID(args.PowerStationID) {
		p.Fail("Invalid Powerstation ID format: - should be: 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX'.")
	}

	// Create a new SemsLoginCreds object via a struct literal
	SemsUserLogin := types.SemsLoginCreds{
		Account:        args.Account,
		Pwd:            args.Pwd,
		PowerStationID: args.PowerStationID,
	}

	// Get the data from the API, return an errors
	return fetchdata.GetData(&SemsUserLogin)
}

// main - program main
func main() {
	if err := run(); err != nil {
		utils.HandleError(err)
		os.Exit(1)
	}
}
