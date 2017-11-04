package vtwilio

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SendMessage Sends a twilio message and returns the twilio message SID
func (v *VTwilio) SendMessage(message string, to string) (*Message, error) {
	if message == "" {
		return nil, fmt.Errorf("must contain a message")
	}

	if to == "" {
		return nil, fmt.Errorf("must contain a phone number to send the message to")
	}

	return v.sendMessage(message, to)
}

func (v *VTwilio) sendMessage(message, to string) (*Message, error) {
	values := url.Values{}
	values.Set("To", to)
	values.Set("From", v.twilioNumber)
	values.Set("Body", message)
	en := values.Encode()

	urlStr := fmt.Sprintf("%s%s%s.json", baseAPI, v.accountSID, messageAPI)
	fmt.Println(urlStr)
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(en))
	if err != nil {
		return nil, err
	}
	setUpRequest(req, v.accountSID, v.authToken)
	return handleMessage(req)
}
