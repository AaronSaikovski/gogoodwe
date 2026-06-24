package auth

import "testing"

func TestCheckUserLoginInfo(t *testing.T) {
	tests := []struct {
		name    string
		creds   *SemsLoginCredentials
		wantErr bool
	}{
		{"valid credentials", &SemsLoginCredentials{Account: "user@test.com", Password: "pass", PowerStationID: "12345678-1234-1234-1234-123456789abc"}, false},
		{"nil credentials", nil, true},
		{"empty account", &SemsLoginCredentials{Account: "", Password: "pass", PowerStationID: "id"}, true},
		{"empty password", &SemsLoginCredentials{Account: "user@test.com", Password: "", PowerStationID: "id"}, true},
		{"empty powerstation ID", &SemsLoginCredentials{Account: "user@test.com", Password: "pass", PowerStationID: ""}, true},
		{"all empty", &SemsLoginCredentials{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkUserLoginInfo(tt.creds)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkUserLoginInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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
