package utils

import "time"

// Get the current date and time Format the date as yyyy-mm-dd
func GetDate() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02")
}
