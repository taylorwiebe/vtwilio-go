package vtwilio

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/twiebe-va/vtwilio-go/internal"
)

// SendMessage Sends a twilio message and returns the twilio message SID
func (v *VTwilio) SendMessage(message string, to string, opts ...SendOption) (*Message, error) {
	if message == "" {
		return nil, fmt.Errorf("must contain a message")
	}
	if to == "" {
		return nil, fmt.Errorf("must contain a phone number to send the message to")
	}
	config := &sendConfiguration{}
	for _, o := range opts {
		o(config)
	}

	return v.sendMessage(message, to, config)
}

func (v *VTwilio) sendMessage(message, to string, config *sendConfiguration) (*Message, error) {
	values := url.Values{}
	values.Set("To", to)
	values.Set("Body", message)
	if config.ServiceSID != "" {
		values.Set("MessagingServiceSid", config.ServiceSID)
	} else if v.twilioNumber != "" && config.ServiceSID == "" {
		values.Set("From", v.twilioNumber)
	} else {
		return nil, fmt.Errorf("a from phone number is required")
	}
	if config.MediaURL != "" {
		values.Set("MediaUrl", config.MediaURL)
	}

	en := values.Encode()

	urlStr := fmt.Sprintf("%s%s%s.json", v.baseAPI, v.accountSID, messageAPI)
	fmt.Println(urlStr)
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(en))
	if err != nil {
		return nil, err
	}
	internal.SetUpRequest(req, v.accountSID, v.authToken)
	return handleMessage(req)
}
