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
