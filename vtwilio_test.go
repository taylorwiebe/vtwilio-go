package vtwilio

import "testing"
import "github.com/stretchr/testify/assert"

type in struct {
	accountSID string
	authToken  string
	opts       []Option
	number     string
}

func TestVTwilio(t *testing.T) {
	tests := []struct {
		name     string
		in       in
		expected *VTwilio
	}{
		{
			name:     "empty vtwilio options",
			in:       in{accountSID: "sid", authToken: "token", opts: []Option{}},
			expected: &VTwilio{accountSID: "sid", authToken: "token", twilioNumber: "", baseAPI: baseAPI},
		},
		{
			name:     "set number",
			in:       in{accountSID: "sid", authToken: "token", opts: []Option{TwilioNumber("12345678910")}},
			expected: &VTwilio{accountSID: "sid", authToken: "token", twilioNumber: "12345678910", baseAPI: baseAPI},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewVTwilio(tt.in.accountSID, tt.in.authToken, tt.in.opts...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
func TestSetPhoneNumber(t *testing.T) {
	tests := []struct {
		name     string
		in       in
		expected *VTwilio
	}{
		{
			name: "empty vtwilio options",
			in:   in{accountSID: "sid", authToken: "token", number: "12345678910"},
			expected: &VTwilio{
				accountSID:   "sid",
				authToken:    "token",
				twilioNumber: "12345678910",
				baseAPI:      baseAPI,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewVTwilio(tt.in.accountSID, tt.in.authToken).SetPhoneNumber(tt.in.number)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
