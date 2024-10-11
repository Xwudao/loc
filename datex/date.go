package datex

import (
	"time"
)

// GetTodayRemainingSec get the remaining seconds of the day
func GetTodayRemainingSec() int {
	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	remainingSeconds := int(endOfDay.Sub(now).Seconds())
	return remainingSeconds
}
