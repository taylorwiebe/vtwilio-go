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

func TestListMessages(t *testing.T) {
	message := &Message{
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

	expected := &List{
		FirstPageURI:    "http://pageuri.com",
		End:             10,
		PreviousPageURI: "http://pageuri.com/prev",
		Messages:        []*Message{message, message},
	}

	tests := []struct {
		name          string
		in            []ListOption
		expected      *List
		expectedError bool
		expectedPath  string
		expectedQuery string
		failServer    bool
		checkResult   bool
		requestURL    string
	}{
		{
			name:          "no options",
			in:            []ListOption{},
			expectedPath:  "/sid/Messages.json",
			expectedQuery: "PageSize=10&Page=0",
		},
		{
			name:          "to",
			in:            []ListOption{To("+12345678910")},
			expectedPath:  "/sid/Messages.json",
			expectedQuery: "PageSize=10&Page=0",
		},
		{
			name:          "on date",
			in:            []ListOption{OnDate(time.Date(2017, time.January, 01, 01, 01, 01, 01, time.UTC))},
			expectedPath:  "/sid/Messages.json",
			expectedQuery: "PageSize=10&Page=0&DateSent=2017-01-01",
		},
		{
			name:          "on and before date",
			in:            []ListOption{OnAndBeforeDate(time.Date(2017, time.January, 01, 01, 01, 01, 01, time.UTC))},
			expectedPath:  "/sid/Messages.json",
			expectedQuery: "PageSize=10&Page=0&DateSent<=2017-01-01",
		},
		{
			name:          "on and after date",
			in:            []ListOption{OnAndAfterDate(time.Date(2017, time.January, 01, 01, 01, 01, 01, time.UTC))},
			expectedPath:  "/sid/Messages.json",
			expectedQuery: "PageSize=10&Page=0&DateSent>=2017-01-01",
		},
		{
			name:          "page size",
			in:            []ListOption{PageSize(20)},
			expectedPath:  "/sid/Messages.json",
			expectedQuery: "PageSize=20&Page=0",
		},
		{
			name:          "page size",
			in:            []ListOption{Page(2)},
			expectedPath:  "/sid/Messages.json",
			expectedQuery: "PageSize=10&Page=2",
		},
		{
			name:          "valid request",
			in:            []ListOption{},
			expected:      expected,
			expectedError: false,
			expectedPath:  "/sid/Messages.json",
			expectedQuery: "PageSize=10&Page=0",
			checkResult:   true,
		},
		{
			name:          "bad request url",
			in:            []ListOption{},
			expectedPath:  "/sid/Messages.json",
			expectedQuery: "PageSize=10&Page=0",
			expected:      nil,
			checkResult:   true,
			expectedError: true,
			requestURL:    "http://192.168.0.%31/",
		},
		{
			name:          "bad request",
			in:            []ListOption{},
			expectedPath:  "/sid/Messages.json",
			expected:      nil,
			checkResult:   true,
			expectedError: true,
			failServer:    true,
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
				assert.Equal(t, tt.expectedQuery, r.URL.RawQuery)

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

			actual, err := v.ListMessages(tt.in...)

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
