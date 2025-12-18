/*
Package main implements a program that authenticates to and queries the SEMS Solar inverter API.
*/
package main

import (
	_ "embed"
	"log"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/app"
)

//go:generate bash get_version.sh
//go:embed version.txt
var version string

// main is the entry point of the Go program.
//
// It creates the root Cobra command with the version string and executes it.
// If an error is returned, it logs the error message and exits the program.
func main() {
	rootCmd := app.NewRootCmd(version)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
