package copilot

import (
	"encoding/json"
	"net/http"

	"github.com/twiebe-va/vtwilio-go"
)

// Links response
type Links struct {
	PhoneNumbers string `json:"phone_numbers"`
	ShortCodes   string `json:"short_codes"`
	AlphaSenders string `json:"alpha_senders"`
}

// Service response
type Service struct {
	AccountID         string         `json:"account_sid"`
	SID               string         `json:"sid"`
	DateCreated       string         `json:"date_created"`
	DateUpdate        string         `json:"date_updated"`
	FriendlyName      string         `json:"friendly_name"`
	InboundRequestURL string         `json:"inbound_request_url"`
	InboudMethod      vtwilio.Method `json:"inbound_method"`
	FallbackURL       string         `json:"fallback_url"`
	FallbackMethod    vtwilio.Method `json:"fallback_method"`
	StatusCallback    string         `json:"status_callback"`
	StickySender      bool           `json:"sticky_sender"`
	MMSConverter      bool           `json:"mms_converter"`
	Links             Links          `json:"links"`
	URL               string         `json:"url"`
}

// Copilot implementation
type Copilot struct {
	twilio vtwilio.VTwilio
}

// NewCopilot returns a copilot instance
func NewCopilot(t vtwilio.VTwilio) *Copilot {
	return &Copilot{twilio: t}
}

// NewService creates a new Copilot service
func (c *Copilot) NewService(friendlyName, callbackURL string) (*Service, error) {
	return nil, nil
}

func handleNewService(req *http.Request) (*Service, error) {
	bodyBytes, err := handleRequest(req)
	if err != nil {
		return nil, err
	}

	var data Service
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
