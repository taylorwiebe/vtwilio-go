package vtwilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testIncomingPhoneNumberArgs struct {
	number  string
	options []IncomingPhoneNumberOption
	vTwilio *VTwilio
}

func TestIncomingNumbersHappyPath(t *testing.T) {
	expected := &IncomingPhoneNumber{
		SID:                 "sid",
		AccountSID:          "aid",
		FriendlyName:        "caller id",
		PhoneNumber:         "12345678910",
		VoiceURL:            "http://url.com",
		VoiceMethod:         "POST",
		VoiceFallbackURL:    "http://url.com",
		VoiceFallbackMethod: "GET",
		DateCreated:         time.Date(2017, time.January, 01, 01, 01, 01, 01, time.UTC),
		DateUpdated:         time.Date(2017, time.January, 01, 01, 01, 01, 01, time.UTC),
		Capabilities:        Capabilities{SMS: true},
		Beta:                false,
		URI:                 "uri",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if r.Method != "POST" {
			t.Errorf("Expected 'POST' got %v", r.Method)
		}

		bytes, err := json.Marshal(expected)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(bytes)
	}))
	defer ts.Close()

	tests := []struct {
		name          string
		in            testIncomingPhoneNumberArgs
		expected      *IncomingPhoneNumber
		expectedError error
	}{
		{
			name: "valid request",
			in: testIncomingPhoneNumberArgs{
				number:  "+10987654321",
				options: nil,
				vTwilio: &VTwilio{
					accountSID:   "sid",
					authToken:    "token",
					twilioNumber: "+12345678910",
					baseAPI:      fmt.Sprintf("%s/", ts.URL),
				},
			},
			expected:      expected,
			expectedError: nil,
		},
		{
			name: "no phone number",
			in: testIncomingPhoneNumberArgs{
				number:  "",
				options: nil,
				vTwilio: &VTwilio{
					accountSID:   "sid",
					authToken:    "token",
					twilioNumber: "+12345678910",
					baseAPI:      fmt.Sprintf("%s/", ts.URL),
				},
			},
			expected:      nil,
			expectedError: fmt.Errorf("phone number is required"),
		},
		{
			name: "fails without a +",
			in: testIncomingPhoneNumberArgs{
				number:  "12345678910",
				options: nil,
				vTwilio: &VTwilio{
					accountSID:   "sid",
					authToken:    "token",
					twilioNumber: "+12345678910",
					baseAPI:      fmt.Sprintf("%s/", ts.URL),
				},
			},
			expected:      nil,
			expectedError: fmt.Errorf("phone number must begin with +"),
		},
		{
			name: "fails with dashes in number",
			in: testIncomingPhoneNumberArgs{
				number:  "+1-234-567-8910",
				options: nil,
				vTwilio: &VTwilio{
					accountSID:   "sid",
					authToken:    "token",
					twilioNumber: "+12345678910",
					baseAPI:      fmt.Sprintf("%s/", ts.URL),
				},
			},
			expected:      nil,
			expectedError: fmt.Errorf("phone number must be numbers"),
		},
		{
			name: "fails with more than 15 digets",
			in: testIncomingPhoneNumberArgs{
				number:  "+1234567891015141",
				options: nil,
				vTwilio: &VTwilio{
					accountSID:   "sid",
					authToken:    "token",
					twilioNumber: "+12345678910",
					baseAPI:      fmt.Sprintf("%s/", ts.URL),
				},
			},
			expected:      nil,
			expectedError: fmt.Errorf("number can only contain 15 digits"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.in.vTwilio.IncomingPhoneNumber(tt.in.number, tt.in.options...)
			assert.Equal(t, tt.expected, actual)
			if err == nil && tt.expectedError != nil {
				t.Error("expected error, got nil")
			} else if err != nil && tt.expectedError == nil {
				t.Errorf("did not expect an error got: %v", err)
			} else if err != nil && tt.expectedError != nil {
				assert.Equal(t, tt.expectedError, err)
			}
		})
	}
}

