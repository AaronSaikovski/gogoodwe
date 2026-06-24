<<<<<<< HEAD:internal/shared/utils/paramcheck_test.go
package utils

import "testing"
=======
package utils_test

import (
	"testing"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)
>>>>>>> origin/main:cmd/gogoodwe/utils/paramcheck_test.go

func TestCheckValidEmail(t *testing.T) {
	tests := []struct {
		name  string
<<<<<<< HEAD:internal/shared/utils/paramcheck_test.go
		email string
		want  bool
	}{
		{"valid email", "user@example.com", true},
		{"valid email with dots", "first.last@example.com", true},
		{"valid email with plus", "user+tag@example.com", true},
		{"valid email with percent", "user%tag@example.com", true},
		{"valid email with subdomain", "user@mail.example.com", true},
		{"valid email with long TLD", "user@example.museum", true},
		{"empty string", "", false},
		{"no at sign", "userexample.com", false},
		{"no domain", "user@", false},
		{"no user", "@example.com", false},
		{"double at", "user@@example.com", false},
		{"no TLD", "user@example", false},
		{"single char TLD", "user@example.c", false},
		{"spaces", "user @example.com", false},
		{"just text", "notanemail", false},
=======
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
>>>>>>> origin/main:cmd/gogoodwe/utils/paramcheck_test.go
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
<<<<<<< HEAD:internal/shared/utils/paramcheck_test.go
			if got := CheckValidEmail(tt.email); got != tt.want {
				t.Errorf("CheckValidEmail(%q) = %v, want %v", tt.email, got, tt.want)
=======
			got := utils.CheckValidEmail(tt.input)
			if got != tt.want {
				t.Errorf("CheckValidEmail(%q) = %v, want %v", tt.input, got, tt.want)
>>>>>>> origin/main:cmd/gogoodwe/utils/paramcheck_test.go
			}
		})
	}
}

func TestCheckValidPowerstationID(t *testing.T) {
	tests := []struct {
<<<<<<< HEAD:internal/shared/utils/paramcheck_test.go
		name string
		id   string
		want bool
	}{
		{"valid lowercase UUID", "12345678-1234-1234-1234-123456789abc", true},
		{"valid uppercase UUID", "12345678-1234-1234-1234-123456789ABC", true},
		{"valid mixed case UUID", "12345678-ABCD-1234-abcd-123456789ABC", true},
		{"all zeros", "00000000-0000-0000-0000-000000000000", true},
		{"all f's lowercase", "ffffffff-ffff-ffff-ffff-ffffffffffff", true},
		{"all F's uppercase", "FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF", true},
		{"empty string", "", false},
		{"too short", "12345678-1234-1234-1234-12345678", false},
		{"too long", "12345678-1234-1234-1234-123456789abcdef", false},
		{"no dashes", "12345678123412341234123456789abc", false},
		{"extra dash", "12345678-1234-1234-1234-1234-56789abc", false},
		{"invalid hex char", "12345678-1234-1234-1234-12345678xxxx", false},
		{"spaces", "12345678 1234 1234 1234 123456789abc", false},
		{"just text", "not-a-uuid-at-all-nope-nopenope1234", false},
=======
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
>>>>>>> origin/main:cmd/gogoodwe/utils/paramcheck_test.go
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
<<<<<<< HEAD:internal/shared/utils/paramcheck_test.go
			if got := CheckValidPowerstationID(tt.id); got != tt.want {
				t.Errorf("CheckValidPowerstationID(%q) = %v, want %v", tt.id, got, tt.want)
=======
			got := utils.CheckValidPowerstationID(tt.input)
			if got != tt.want {
				t.Errorf("CheckValidPowerstationID(%q) = %v, want %v", tt.input, got, tt.want)
>>>>>>> origin/main:cmd/gogoodwe/utils/paramcheck_test.go
			}
		})
	}
}
