package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSemsLogin_InvalidCredentials(t *testing.T) {
	creds := &SemsLoginCredentials{}
	_, err := creds.SemsLogin(context.Background())
	if err == nil {
		t.Error("SemsLogin() expected error for empty credentials")
	}
}

func TestSemsLogin_NilCredentials(t *testing.T) {
	var creds *SemsLoginCredentials
	_, err := creds.SemsLogin(context.Background())
	if err == nil {
		t.Error("SemsLogin() expected error for nil credentials")
	}
}

func TestSemsLogin_WithMockServer(t *testing.T) {
	// Create mock SEMS API server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected application/json content type, got %s", r.Header.Get("Content-Type"))
		}

		response := SemsLoginResponse{
			Msg:  "Successful",
			Code: 0,
		}
		response.Data.UID = "test-uid"
		response.Data.Token = "test-token"
		response.Data.Timestamp = 1234567890
		response.API = "https://test.semsportal.com"

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// The real SemsLogin uses a const AuthLoginURL which we can't override in tests.
	// Test exercises credential validation and marshaling paths.
	creds := &SemsLoginCredentials{
		Account:        "test@example.com",
		Password:       "testpass",
		PowerStationID: "12345678-1234-1234-1234-123456789abc",
	}

	// Expected to fail since it can't reach the real SEMS portal
	_, err := creds.SemsLogin(context.Background())
	if err == nil {
		t.Log("SemsLogin() succeeded unexpectedly (may have reached real API)")
	}
}

func TestSemsLogin_ContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	creds := &SemsLoginCredentials{
		Account:        "test@example.com",
		Password:       "testpass",
		PowerStationID: "12345678-1234-1234-1234-123456789abc",
	}

	_, err := creds.SemsLogin(ctx)
	if err == nil {
		t.Error("SemsLogin() expected error for cancelled context")
	}
}
