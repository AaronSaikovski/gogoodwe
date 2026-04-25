package utils

import (
	"regexp"
	"strings"
)

var (
	emailPattern          = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	powerstationIDPattern = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
)

// CheckValidEmail checks if the loginEmail string is in a valid email format.
func CheckValidEmail(loginEmail string) bool {
	return emailPattern.MatchString(loginEmail)
}

// CheckValidPowerstationID checks for valid powerstation id (case-insensitive UUID).
func CheckValidPowerstationID(powerStationID string) bool {
	return powerstationIDPattern.MatchString(strings.ToLower(powerStationID))
}
