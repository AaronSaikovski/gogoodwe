<<<<<<< HEAD:internal/shared/utils/datetime_test.go
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
=======
package utils_test

import (
	"testing"
	"time"

	"github.com/AaronSaikovski/gogoodwe/pkg/utils"
)

func TestGetDate(t *testing.T) {
	now := time.Now()
	expected := now.Format("2006-01-02")

	got := utils.GetDate()

	if got != expected {
		t.Errorf("GetDate() = %q, want %q", got, expected)
	}

	// Verify format is exactly YYYY-MM-DD (10 characters)
	if len(got) != 10 {
		t.Errorf("GetDate() length = %d, want 10", len(got))
	}

	// Verify it looks like a date (contains hyphens at expected positions)
	if got[4] != '-' || got[7] != '-' {
		t.Errorf("GetDate() format incorrect: %q", got)
>>>>>>> origin/main:pkg/utils/datetime_test.go
	}
}
