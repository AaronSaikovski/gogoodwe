package utils

import (
	"testing"
)

func TestParseOutput(t *testing.T) {
	t.Run("valid JSON", func(t *testing.T) {
		input := []byte(`{"key":"value","num":42}`)
		result, err := ParseOutput(input)
		if err != nil {
			t.Fatalf("ParseOutput() unexpected error: %v", err)
		}
		if result == nil {
			t.Fatal("ParseOutput() returned nil")
		}
		if result.GetStringBytes("key") == nil {
			t.Error("ParseOutput() result missing 'key' field")
		}
	})

	t.Run("valid JSON array", func(t *testing.T) {
		input := []byte(`[1,2,3]`)
		result, err := ParseOutput(input)
		if err != nil {
			t.Fatalf("ParseOutput() unexpected error: %v", err)
		}
		if result == nil {
			t.Fatal("ParseOutput() returned nil")
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		input := []byte(`not valid json`)
		_, err := ParseOutput(input)
		if err == nil {
			t.Error("ParseOutput() expected error for invalid JSON")
		}
	})

	t.Run("empty input", func(t *testing.T) {
		input := []byte(``)
		_, err := ParseOutput(input)
		if err == nil {
			t.Error("ParseOutput() expected error for empty input")
		}
	})
}
