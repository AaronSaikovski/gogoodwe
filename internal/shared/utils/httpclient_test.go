<<<<<<< HEAD:internal/shared/utils/httpclient_test.go
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
=======
package utils_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/AaronSaikovski/gogoodwe/pkg/utils"
)

func TestNewHTTPTransport(t *testing.T) {
	transport := utils.NewHTTPTransport()

	if transport.MaxIdleConns != 100 {
		t.Errorf("MaxIdleConns = %d, want 100", transport.MaxIdleConns)
	}

	if transport.MaxIdleConnsPerHost != 10 {
		t.Errorf("MaxIdleConnsPerHost = %d, want 10", transport.MaxIdleConnsPerHost)
	}

	if transport.MaxConnsPerHost != 100 {
		t.Errorf("MaxConnsPerHost = %d, want 100", transport.MaxConnsPerHost)
	}

	if transport.IdleConnTimeout != 90*time.Second {
		t.Errorf("IdleConnTimeout = %v, want 90s", transport.IdleConnTimeout)
	}

	if transport.TLSHandshakeTimeout.String() != "10s" {
		t.Errorf("TLSHandshakeTimeout = %v, want 10s", transport.TLSHandshakeTimeout)
	}

	if transport.ResponseHeaderTimeout.String() != "10s" {
		t.Errorf("ResponseHeaderTimeout = %v, want 10s", transport.ResponseHeaderTimeout)
	}

	if transport.ExpectContinueTimeout.String() != "1s" {
		t.Errorf("ExpectContinueTimeout = %v, want 1s", transport.ExpectContinueTimeout)
	}

	if transport.DisableCompression {
		t.Error("DisableCompression = true, want false")
	}

>>>>>>> origin/main:pkg/utils/httpclient_test.go
	if !transport.ForceAttemptHTTP2 {
		t.Error("ForceAttemptHTTP2 = false, want true")
	}
}
<<<<<<< HEAD:internal/shared/utils/httpclient_test.go
=======

func TestNewHTTPTransportCreatesUsableClient(t *testing.T) {
	transport := utils.NewHTTPTransport()
	_ = &http.Client{Transport: transport}
}
>>>>>>> origin/main:pkg/utils/httpclient_test.go
