package utils

import "testing"

func TestIsWithin7Days(t *testing.T) {
	tests := []struct {
		name    string
		date1   string
		date2   string
		want    bool
		wantErr bool
	}{
		{"same time", "2024-01-01 12:00", "2024-01-01 12:00", true, false},
		{"1 day apart", "2024-01-01 12:00", "2024-01-02 12:00", true, false},
		{"exactly 7 days", "2024-01-01 12:00", "2024-01-08 12:00", true, false},
		{"8 days apart", "2024-01-01 12:00", "2024-01-09 12:00", false, false},
		{"reversed dates within 7 days", "2024-01-05 12:00", "2024-01-01 12:00", true, false},
		{"reversed dates over 7 days", "2024-01-10 12:00", "2024-01-01 12:00", false, false},
		{"invalid first date", "not-a-date", "2024-01-01 12:00", false, true},
		{"invalid second date", "2024-01-01 12:00", "not-a-date", false, true},
		{"both invalid", "bad", "bad", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsWithin7Days(tt.date1, tt.date2)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsWithin7Days() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsWithin7Days(%q, %q) = %v, want %v", tt.date1, tt.date2, got, tt.want)
			}
		})
	}
}
