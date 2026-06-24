package utils

import (
	"regexp"
	"testing"
	"time"
)

func TestGetDate(t *testing.T) {
	result := GetDate()

	// Verify format matches YYYY-MM-DD
	pattern := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !pattern.MatchString(result) {
		t.Errorf("GetDate() = %q, want format YYYY-MM-DD", result)
	}

	// Verify it matches today's date
	expected := time.Now().Format("2006-01-02")
	if result != expected {
		t.Errorf("GetDate() = %q, want %q", result, expected)
	}

	// Verify format is exactly YYYY-MM-DD (10 characters)
	if len(result) != 10 {
		t.Errorf("GetDate() length = %d, want 10", len(result))
	}

	// Verify it looks like a date (contains hyphens at expected positions)
	if result[4] != '-' || result[7] != '-' {
		t.Errorf("GetDate() format incorrect: %q", result)
	}
}
