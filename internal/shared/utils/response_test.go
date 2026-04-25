package utils

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestFetchResponseBody(t *testing.T) {
	t.Run("normal response", func(t *testing.T) {
		body := bytes.NewBufferString(`{"status":"ok"}`)
		result, err := FetchResponseBody(body)
		if err != nil {
			t.Fatalf("FetchResponseBody() unexpected error: %v", err)
		}
		if string(result) != `{"status":"ok"}` {
			t.Errorf("FetchResponseBody() = %s, want %s", result, `{"status":"ok"}`)
		}
	})

	t.Run("empty response", func(t *testing.T) {
		body := bytes.NewBufferString("")
		result, err := FetchResponseBody(body)
		if err != nil {
			t.Fatalf("FetchResponseBody() unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("FetchResponseBody() = %s, want empty", result)
		}
	})

	t.Run("oversized response", func(t *testing.T) {
		// Create a reader that returns more than MaxResponseSize
		body := strings.NewReader(strings.Repeat("x", MaxResponseSize+1))
		_, err := FetchResponseBody(body)
		if err == nil {
			t.Error("FetchResponseBody() expected error for oversized response")
		}
	})

	t.Run("read error", func(t *testing.T) {
		body := &errorReader{}
		_, err := FetchResponseBody(body)
		if err == nil {
			t.Error("FetchResponseBody() expected error for read failure")
		}
	})
}

type errorReader struct{}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, io.ErrUnexpectedEOF
}
