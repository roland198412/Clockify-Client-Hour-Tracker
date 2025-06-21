package utils

import (
	"fmt"
	"time"
)

func FormatSecondsToHHMMSS(seconds int64) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	remainingSeconds := seconds % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, remainingSeconds)
}

func HoursToSeconds(hours int) int64 {
	return int64(hours * 60 * 60)
}

func TotalWeekdaysBetweenTwoDates(startDate, endDate time.Time) int {
	// Normalize time to midnight
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, endDate.Location())

	// Swap if startDate is after endDate
	if startDate.After(endDate) {
		startDate, endDate = endDate, startDate
	}

	count := 0
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		switch d.Weekday() {
		case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
			count++
		}
	}

	return count
}
