package utils

import "time"

func GetLastWeekday(now time.Time) time.Time {
	prevDay := now.AddDate(0, 0, -1)
	for prevDay.Weekday() == time.Saturday || prevDay.Weekday() == time.Sunday {
		prevDay = prevDay.AddDate(0, 0, -1)
	}
	return prevDay
}
