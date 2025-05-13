/*
# Name: jsonutils - helper functions to get the Powerstation Data from the API
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package utils

import (
	"encoding/json"
)

// UnmarshalDataToStruct unmarshals the JSON data from the `respBody` byte slice into the `targetStruct` interface.
//
// Parameters:
// - respBody: a byte slice containing the JSON data to be unmarshaled.
// - targetStruct: an interface{} to store the unmarshaled JSON data.
//
// Returns:
// - error: an error if the unmarshaling process fails.
func UnmarshalDataToStruct(respBody []byte, targetStruct interface{}) error {
	return json.Unmarshal(respBody, &targetStruct)
}

// MarshalStructToJSON marshals the struct pointer to JSON.
//
// Parameters:
// - targetStruct: the struct pointer to be marshaled.
//
// Returns:
// - []byte: the JSON representation of the struct pointer.
// - error: an error if the marshaling process fails.
func MarshalStructToJSON(targetStruct interface{}) ([]byte, error) {
	return json.Marshal(&targetStruct)
}
