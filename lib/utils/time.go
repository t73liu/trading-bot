package utils

import "time"

func GetLastWeekday(now time.Time) time.Time {
	prevDay := now.AddDate(0, 0, -1)
	for prevDay.Weekday() == time.Saturday || prevDay.Weekday() == time.Sunday {
		prevDay = prevDay.AddDate(0, 0, -1)
	}
	return prevDay
}

func GetNYSELocation() *time.Location {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	return location
}

// Assuming location = "America/New_York"
func IsWithinNYSETradingHours(moment time.Time) bool {
	hour, minute, _ := moment.Clock()
	if hour == 9 {
		return minute >= 30
	}
	return hour > 9 && hour < 16
}
