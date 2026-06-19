package exporthistory

import (
	"testing"

	"github.com/AaronSaikovski/gogoodwe/internal/shared/auth"
)

func TestParseTargets(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantCount int
		wantError bool
	}{
		{
			name:      "empty returns defaults",
			input:     "",
			wantCount: 14,
			wantError: false,
		},
		{
			name:      "single valid target",
			input:     "Vpv1",
			wantCount: 1,
			wantError: false,
		},
		{
			name:      "multiple valid targets",
			input:     "Vpv1,Vpv2,Ipv1",
			wantCount: 3,
			wantError: false,
		},
		{
			name:      "targets with spaces",
			input:     "Vpv1, Vpv2, Ipv1",
			wantCount: 3,
			wantError: false,
		},
		{
			name:      "all valid targets",
			input:     "Vpv1,Vpv2,Ipv1,Ipv2,Vac1,Iac1,Fac1,Pac,WorkMode,Tempperature,ETotal,HTotal,Reserved5,PF",
			wantCount: 14,
			wantError: false,
		},
		{
			name:      "invalid target key",
			input:     "Vpv1,InvalidKey",
			wantCount: 0,
			wantError: true,
		},
		{
			name:      "wrong case target",
			input:     "vpv1",
			wantCount: 0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			targets, err := ParseTargets(tt.input)

			if (err != nil) != tt.wantError {
				t.Errorf("ParseTargets() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !tt.wantError && len(targets) != tt.wantCount {
				t.Errorf("ParseTargets() returned %d targets, want %d", len(targets), tt.wantCount)
			}
		})
	}
}

func TestParseTargetsCorrectIndex(t *testing.T) {
	targets, err := ParseTargets("Vpv1,PF")
	if err != nil {
		t.Fatalf("ParseTargets() unexpected error: %v", err)
	}

	if targets[0].TargetKey != "Vpv1" || targets[0].TargetIndex != 1 {
		t.Errorf("Expected Vpv1 with index 1, got %s with index %d", targets[0].TargetKey, targets[0].TargetIndex)
	}
	if targets[1].TargetKey != "PF" || targets[1].TargetIndex != 407 {
		t.Errorf("Expected PF with index 407, got %s with index %d", targets[1].TargetKey, targets[1].TargetIndex)
	}
}

func TestBuildRequest(t *testing.T) {
	loginInfo := &auth.LoginInfo{
		SemsLoginCredentials: &auth.SemsLoginCredentials{
			Account:        "test@example.com",
			PowerStationID: "d035a6b1-a99a-46cf-84a0-ddd5730d7f5c",
		},
		SemsLoginResponse: &auth.SemsLoginResponse{},
	}

	tests := []struct {
		name         string
		startTime    string
		endTime      string
		wantTimes    int
		wantError    bool
	}{
		{
			name:      "valid 5 day range",
			startTime: "2025-12-11 00:00",
			endTime:   "2025-12-15 00:00",
			wantTimes: 5,
			wantError: false,
		},
		{
			name:      "valid 1 day range",
			startTime: "2025-12-11 00:00",
			endTime:   "2025-12-11 00:00",
			wantTimes: 1,
			wantError: false,
		},
		{
			name:      "invalid start format",
			startTime: "2025-12-11",
			endTime:   "2025-12-15 00:00",
			wantTimes: 0,
			wantError: true,
		},
		{
			name:      "invalid end format",
			startTime: "2025-12-11 00:00",
			endTime:   "bad-date",
			wantTimes: 0,
			wantError: true,
		},
	}

	targets := defaultTargets()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := BuildRequest(loginInfo, tt.startTime, tt.endTime, "55000DSC22CW3619", "Test User", "123 Test St", targets)

			if (err != nil) != tt.wantError {
				t.Errorf("BuildRequest() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !tt.wantError {
				if req.Times != tt.wantTimes {
					t.Errorf("BuildRequest() Times = %d, want %d", req.Times, tt.wantTimes)
				}
				if req.QryTimeStart != tt.startTime {
					t.Errorf("BuildRequest() QryTimeStart = %s, want %s", req.QryTimeStart, tt.startTime)
				}
				if req.QryTimeEnd != tt.endTime {
					t.Errorf("BuildRequest() QryTimeEnd = %s, want %s", req.QryTimeEnd, tt.endTime)
				}
				if len(req.PwsHistorys) != 1 {
					t.Errorf("BuildRequest() PwsHistorys count = %d, want 1", len(req.PwsHistorys))
				}
				if req.PwsHistorys[0].ID != "d035a6b1-a99a-46cf-84a0-ddd5730d7f5c" {
					t.Errorf("BuildRequest() PwsHistorys[0].ID = %s, want powerstation ID", req.PwsHistorys[0].ID)
				}
				if len(req.Targets) != 14 {
					t.Errorf("BuildRequest() Targets count = %d, want 14", len(req.Targets))
				}
			}
		})
	}
}

func TestDefaultTargets(t *testing.T) {
	targets := defaultTargets()

	if len(targets) != 14 {
		t.Errorf("defaultTargets() returned %d targets, want 14", len(targets))
	}

	// Verify first and last target
	if targets[0].TargetKey != "Vpv1" || targets[0].TargetIndex != 1 {
		t.Errorf("First target should be Vpv1:1, got %s:%d", targets[0].TargetKey, targets[0].TargetIndex)
	}
	if targets[13].TargetKey != "PF" || targets[13].TargetIndex != 407 {
		t.Errorf("Last target should be PF:407, got %s:%d", targets[13].TargetKey, targets[13].TargetIndex)
	}
}
