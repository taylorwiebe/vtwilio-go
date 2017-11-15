package vtwilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendRequest(t *testing.T) {
	tests := []struct {
		name         string
		opts         []SendOption
		message, to  string
		expectedPath string
		expectedBody string
	}{
		{
			name:         "no options",
			opts:         []SendOption{},
			message:      "Text message",
			to:           "+09876543210",
			expectedPath: "/sid/Messages.json",
			expectedBody: "Body=Text+message&From=%2B12345678910&To=%2B09876543210",
		},
		{
			name:         "media url",
			opts:         []SendOption{MediaURL("http://url.com")},
			message:      "Text message",
			to:           "+09876543210",
			expectedPath: "/sid/Messages.json",
			expectedBody: "Body=Text+message&From=%2B12345678910&MediaUrl=http%3A%2F%2Furl.com&To=%2B09876543210",
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
			v.SendMessage(tt.message, tt.to, tt.opts...)
		})
	}
}

func TestSendValidation(t *testing.T) {
	tests := []struct {
		name, message, to string
		expectedError     error
	}{
		{"no message", "", "+12345678910", fmt.Errorf("must contain a message")},
		{"no message", "Message", "", fmt.Errorf("must contain a phone number to send the message to")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VTwilio{
				accountSID:   "sid",
				authToken:    "token",
				twilioNumber: "+12345678910",
			}
			_, err := v.SendMessage(tt.message, tt.to)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestSendBadResponse(t *testing.T) {
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
		actual, err := v.SendMessage("message", "+12345678910")

		assert.Nil(t, actual)
		assert.Equal(t, fmt.Errorf("Error: invalid request"), err)
	})
}

func TestSendBadRequest(t *testing.T) {
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
		actual, err := v.SendMessage("message", "+12345678910")

		assert.Nil(t, actual)
		assert.Error(t, err)
	})
}
