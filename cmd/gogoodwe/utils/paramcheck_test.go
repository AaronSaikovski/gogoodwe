package utils_test

import (
	"testing"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

func TestCheckValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "valid email", input: "user@example.com", want: true},
		{name: "valid email with dots", input: "user.name@example.com", want: true},
		{name: "valid email with plus", input: "user+tag@example.com", want: true},
		{name: "valid email with underscores", input: "user_name@example.com", want: true},
		{name: "valid email with hyphens", input: "user-name@example-domain.com", want: true},
		{name: "valid email numeric domain", input: "user@123.com", want: true},
		{name: "missing @", input: "userexample.com", want: false},
		{name: "missing domain", input: "user@", want: false},
		{name: "missing local part", input: "@example.com", want: false},
		{name: "no TLD", input: "user@example", want: false},
		{name: "empty string", input: "", want: false},
		{name: "spaces", input: "user @example.com", want: false},
		{name: "two @", input: "user@@example.com", want: false},
		{name: "trailing dot TLD", input: "user@example.", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.CheckValidEmail(tt.input)
			if got != tt.want {
				t.Errorf("CheckValidEmail(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCheckValidPowerstationID(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "valid lowercase", input: "12345678-abcd-efab-1234-123456789abc", want: true},
		{name: "valid all lowercase", input: "12345678-abcd-efab-1234-123456789abc", want: true},
		{name: "missing dashes", input: "12345678abcdefghij1234123456789abc", want: false},
		{name: "wrong length", input: "1234567-abcd-efab-1234-123456789abc", want: false},
		{name: "extra dash", input: "123456789-abcd-efab-1234-123456789abc", want: false},
		{name: "empty string", input: "", want: false},
		{name: "spaces", input: "12345678-abcd-efab-1234-123456789abc ", want: false},
		{name: "special chars", input: "12345678-abcd-efab-1234-123456789abc!", want: false},
		{name: "non-hex chars", input: "g2345678-abcd-efab-1234-123456789abc", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.CheckValidPowerstationID(tt.input)
			if got != tt.want {
				t.Errorf("CheckValidPowerstationID(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
