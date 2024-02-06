package types

/*
# Name: LoginDataFlow - Struct to hold pointers to User login data structs
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
type LoginDataFlow struct {
	LoginCreds *LoginCredentials
	LoginResp  *LoginResponse
}
