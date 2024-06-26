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

/*
Package main implements a program that authenticates to and queries the SEMS Solar inverter API.
*/
package main

import (
	"context"
	_ "embed"
	"log"
	"time"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/app"
)

const (

	//Context default timeout
	contextTimeout = (time.Second * 60)
)

//go:generate bash get_version.sh
//go:embed version.txt
var version string

// main is the entry point of the Go program.
//
// It calls the app.Run function with the version string as a parameter.
// If an error is returned, it logs the error message and exits the program.
func main() {

	// Create a context with cancellation capability and 60 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeout))
	defer cancel()

	// Main run
	if err := app.Run(ctx, version); err != nil {
		log.Fatalf("Error: %v", err)
	}

}
