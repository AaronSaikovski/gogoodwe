package utils_test

import (
	"bytes"
	"testing"

	"github.com/AaronSaikovski/gogoodwe/pkg/utils"
	"github.com/valyala/fastjson"
)

func TestParseOutput(t *testing.T) {
	tests := []struct {
		name     string
		jsonData []byte
		wantErr  bool
	}{
		{
			name:     "valid JSON object",
			jsonData: []byte(`{"key":"value","num":42}`),
			wantErr:  false,
		},
		{
			name:     "valid JSON array",
			jsonData: []byte(`[1,2,3]`),
			wantErr:  false,
		},
		{
			name:     "valid JSON string",
			jsonData: []byte(`"hello"`),
			wantErr:  false,
		},
		{
			name:     "empty JSON",
			jsonData: []byte(``),
			wantErr:  true,
		},
		{
			name:     "invalid JSON",
			jsonData: []byte(`{invalid}`),
			wantErr:  true,
		},
		{
			name:     "valid JSON number",
			jsonData: []byte(`42`),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := utils.ParseOutput(tt.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result == nil {
				t.Error("ParseOutput() returned nil for valid JSON")
			}
			if !tt.wantErr {
				// Verify it's a *fastjson.Value
				var _ *fastjson.Value = result
			}
		})
	}
}

func TestFetchResponseBody(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantLen    int
		wantErr    bool
		wantErrSub string
	}{
		{
			name:    "normal response",
			input:   `{"key":"value"}`,
			wantLen: 15,
			wantErr: false,
		},
		{
			name:    "empty response",
			input:   ``,
			wantLen: 0,
			wantErr: false,
		},
		{
			name:       "oversized response",
			input:      string(make([]byte, 11*1024*1024)),
			wantErr:    true,
			wantErrSub: "too large",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := utils.FetchResponseBody(bytes.NewReader([]byte(tt.input)))
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchResponseBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(result) != tt.wantLen {
				t.Errorf("FetchResponseBody() len = %d, want %d", len(result), tt.wantLen)
			}
			if err != nil && tt.wantErrSub != "" && !bytes.Contains([]byte(err.Error()), []byte(tt.wantErrSub)) {
				t.Errorf("FetchResponseBody() error = %v, want substring %q", err, tt.wantErrSub)
			}
		})
	}
}
