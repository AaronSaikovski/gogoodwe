package auth_test

import (
	"testing"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
)

func TestNewSemsLoginCredentials(t *testing.T) {
	tests := []struct {
		name           string
		account        string
		password       string
		powerStationID string
		checkFunc      func(t *testing.T, creds auth.SemsLoginCredentials)
	}{
		{
			name:           "valid credentials",
			account:        "user@example.com",
			password:       "secret123",
			powerStationID: "12345678-abcd-efab-1234-123456789abc",
			checkFunc: func(t *testing.T, creds auth.SemsLoginCredentials) {
				if creds.Account != "user@example.com" {
					t.Errorf("Account = %q, want %q", creds.Account, "user@example.com")
				}
				if creds.Password != "secret123" {
					t.Errorf("Password = %q, want %q", creds.Password, "secret123")
				}
				if creds.PowerStationID != "12345678-abcd-efab-1234-123456789abc" {
					t.Errorf("PowerStationID = %q, want %q", creds.PowerStationID, "12345678-abcd-efab-1234-123456789abc")
				}
				if creds.ID != "12345678-abcd-efab-1234-123456789abc" {
					t.Errorf("ID = %q, want %q", creds.ID, "12345678-abcd-efab-1234-123456789abc")
				}
				if creds.Date == "" {
					t.Error("Date is empty")
				}
			},
		},
		{
			name:           "empty credentials",
			account:        "",
			password:       "",
			powerStationID: "",
			checkFunc: func(t *testing.T, creds auth.SemsLoginCredentials) {
				if creds.Account != "" {
					t.Errorf("Account = %q, want empty", creds.Account)
				}
				if creds.Password != "" {
					t.Errorf("Password = %q, want empty", creds.Password)
				}
				if creds.PowerStationID != "" {
					t.Errorf("PowerStationID = %q, want empty", creds.PowerStationID)
				}
				if creds.Date == "" {
					t.Error("Date is empty")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds := auth.NewSemsLoginCredentials(tt.account, tt.password, tt.powerStationID)
			tt.checkFunc(t, creds)
		})
	}
}
