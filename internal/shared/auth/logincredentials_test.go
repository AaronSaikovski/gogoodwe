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
	}
}
