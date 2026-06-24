<<<<<<< HEAD:internal/shared/auth/logincredentials_test.go
package auth

import (
	"testing"
	"time"
)

func TestNewSemsLoginCredentials(t *testing.T) {
	creds := NewSemsLoginCredentials("user@test.com", "password123", "station-id")

	if creds.Account != "user@test.com" {
		t.Errorf("Account = %q, want %q", creds.Account, "user@test.com")
	}
	if creds.Password != "password123" {
		t.Errorf("Password = %q, want %q", creds.Password, "password123")
	}
	if creds.PowerStationID != "station-id" {
		t.Errorf("PowerStationID = %q, want %q", creds.PowerStationID, "station-id")
	}
	if creds.ID != "station-id" {
		t.Errorf("ID = %q, want %q (should match PowerStationID)", creds.ID, "station-id")
	}

	expectedDate := time.Now().Format("2006-01-02")
	if creds.Date != expectedDate {
		t.Errorf("Date = %q, want %q", creds.Date, expectedDate)
=======
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
>>>>>>> origin/main:pkg/auth/logincredentials_test.go
	}
}
