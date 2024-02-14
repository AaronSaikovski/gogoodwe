package utils

import (
	"io"
)

// FetchResponseBody - Get the response body from a HTTP response
func FetchResponseBody(resp io.Reader) ([]byte, error) {
	respBody, err := io.ReadAll(resp)
	return respBody, err
}
