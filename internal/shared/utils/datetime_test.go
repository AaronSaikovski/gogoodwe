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
}
