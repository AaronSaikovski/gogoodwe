package apihelpers

import (
	"context"
	"testing"

	"github.com/AaronSaikovski/gogoodwe/internal/shared/auth"
)

func TestValidateAPIURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{"valid SEMS URL", "https://eu.semsportal.com/api/", false},
		{"valid SEMS root", "https://semsportal.com/api/", false},
		{"valid SEMS subdomain", "https://au.semsportal.com/api/", false},
		{"HTTP not HTTPS", "http://eu.semsportal.com/api/", true},
		{"wrong domain", "https://evil.com/api/", true},
		{"subdomain spoof", "https://semsportal.com.evil.com/api/", true},
		{"empty URL", "", true},
		{"no scheme", "eu.semsportal.com/api/", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAPIURL(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateAPIURL(%q) error = %v, wantErr %v", tt.url, err, tt.wantErr)
			}
		})
	}
}

func TestFetchMonitorAPIData_NilAuth(t *testing.T) {
	_, err := FetchMonitorAPIData(context.Background(), nil, "/test", &struct{}{})
	if err == nil {
		t.Error("FetchMonitorAPIData() expected error for nil auth")
	}
}

func TestFetchMonitorAPIData_EmptyURL(t *testing.T) {
	loginInfo := &auth.LoginInfo{
		SemsLoginCredentials: &auth.SemsLoginCredentials{},
		SemsLoginResponse:    &auth.SemsLoginResponse{},
	}
	_, err := FetchMonitorAPIData(context.Background(), loginInfo, "", &struct{}{})
	if err == nil {
		t.Error("FetchMonitorAPIData() expected error for empty URL")
	}
}

func TestFetchMonitorAPIData_NilResponse(t *testing.T) {
	loginInfo := &auth.LoginInfo{
		SemsLoginCredentials: &auth.SemsLoginCredentials{},
		SemsLoginResponse:    nil,
	}
	_, err := FetchMonitorAPIData(context.Background(), loginInfo, "/test", &struct{}{})
	if err == nil {
		t.Error("FetchMonitorAPIData() expected error for nil response")
	}
}

func TestFetchMonitorAPIData_NilCredentials(t *testing.T) {
	loginInfo := &auth.LoginInfo{
		SemsLoginCredentials: nil,
		SemsLoginResponse:    &auth.SemsLoginResponse{},
	}
	_, err := FetchMonitorAPIData(context.Background(), loginInfo, "/test", &struct{}{})
	if err == nil {
		t.Error("FetchMonitorAPIData() expected error for nil credentials")
	}
}

func TestFetchMonitorAPIData_InvalidAPIURL(t *testing.T) {
	response := &auth.SemsLoginResponse{}
	response.API = "http://evil.com" // not HTTPS, not semsportal.com
	loginInfo := &auth.LoginInfo{
		SemsLoginCredentials: &auth.SemsLoginCredentials{PowerStationID: "test"},
		SemsLoginResponse:    response,
	}
	_, err := FetchMonitorAPIData(context.Background(), loginInfo, "/test", &struct{}{})
	if err == nil {
		t.Error("FetchMonitorAPIData() expected error for invalid API URL")
	}
}
