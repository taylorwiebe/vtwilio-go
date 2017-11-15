package vtwilio

import (
	"fmt"
	"strings"
	"time"
)

// ToTime convers a twilio api time response to time.Time
func ToTime(timeStr string) (time.Time, error) {
	t, err := time.Parse("Mon, 2 Jan 2006 15:04:05 +0000", timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func handleDateRange(date time.Time, opt dateOption) string {
	t := date.Format("2006-01-02T15:04:05.999999-07:00")
	t = strings.Split(t, "T")[0]
	if opt == before {
		return fmt.Sprintf("DateSent<=%v", t)
	} else if opt == after {
		return fmt.Sprintf("DateSent>=%v", t)
	} else if opt == equal {
		return fmt.Sprintf("DateSent=%v", t)
	}
	return ""
}
