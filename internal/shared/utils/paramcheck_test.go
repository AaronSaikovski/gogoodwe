package utils

import "testing"

func TestCheckValidEmail(t *testing.T) {
	tests := []struct {
		name  string
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValidEmail(tt.email); got != tt.want {
				t.Errorf("CheckValidEmail(%q) = %v, want %v", tt.email, got, tt.want)
			}
		})
	}
}

func TestCheckValidPowerstationID(t *testing.T) {
	tests := []struct {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValidPowerstationID(tt.id); got != tt.want {
				t.Errorf("CheckValidPowerstationID(%q) = %v, want %v", tt.id, got, tt.want)
			}
		})
	}
}
