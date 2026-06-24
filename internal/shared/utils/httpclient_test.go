package utils

import (
	"testing"
	"time"
)

func TestSharedHTTPClient(t *testing.T) {
	if SharedHTTPClient == nil {
		t.Fatal("SharedHTTPClient is nil")
	}
	if SharedHTTPClient.Transport == nil {
		t.Fatal("SharedHTTPClient.Transport is nil")
	}
}

func TestNewHTTPTransport(t *testing.T) {
	transport := NewHTTPTransport()

	if transport == nil {
		t.Fatal("NewHTTPTransport() returned nil")
	}

	if transport.MaxIdleConns != 2 {
		t.Errorf("MaxIdleConns = %d, want 2", transport.MaxIdleConns)
	}
	if transport.MaxIdleConnsPerHost != 1 {
		t.Errorf("MaxIdleConnsPerHost = %d, want 1", transport.MaxIdleConnsPerHost)
	}
	if transport.MaxConnsPerHost != 2 {
		t.Errorf("MaxConnsPerHost = %d, want 2", transport.MaxConnsPerHost)
	}
	if transport.IdleConnTimeout != 10*time.Second {
		t.Errorf("IdleConnTimeout = %v, want %v", transport.IdleConnTimeout, 10*time.Second)
	}
	if transport.TLSHandshakeTimeout != 10*time.Second {
		t.Errorf("TLSHandshakeTimeout = %v, want %v", transport.TLSHandshakeTimeout, 10*time.Second)
	}

	if !transport.ForceAttemptHTTP2 {
		t.Error("ForceAttemptHTTP2 = false, want true")
	}
}
