package utils

import (
	"fmt"
	"io"
)

const (
	// MaxResponseSize is the maximum allowed response body size (10MB)
	MaxResponseSize = 10 * 1024 * 1024
)

// FetchResponseBody Get the response body from a HTTP response.
//
// Parameters:
// - resp: an io.Reader representing the HTTP response.
// Returns:
// - []byte: the response body as a byte slice.
// - error: an error if the operation fails.
func FetchResponseBody(resp io.Reader) ([]byte, error) {
	// Limit the response body size to prevent excessive memory usage
	limitedReader := io.LimitReader(resp, MaxResponseSize)
	data, err := io.ReadAll(limitedReader)
	if err != nil {
		return nil, err
	}

	// Check if we hit the limit
	if len(data) >= MaxResponseSize {
		return nil, fmt.Errorf("response body too large (exceeds %d bytes)", MaxResponseSize)
	}

	return data, nil
}
