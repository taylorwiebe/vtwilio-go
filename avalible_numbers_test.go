package vtwilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO test expected value

func TestAvailableNumbers(t *testing.T) {
	tests := []struct {
		name          string
		in            []AvailableOption
		expectedPath  string
		expectedQuery string
	}{
		{
			name:         "no options",
			in:           []AvailableOption{},
			expectedPath: "/sid/AvailablePhoneNumbers/US/Local.json",
		},
		{
			name:          "near number",
			in:            []AvailableOption{NearNumber("12345678910")},
			expectedPath:  "/sid/AvailablePhoneNumbers/US/Local.json",
			expectedQuery: "NearNumber=12345678910",
		},
		{
			name:          "near lat long",
			in:            []AvailableOption{NearLatLong("34.0928", "118.3287")},
			expectedPath:  "/sid/AvailablePhoneNumbers/US/Local.json",
			expectedQuery: "NearLatLong=34.0928,118.3287",
		},
		{
			name:          "distance",
			in:            []AvailableOption{Distance(25)},
			expectedPath:  "/sid/AvailablePhoneNumbers/US/Local.json",
			expectedQuery: "Distance=25",
		},
		{
			name:          "in postal code",
			in:            []AvailableOption{InPostalCode("90210")},
			expectedPath:  "/sid/AvailablePhoneNumbers/US/Local.json",
			expectedQuery: "InPostalCode=90210",
		},
		{
			name:          "in locality",
			in:            []AvailableOption{InLocality("Hollywood")},
			expectedPath:  "/sid/AvailablePhoneNumbers/US/Local.json",
			expectedQuery: "InLocality=Hollywood",
		},
		{
			name:          "in region",
			in:            []AvailableOption{InRegion("CALIFORNIA")},
			expectedPath:  "/sid/AvailablePhoneNumbers/US/Local.json",
			expectedQuery: "InRegion=CALIFORNIA",
		},
		{
			name:          "in rate center",
			in:            []AvailableOption{InRateCenter("rate")},
			expectedPath:  "/sid/AvailablePhoneNumbers/US/Local.json",
			expectedQuery: "InRateCenter=rate",
		},
		{
			name:          "in LATA",
			in:            []AvailableOption{InLATA("lata")},
			expectedPath:  "/sid/AvailablePhoneNumbers/US/Local.json",
			expectedQuery: "InLata=lata",
		},
		{
			name:          "multiple options",
			in:            []AvailableOption{InLATA("lata"), Distance(25), InRegion("CALIFORNIA")},
			expectedPath:  "/sid/AvailablePhoneNumbers/US/Local.json",
			expectedQuery: "Distance=25&InRegion=CALIFORNIA&InLata=lata",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tt.expectedPath, r.URL.Path)
				assert.Equal(t, tt.expectedQuery, r.URL.RawQuery)
			}))
			defer ts.Close()

			v := &VTwilio{
				accountSID:   "sid",
				authToken:    "token",
				twilioNumber: "+12345678910",
				baseAPI:      fmt.Sprintf("%s/", ts.URL),
			}
			v.AvailablePhoneNumbers("US", tt.in...)
		})
	}
}

func TestBuildValues(t *testing.T) {
	tests := []struct {
		name     string
		in       *availableConfiguration
		expected string
	}{
		{
			name: "builds valid value",
			in: &availableConfiguration{
				InPostalCode: "90210",
			},
			expected: "InPostalCode=90210",
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
			expected: "NearNumber=12345678910&NearLatLong=34.0928118.3287&Distance=25&InPostalCode=90210&InLocality=HOLLYWOOD&InRegion=CALIFORNIA&InRateCenter=CALIFORNIA",
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
			expected: "NearLatLong=34.0928118.3287&Distance=25&InPostalCode=90210&InLocality=HOLLYWOOD&InRegion=CALIFORNIA&InRateCenter=CALIFORNIA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := buildValues(tt.in)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestAvaliblePhoneNumberBadResponse(t *testing.T) {
	t.Run("bad request", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			resp := struct {
				Status  int
				Message string
			}{
				Status:  400,
				Message: "invalid request",
			}

			w.WriteHeader(http.StatusBadRequest)
			bytes, err := json.Marshal(&resp)
			if err != nil {
				t.Error(err)
			}
			w.Write(bytes)
		}))
		defer ts.Close()

		v := &VTwilio{
			accountSID:   "sid",
			authToken:    "token",
			twilioNumber: "+12345678910",
			baseAPI:      fmt.Sprintf("%s/", ts.URL),
		}
		actual, err := v.AvailablePhoneNumbers("US")

		assert.Nil(t, actual)
		assert.Equal(t, fmt.Errorf("Error: invalid request"), err)
	})
}

func TestAvaliblePhoneNumberBadRequest(t *testing.T) {
	t.Run("bad url", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			resp := struct {
				Status  int
				Message string
			}{
				Status:  400,
				Message: "invalid request",
			}

			w.WriteHeader(http.StatusBadRequest)
			bytes, err := json.Marshal(&resp)
			if err != nil {
				t.Error(err)
			}
			w.Write(bytes)
		}))
		defer ts.Close()

		v := &VTwilio{
			accountSID:   "sid",
			authToken:    "token",
			twilioNumber: "+12345678910",
			baseAPI:      "http://192.168.0.%31/",
		}
		actual, err := v.AvailablePhoneNumbers("US")

		assert.Nil(t, actual)
		assert.Error(t, err)
	})
}