func TestIncomingNumbersRequests(t *testing.T) {
	tests := []struct {
		name         string
		in           []IncomingPhoneNumberOption
		expectedPath string
		expectedBody string
	}{
		{
			name:         "no options",
			in:           []IncomingPhoneNumberOption{},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910",
		},
		{
			name:         "area code option",
			in:           []IncomingPhoneNumberOption{AreaCode("90210")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "AreaCode=90210&PhoneNumber=%2B12345678910",
		},
		{
			name:         "valid api version",
			in:           []IncomingPhoneNumberOption{APIVersion("2010-04-01")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "ApiVersion=2010-04-01&PhoneNumber=%2B12345678910",
		},
		{
			name:         "valid old api version",
			in:           []IncomingPhoneNumberOption{APIVersion("2008-08-01")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "ApiVersion=2008-08-01&PhoneNumber=%2B12345678910",
		},
		{
			name:         "invalid api version",
			in:           []IncomingPhoneNumberOption{APIVersion("")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "ApiVersion=2010-04-01&PhoneNumber=%2B12345678910",
		},
		{
			name:         "friendly name",
			in:           []IncomingPhoneNumberOption{FriendlyName("name")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "FriendlyName=name&PhoneNumber=%2B12345678910",
		},
		{
			name:         "voice url",
			in:           []IncomingPhoneNumberOption{VoiceURL("http://url.com")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&VoiceUrl=http%3A%2F%2Furl.com",
		},
		{
			name:         "voice method get",
			in:           []IncomingPhoneNumberOption{VoiceMethod(GET)},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&VoiceMethod=GET",
		},
		{
			name:         "voice method post",
			in:           []IncomingPhoneNumberOption{VoiceMethod(POST)},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&VoiceMethod=POST",
		},
		{
			name:         "voice fallback url",
			in:           []IncomingPhoneNumberOption{VoiceFallBackURL("http://url.com")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&VoiceFallbackUrl=http%3A%2F%2Furl.com",
		},
		{
			name:         "voice method get",
			in:           []IncomingPhoneNumberOption{VoiceFallBackMethod(GET)},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&VoiceFallbackMethod=GET",
		},
		{
			name:         "voice method post",
			in:           []IncomingPhoneNumberOption{VoiceFallBackMethod(POST)},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&VoiceFallbackMethod=POST",
		},
		{
			name:         "status callback url",
			in:           []IncomingPhoneNumberOption{StatusCallback("http://url.com")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&StatusCallback=http%3A%2F%2Furl.com",
		},
		{
			name:         "status callback get",
			in:           []IncomingPhoneNumberOption{StatusCallbackMethod(GET)},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&StatusCallbackMethod=GET",
		},
		{
			name:         "status callback post",
			in:           []IncomingPhoneNumberOption{StatusCallbackMethod(POST)},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&StatusCallbackMethod=POST",
		},
		{
			name:         "voice caller id lookup true",
			in:           []IncomingPhoneNumberOption{VoiceCallerIDLookup(true)},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&VoiceCallerIdLookup=true",
		},
		{
			name:         "voice caller id lookup false",
			in:           []IncomingPhoneNumberOption{VoiceCallerIDLookup(false)},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&VoiceCallerIdLookup=false",
		},
		{
			name:         "voice application sid",
			in:           []IncomingPhoneNumberOption{VoiceApplicationSID("sid")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&VoiceApplicationSid=sid",
		},
		{
			name:         "voice application sid",
			in:           []IncomingPhoneNumberOption{TrunkSID("sid")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&TrunkSid=sid",
		},
		{
			name:         "sms url",
			in:           []IncomingPhoneNumberOption{SMSURL("http://url.com")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&SmsUrl=http%3A%2F%2Furl.com",
		},
		{
			name:         "sms get",
			in:           []IncomingPhoneNumberOption{SMSMethod(GET)},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&SmsMethod=GET",
		},
		{
			name:         "sms post",
			in:           []IncomingPhoneNumberOption{SMSMethod(POST)},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&SmsMethod=POST",
		},
		{
			name:         "sms fallback url",
			in:           []IncomingPhoneNumberOption{SMSFallbackURL("http://url.com")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&SmsFallbackUrl=http%3A%2F%2Furl.com",
		},
		{
			name:         "sms fallback get",
			in:           []IncomingPhoneNumberOption{SMSFallbackMethod(GET)},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&SmsFallbackMethod=GET",
		},
		{
			name:         "sms fallback post",
			in:           []IncomingPhoneNumberOption{SMSFallbackMethod(POST)},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&SmsFallbackMethod=POST",
		},
		{
			name:         "sms application sid",
			in:           []IncomingPhoneNumberOption{SMSApplicationSID("sid")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "PhoneNumber=%2B12345678910&SmsApplicationSid=sid",
		},
		{
			name:         "account sid",
			in:           []IncomingPhoneNumberOption{AccountSID("sid")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "AccountSid=sid&PhoneNumber=%2B12345678910",
		},
		{
			name:         "account sid",
			in:           []IncomingPhoneNumberOption{AddressSID("sid")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "AddressSid=sid&PhoneNumber=%2B12345678910",
		},
		{
			name:         "multiple options",
			in:           []IncomingPhoneNumberOption{AddressSID("sid"), FriendlyName("name")},
			expectedPath: "/sid/IncomingPhoneNumbers.json",
			expectedBody: "AddressSid=sid&FriendlyName=name&PhoneNumber=%2B12345678910",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tt.expectedPath, r.URL.Path)
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					t.Errorf("failed to read body: %v", err)
				}
				assert.Equal(t, tt.expectedBody, string(body))
			}))
			defer ts.Close()

			v := &VTwilio{
				accountSID:   "sid",
				authToken:    "token",
				twilioNumber: "+12345678910",
				baseAPI:      fmt.Sprintf("%s/", ts.URL),
			}
			v.IncomingPhoneNumber("+12345678910", tt.in...)
		})
	}
}

func TestIncomingNumbersBadResponse(t *testing.T) {
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
		actual, err := v.IncomingPhoneNumber("+10987654321")

		assert.Nil(t, actual)
		assert.Equal(t, fmt.Errorf("Error: invalid request"), err)
	})
}

func TestBadRequest(t *testing.T) {
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
		actual, err := v.IncomingPhoneNumber("+10987654321")

		assert.Nil(t, actual)
		assert.Error(t, err)
	})
}

func TestBuildPostValues(t *testing.T) {
	tests := []struct {
		name     string
		in       *incomingNumberConfiguration
		expected string
	}{
		{
			name: "builds values",
			in: &incomingNumberConfiguration{
				PhoneNumber: "+12345678910",
				AreaCode:    "306",
			},
			expected: "AreaCode=306&PhoneNumber=%2B12345678910",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := buildPostValues(tt.in)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
