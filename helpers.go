package vtwilio

import "time"

// ToTime convers a twilio api time response to time.Time
func ToTime(timeStr string) (time.Time, error) {
	t, err := time.Parse("Mon, 2 Jan 2006 15:04:05 +0000", timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
