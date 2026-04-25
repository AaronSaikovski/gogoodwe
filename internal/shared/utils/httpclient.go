package utils

import (
	"net/http"
	"time"
)

// SharedHTTPClient is a single reusable HTTP client for the entire application.
// Sized appropriately for a CLI tool making 2 sequential requests per invocation.
var SharedHTTPClient = &http.Client{
	Transport: NewHTTPTransport(),
}

// NewHTTPTransport creates an HTTP transport sized for a CLI tool (2 sequential requests).
func NewHTTPTransport() *http.Transport {
	return &http.Transport{
		MaxIdleConns:          2,
		MaxIdleConnsPerHost:   1,
		MaxConnsPerHost:       2,
		IdleConnTimeout:       10 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableCompression:    false,
		ForceAttemptHTTP2:     true,
	}
}
