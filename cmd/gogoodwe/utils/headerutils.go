/*
# Name: powerstationhelper - helper functions to get the Powerstation Data from the API
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package utils

import (
	"net/http"
)

// setHeaders - Set the headers for the SEMS Data API
func SetHeaders(r *http.Request, tokenstring []byte) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Token", string(tokenstring))
}
