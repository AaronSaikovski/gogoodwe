/*
# Name: jsonutils - helper functions to get the Powerstation Data from the API
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package utils

import (
	"encoding/json"
)

// UnmarshalDataToStruct - Unmarshall http response to target struct
func UnmarshalDataToStruct(respBody []byte, targetStruct interface{}) error {
	resperr := json.Unmarshal(respBody, &targetStruct)
	return resperr
}

// MarshalStructToJSON - Marshall the struct pointer to JSON
func MarshalStructToJSON(targetStruct interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(&targetStruct)
	return jsonData, err
}
