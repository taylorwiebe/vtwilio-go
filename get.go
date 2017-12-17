package vtwilio

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/twiebe-va/vtwilio-go/internal"
)

// GetMessage gets a message by it's sid
func (v *VTwilio) GetMessage(messageSID string) (*Message, error) {
	if messageSID == "" {
		return nil, fmt.Errorf("must contain a message SID")
	}
	return v.getMessage(messageSID)
}

func (v *VTwilio) getMessage(messageSID string) (*Message, error) {
	urlStr := fmt.Sprintf("%v%v%v/%v.json", v.baseAPI, v.accountSID, messageAPI, messageSID)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	internal.SetUpRequest(req, v.accountSID, v.authToken)
	return handleMessage(req)
}

func handleMessage(req *http.Request) (*Message, error) {
	bodyBytes, err := internal.HandleRequest(req)
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
