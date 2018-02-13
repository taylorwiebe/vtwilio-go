package vtwilio

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
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
	from := v.twilioNumber
	if config.From != "" {
		from = config.From
	}

	values := url.Values{}
	values.Set("To", to)
	values.Set("From", from)
	values.Set("Body", message)
	if config.MediaURL != "" {
		values.Set("MediaUrl", config.MediaURL)
	}

	if config.CallbackURL != "" && config.CallbackMethod != "" {
		values.Set("statusCallback", config.CallbackURL)
		values.Set("StatusCallbackMethod", config.CallbackMethod.String())
	}

	en := values.Encode()
	urlStr := fmt.Sprintf("%s%s%s.json", v.baseAPI, v.accountSID, messageAPI)
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(en))
	if err != nil {
		return nil, err
	}
	setUpRequest(req, v.accountSID, v.authToken)
	return handleMessage(req)
}
