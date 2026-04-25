package utils

import (
	"testing"
)

func TestProcessRawJSON(t *testing.T) {
	t.Run("valid JSON", func(t *testing.T) {
		err := ProcessRawJSON([]byte(`{"key":"value"}`))
		if err != nil {
			t.Errorf("ProcessRawJSON() unexpected error: %v", err)
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		err := ProcessRawJSON([]byte(`not json`))
		if err == nil {
			t.Error("ProcessRawJSON() expected error for invalid JSON")
		}
	})

	t.Run("empty input", func(t *testing.T) {
		err := ProcessRawJSON([]byte(``))
		if err == nil {
			t.Error("ProcessRawJSON() expected error for empty input")
		}
	})
}
