package utils

import (
	"encoding/json"
	"io"
)

// FetchResponseBody - Get the response body from a HTTP response
func FetchResponseBody(resp io.Reader) ([]byte, error) {
	respBody, err := io.ReadAll(resp)
	return respBody, err
}

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
