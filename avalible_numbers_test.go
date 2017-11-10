package vtwilio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildValues(t *testing.T) {
	tests := []struct {
		name          string
		in            *availableConfiguration
		expected      string
		expectedError error
	}{
		{
			name: "builds valid value",
			in: &availableConfiguration{
				InPostalCode: "90210",
			},
			expected:      "InPostalCode=90210",
			expectedError: nil,
		},
		{
			name: "builds valid value",
			in: &availableConfiguration{
				NearNumber:   "12345678910",
				NearLatLong:  "34.0928118.3287",
				Distance:     "25",
				InPostalCode: "90210",
				InLocality:   "HOLLYWOOD",
				InRegion:     "CALIFORNIA",
				InRateCenter: "CALIFORNIA",
				InLata:       "",
			},
			expected:      "NearNumber=12345678910&NearLatLong=34.0928118.3287&Distance=25&InPostalCode=90210&InLocality=HOLLYWOOD&InRegion=CALIFORNIA&InRateCenter=CALIFORNIA",
			expectedError: nil,
		},
		{
			name: "removes empty fields",
			in: &availableConfiguration{
				NearNumber:   "",
				NearLatLong:  "34.0928118.3287",
				Distance:     "25",
				InPostalCode: "90210",
				InLocality:   "HOLLYWOOD",
				InRegion:     "CALIFORNIA",
				InRateCenter: "CALIFORNIA",
				InLata:       "",
			},
			expected:      "NearLatLong=34.0928118.3287&Distance=25&InPostalCode=90210&InLocality=HOLLYWOOD&InRegion=CALIFORNIA&InRateCenter=CALIFORNIA",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := buildValues(tt.in)
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
