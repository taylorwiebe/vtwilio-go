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
		expectedError error
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ToTime(tt.in)
			assert.Equal(t, tt.expected, actual)
			if err != nil && tt.expectedError == nil {
				t.Errorf("did not expect an error, got: %v", err)
			} else if err == nil && tt.expectedError != nil {
				t.Error("expected an error, got nil")
			} else {
				assert.Equal(t, tt.expectedError, err)
			}
		})
	}
}
