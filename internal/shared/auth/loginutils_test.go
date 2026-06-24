package auth

<<<<<<< HEAD:internal/shared/auth/loginutils_test.go
import "testing"
=======
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
>>>>>>> origin/main:pkg/auth/loginutils_test.go

func TestCheckUserLoginInfo(t *testing.T) {
	tests := []struct {
		name    string
		creds   *SemsLoginCredentials
		wantErr bool
	}{
<<<<<<< HEAD:internal/shared/auth/loginutils_test.go
		{"valid credentials", &SemsLoginCredentials{Account: "user@test.com", Password: "pass", PowerStationID: "12345678-1234-1234-1234-123456789abc"}, false},
		{"nil credentials", nil, true},
		{"empty account", &SemsLoginCredentials{Account: "", Password: "pass", PowerStationID: "id"}, true},
		{"empty password", &SemsLoginCredentials{Account: "user@test.com", Password: "", PowerStationID: "id"}, true},
		{"empty powerstation ID", &SemsLoginCredentials{Account: "user@test.com", Password: "pass", PowerStationID: ""}, true},
		{"all empty", &SemsLoginCredentials{}, true},
=======
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
>>>>>>> origin/main:pkg/auth/loginutils_test.go
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkUserLoginInfo(tt.creds)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkUserLoginInfo() error = %v, wantErr %v", err, tt.wantErr)
<<<<<<< HEAD:internal/shared/auth/loginutils_test.go
=======
				return
>>>>>>> origin/main:pkg/auth/loginutils_test.go
			}
		})
	}
}
<<<<<<< HEAD:internal/shared/auth/loginutils_test.go

func TestCheckUserLoginResponse(t *testing.T) {
	tests := []struct {
		name     string
		response string
		wantErr  bool
	}{
		{"successful login", "Successful", false},
		{"failed login", "Invalid credentials", true},
		{"empty response", "", true},
		{"wrong case", "successful", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkUserLoginResponse(tt.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkUserLoginResponse(%q) error = %v, wantErr %v", tt.response, err, tt.wantErr)
			}
		})
	}
}

func TestSetHeaders(t *testing.T) {
	// Import net/http is implicit via the package
	req, _ := newTestRequest()

	setHeaders(req)

	if got := req.Header.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}
	if got := req.Header.Get("Token"); got == "" {
		t.Error("Token header not set")
	}
}
=======
>>>>>>> origin/main:pkg/auth/loginutils_test.go
