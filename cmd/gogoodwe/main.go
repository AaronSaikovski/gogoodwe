/*
Package main implements a program that authenticates to and queries the SEMS Solar inverter API.
*/
package main

import (
	"context"
	_ "embed"
	"log"
	"time"
	// "os"
	// "runtime/pprof"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/app"
)

const (
	// Context default timeout
	contextTimeout = 60 * time.Second
)

//go:generate bash get_version.sh
//go:embed version.txt
var version string

// main is the entry point of the Go program.
//
// It calls the app.Run function with the version string as a parameter.
// If an error is returned, it logs the error message and exits the program.
func main() {

	// f, err := os.Create("cpu.prof")
	// if err != nil {
	//     panic(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	// Create a context with cancellation capability and 60 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	// Main run
	if err := app.Run(ctx, version); err != nil {
		log.Fatalf("Error: %v", err)
	}

}
