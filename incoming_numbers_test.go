package vtwilio

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateNumber(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected error
	}{
		{
			name:     "valid phone number",
			in:       "+12345678910",
			expected: nil,
		},
		{
			name:     "fails without +",
			in:       "12345678910",
			expected: fmt.Errorf("phone number must begin with +"),
		},
		{
			name:     "fails with dashes",
			in:       "+1-234-567-8910",
			expected: fmt.Errorf("phone number must be numbers"),
		},
		{
			name:     "fails with more than 15 digets",
			in:       "+1234567891015141",
			expected: fmt.Errorf("number can only contain 15 digits"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateNumber(tt.in)
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestBuildPostValues(t *testing.T) {
	tests := []struct {
		name          string
		in            *incomingNumberConfiguration
		expected      string
		expectedError error
	}{
		{
			name: "builds values",
			in: &incomingNumberConfiguration{
				PhoneNumber: "+12345678910",
				AreaCode:    "306",
			},
			expected:      "AreaCode=306&PhoneNumber=%2B12345678910",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := buildPostValues(tt.in)
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
