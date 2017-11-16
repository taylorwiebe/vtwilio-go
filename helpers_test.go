package vtwilio

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToTime(t *testing.T) {
	tests := []struct {
		name          string
		in            string
		expected      time.Time
		expectedError bool
	}{
		{
			name:          "valid time string",
			in:            "Thu, 31 Aug 2017 01:10:56 +0000",
			expected:      time.Date(2017, time.August, 31, 1, 10, 56, 0, time.UTC),
			expectedError: false,
		},
		{
			name:          "invalid time",
			in:            "August 31 2017 01:10:56 +0000",
			expected:      time.Time{},
			expectedError: true,
		},
		{
			name:          "no time",
			in:            "",
			expected:      time.Time{},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ToTime(tt.in)
			assert.Equal(t, tt.expected, actual)
			if err != nil && !tt.expectedError {
				t.Errorf("did not expect an error, got: %v", err)
			} else if err == nil && tt.expectedError {
				t.Error("expected an error, got nil")
			}
		})
	}
}

func TestHandleDateRange(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		option   dateOption
		expected string
	}{
		{
			name:     "before",
			date:     time.Date(2017, time.August, 31, 1, 10, 56, 0, time.UTC),
			option:   before,
			expected: "DateSent<=2017-08-31",
		},
		{
			name:     "after",
			date:     time.Date(2017, time.August, 31, 1, 10, 56, 0, time.UTC),
			option:   after,
			expected: "DateSent>=2017-08-31",
		},
		{
			name:     "equal",
			date:     time.Date(2017, time.August, 31, 1, 10, 56, 0, time.UTC),
			option:   equal,
			expected: "DateSent=2017-08-31",
		},
		{
			name:     "bad option",
			date:     time.Date(2017, time.August, 31, 1, 10, 56, 0, time.UTC),
			option:   -1,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := handleDateRange(tt.date, tt.option)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
