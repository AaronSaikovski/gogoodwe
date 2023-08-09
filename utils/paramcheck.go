package utils

import (
	"regexp"
)

// CheckValidEmail - check for a valid email address input
func CheckValidEmail(loginEmail string) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return pattern.MatchString(loginEmail)
}

// CheckValidPowerstationID - check for valid powerstation id
func CheckValidPowerstationID(powerStationID string) bool {
	pattern := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	return pattern.MatchString(powerStationID)
}
