package utils

import "github.com/AaronSaikovski/gogoodwe/pkg/goodwe/constants"

// Args - struct using go-arg- https://github.com/alexflint/go-arg
type Args struct {
	Account        string `arg:"required,-a,--account" help:"SEMS Email Account."`
	Pwd            string `arg:"required,-p,--pwd" help:"SEMS Account password."`
	PowerStationID string `arg:"required,-i,--powerstationid" help:"SEMS Powerstation ID."`
}

// Description - App description
func (Args) Description() string {
	return "A command line tool to query the GOODWE SEMS Portal APIs and Solar SEMS API."
}

// Version - Version info
func (Args) Version() string {
	return constants.VersionString
}