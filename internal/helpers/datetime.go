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
