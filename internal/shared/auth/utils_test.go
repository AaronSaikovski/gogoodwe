package auth

import (
	"encoding/json"
	"net/http"
	"testing"
)

func newTestRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodPost, "https://example.com", nil)
}

func TestSetHeadersData(t *testing.T) {
	req, _ := newTestRequest()
	tokenData := []byte(`{"token":"abc123"}`)

	SetHeaders(req, tokenData)

	if got := req.Header.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}
	if got := req.Header.Get("Token"); got != `{"token":"abc123"}` {
		t.Errorf("Token = %q, want %q", got, `{"token":"abc123"}`)
	}
}

func TestSetPowerPlantHeaders(t *testing.T) {
	req, _ := newTestRequest()
	tokenData := []byte(`{"token":"abc"}`)
	ppData := []byte(`{"id":"xyz"}`)

	SetPowerPlantHeaders(req, tokenData, ppData)

	if got := req.Header.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}
	if got := req.Header.Get("Token"); got != `{"token":"abc"}` {
		t.Errorf("Token = %q, want %q", got, `{"token":"abc"}`)
	}
	if got := req.Header.Get("data"); got != `{"id":"xyz"}` {
		t.Errorf("data = %q, want %q", got, `{"id":"xyz"}`)
	}
}

func TestPowerStationIdJSON(t *testing.T) {
	creds := &SemsLoginCredentials{PowerStationID: "test-id-123"}

	result, err := PowerStationIdJSON(creds)
	if err != nil {
		t.Fatalf("PowerStationIdJSON() unexpected error: %v", err)
	}

	var parsed struct {
		PowerStationID string `json:"powerStationId"`
	}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}
	if parsed.PowerStationID != "test-id-123" {
		t.Errorf("PowerStationID = %q, want %q", parsed.PowerStationID, "test-id-123")
	}
}

func TestDataTokenJSON(t *testing.T) {
	response := &SemsLoginResponse{}
	response.Data.UID = "uid-123"
	response.Data.Token = "token-abc"
	response.Data.Timestamp = 1234567890

	result, err := DataTokenJSON(response)
	if err != nil {
		t.Fatalf("DataTokenJSON() unexpected error: %v", err)
	}

	var parsed struct {
		Version   string `json:"version"`
		Client    string `json:"client"`
		Language  string `json:"language"`
		Timestamp int64  `json:"timestamp"`
		UID       string `json:"uid"`
		Token     string `json:"token"`
	}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	if parsed.Version != apiVersion {
		t.Errorf("Version = %q, want %q", parsed.Version, apiVersion)
	}
	if parsed.Client != apiClient {
		t.Errorf("Client = %q, want %q", parsed.Client, apiClient)
	}
	if parsed.Language != apiLanguage {
		t.Errorf("Language = %q, want %q", parsed.Language, apiLanguage)
	}
	if parsed.UID != "uid-123" {
		t.Errorf("UID = %q, want %q", parsed.UID, "uid-123")
	}
	if parsed.Token != "token-abc" {
		t.Errorf("Token = %q, want %q", parsed.Token, "token-abc")
	}
	if parsed.Timestamp != 1234567890 {
		t.Errorf("Timestamp = %d, want %d", parsed.Timestamp, 1234567890)
	}
}

func TestPowerPlantdataTokenJSON(t *testing.T) {
	response := &SemsLoginResponse{}
	response.Data.UID = "uid-456"

	result, err := PowerPlantdataTokenJSON(response)
	if err != nil {
		t.Fatalf("PowerPlantdataTokenJSON() unexpected error: %v", err)
	}

	var parsed struct {
		ID   string `json:"id"`
		Date string `json:"date"`
	}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}
	if parsed.ID != "uid-456" {
		t.Errorf("ID = %q, want %q", parsed.ID, "uid-456")
	}
	if parsed.Date == "" {
		t.Error("Date should not be empty")
	}
}
