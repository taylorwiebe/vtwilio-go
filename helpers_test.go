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
