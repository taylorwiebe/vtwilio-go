package vtwilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMessage(t *testing.T) {
	expected := &Message{
		SID:                 "sid",
		DateCreated:         time.Date(2017, time.January, 01, 01, 01, 01, 01, time.UTC).String(),
		DateUpdated:         time.Date(2017, time.January, 01, 01, 01, 01, 01, time.UTC).String(),
		DateSent:            time.Date(2017, time.January, 01, 01, 01, 01, 01, time.UTC).String(),
		AccountSID:          "account_sid",
		To:                  "+123445678910",
		From:                "+10987654321",
		MessagingServiceSID: "messaging_sid",
		Body:                "message",
		Status:              "200",
		NumSegments:         "1",
		NumMedia:            "0",
		Direction:           "",
		APIVersion:          "2010-04-01",
		Price:               "0.00",
		PriceUnit:           "1",
		ErrorCode:           "",
		ErrorMessage:        "",
		URI:                 "uri",
		SubresourceURIs: Media{
			Media: "media",
		},
	}

	tests := []struct {
		name          string
		in            string
		checkResult   bool
		expected      *Message
		expectedError bool
		failServer    bool
		expectedPath  string
		requestURL    string
	}{
		{
			name:          "missing sid",
			in:            "",
			checkResult:   true,
			expected:      nil,
			expectedError: true,
		},
		{
			name:          "valid request",
			in:            "sid",
			checkResult:   true,
			expected:      expected,
			expectedError: false,
			expectedPath:  "/sid/Messages/sid.json",
		},
		{
			name:          "bad request",
			in:            "sid",
			checkResult:   true,
			failServer:    true,
			expected:      nil,
			expectedError: true,
			expectedPath:  "/sid/Messages/sid.json",
		},
		{
			name:          "bad url",
			in:            "sid",
			checkResult:   true,
			failServer:    true,
			expected:      nil,
			expectedError: true,
			expectedPath:  "/sid/Messages/sid.json",
			requestURL:    "http://192.168.0.%31/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.failServer {
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
					return
				}
				w.WriteHeader(http.StatusOK)
				if r.Method != "GET" {
					t.Errorf("Expected 'GET' got %v", r.Method)
				}

				assert.Equal(t, tt.expectedPath, r.URL.Path)

				bytes, err := json.Marshal(expected)
				if err != nil {
					t.Fatal(err)
				}
				w.Write(bytes)
			}))
			defer ts.Close()

			requestURL := fmt.Sprintf("%s/", ts.URL)
			if tt.requestURL != "" {
				requestURL = tt.requestURL
			}

			v := &VTwilio{
				accountSID:   "sid",
				authToken:    "token",
				twilioNumber: "+12345678910",
				baseAPI:      requestURL,
			}

			actual, err := v.GetMessage(tt.in)

			if tt.checkResult {
				assert.Equal(t, tt.expected, actual)
				if err == nil && tt.expectedError {
					t.Error("expected error got nil")
				} else if err != nil && !tt.expectedError {
					t.Errorf("did not expect error, got: %v", err)
				}
			}
		})
	}
}
