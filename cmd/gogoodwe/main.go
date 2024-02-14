/*
Package main implements a program that authenticates to and queries the SEMS Solar inverter API.
*/
package main

import (
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/app"
	"log"
)

func main() {
	if err := runApp(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func runApp() error {
	return app.Run()
}
