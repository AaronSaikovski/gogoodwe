/*
# Name: LoginCredentials - Struct to hold User login data
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package types

// LoginCredentials - Struct to hold User login data
type LoginCredentials struct {
	Account        string `json:"account"`
	Password       string `json:"pwd"`
	PowerStationID string `json:"powerstationid"`
}
