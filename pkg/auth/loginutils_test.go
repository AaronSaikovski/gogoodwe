package auth

import (
	"net/http"
	"testing"
)

func TestSetHeaders(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
	setHeaders(req)

	if got := req.Header.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}

	if got := req.Header.Get("Token"); got != tokenHeaderValue {
		t.Errorf("Token = %q, want %q", got, tokenHeaderValue)
	}
}

func TestCheckUserLoginResponse(t *testing.T) {
	tests := []struct {
		name          string
		loginResponse string
		wantErr       bool
		errContains   string
	}{
		{name: "successful login", loginResponse: "Successful", wantErr: false},
		{name: "failed login", loginResponse: "Failed", wantErr: true, errContains: "Failed"},
		{name: "error message", loginResponse: "Invalid credentials", wantErr: true, errContains: "Invalid credentials"},
		{name: "empty response", loginResponse: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkUserLoginResponse(tt.loginResponse)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkUserLoginResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.errContains != "" {
				expected := "**API Login Error: " + tt.loginResponse
				if err.Error() != expected {
					t.Errorf("checkUserLoginResponse() error = %v, want %v", err, expected)
				}
			}
		})
	}
}

func TestCheckUserLoginInfo(t *testing.T) {
	tests := []struct {
		name    string
		creds   *SemsLoginCredentials
		wantErr bool
	}{
		{
			name: "valid credentials",
			creds: &SemsLoginCredentials{
				Account:        "user@example.com",
				Password:       "secret",
				PowerStationID: "12345678-abcd-efab-1234-123456789abc",
			},
			wantErr: false,
		},
		{
			name:    "nil credentials",
			creds:   nil,
			wantErr: true,
		},
		{
			name:    "empty account",
			creds:   &SemsLoginCredentials{Password: "secret", PowerStationID: "12345678-abcd-efab-1234-123456789abc"},
			wantErr: true,
		},
		{
			name:    "empty password",
			creds:   &SemsLoginCredentials{Account: "user@example.com", PowerStationID: "12345678-abcd-efab-1234-123456789abc"},
			wantErr: true,
		},
		{
			name:    "empty powerstation ID",
			creds:   &SemsLoginCredentials{Account: "user@example.com", Password: "secret"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkUserLoginInfo(tt.creds)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkUserLoginInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
