package utils

import (
	"regexp"
)

// CheckValidEmail checks if the loginEmail string is in a valid email format.
//
// It takes a loginEmail string as a parameter and returns a boolean value.
func CheckValidEmail(loginEmail string) bool {
	var pattern = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return pattern.MatchString(loginEmail)
}

// CheckValidPowerstationID check for valid powerstation id.
//
// powerStationID: a string representing the powerstation ID.
// Returns a boolean indicating if the powerstation ID is valid.
func CheckValidPowerstationID(powerStationID string) bool {
	var pattern = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	return pattern.MatchString(powerStationID)
}
