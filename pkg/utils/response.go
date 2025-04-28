package utils

import (
	"io"
)

// FetchResponseBody Get the response body from a HTTP response.
//
// Parameters:
// - resp: an io.Reader representing the HTTP response.
// Returns:
// - []byte: the response body as a byte slice.
// - error: an error if the operation fails.
func FetchResponseBody(resp io.Reader) ([]byte, error) {
	return io.ReadAll(resp)
}
