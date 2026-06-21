package apihelpers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
)

func TestSetHeaders(t *testing.T) {
	token := []byte(`{"version":"v2.1.0","client":"ios"}`)
	req, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
	SetHeaders(req, token)

	if got := req.Header.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}

	if got := req.Header.Get("Token"); got != string(token) {
		t.Errorf("Token = %q, want %q", got, string(token))
	}
}

func TestSetPowerPlantHeaders(t *testing.T) {
	token := []byte(`{"version":"v2.1.0"}`)
	plantToken := []byte(`{"id":"uid","date":"2024-01-01"}`)
	req, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
	SetPowerPlantHeaders(req, token, plantToken)

	if got := req.Header.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}

	if got := req.Header.Get("Token"); got != string(token) {
		t.Errorf("Token = %q, want %q", got, string(token))
	}

	if got := req.Header.Get("data"); got != string(plantToken) {
		t.Errorf("data = %q, want %q", got, string(plantToken))
	}
}

func TestPowerStationIdJSON(t *testing.T) {
	creds := auth.SemsLoginCredentials{PowerStationID: "12345678-abcd-efab-1234-123456789abc"}
	result, err := PowerStationIdJSON(&creds)
	if err != nil {
		t.Fatalf("PowerStationIdJSON() error = %v", err)
	}

	expected := `{"powerStationId":"12345678-abcd-efab-1234-123456789abc"}`
	if string(result) != expected {
		t.Errorf("PowerStationIdJSON() = %s, want %s", string(result), expected)
	}
}

func TestDataTokenJSON(t *testing.T) {
	response := &auth.SemsLoginResponse{
		Data: struct {
			UID       string `json:"uid"`
			Timestamp int64  `json:"timestamp"`
			Token     string `json:"token"`
			Client    string `json:"client"`
			Version   string `json:"version"`
			Language  string `json:"language"`
		}{
			UID:       "user-123",
			Timestamp: 1700000000,
			Token:     "abc-token-xyz",
			Client:    "ios",
			Version:   "v2.1.0",
			Language:  "en",
		},
	}

	result, err := DataTokenJSON(response)
	if err != nil {
		t.Fatalf("DataTokenJSON() error = %v", err)
	}

	expected := `{"version":"v2.1.0","client":"ios","language":"en","timestamp":1700000000,"uid":"user-123","token":"abc-token-xyz"}`
	if string(result) != expected {
		t.Errorf("DataTokenJSON() = %s, want %s", string(result), expected)
	}
}

func TestPowerPlantdataTokenJSON(t *testing.T) {
	response := &auth.SemsLoginResponse{
		Data: struct {
			UID       string `json:"uid"`
			Timestamp int64  `json:"timestamp"`
			Token     string `json:"token"`
			Client    string `json:"client"`
			Version   string `json:"version"`
			Language  string `json:"language"`
		}{
			UID: "user-123",
		},
	}

	result, err := PowerPlantdataTokenJSON(response)
	if err != nil {
		t.Fatalf("PowerPlantdataTokenJSON() error = %v", err)
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Result is not valid JSON: %v", err)
	}

	if parsed["id"] != "user-123" {
		t.Errorf("id = %q, want %q", parsed["id"], "user-123")
	}

	if parsed["date"] == "" {
		t.Error("date is empty")
	}
}
