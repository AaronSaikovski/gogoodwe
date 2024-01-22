/*
# Name: GoGoodwe - Authenticates to and queries the SEMS Solar inverter API
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package main

import (
	"github.com/AaronSaikovski/gogoodwe/app"
)

// main - program main
func main() {

	//setup and run app
	err := app.Run()

	if err != nil {
		panic(err)
	}
}
