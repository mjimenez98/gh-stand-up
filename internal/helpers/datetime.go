// Package helpers provides utility functions.
package helpers

import (
	"time"
)

// GetYesterdayDate returns the date of yesterday in the format "YYYY-MM-DD".
func GetYesterdayDate() string {
	currentTime := time.Now()
	return currentTime.AddDate(0, 0, -1).Format("2006-01-02")
}

// GetYesterdayDateTime returns the date and time of yesterday at 00:00:00 in the local timezone.
func GetYesterdayDateTime() time.Time {
	currentTime := time.Now()
	currentTime = currentTime.AddDate(0, 0, -1)
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
}
