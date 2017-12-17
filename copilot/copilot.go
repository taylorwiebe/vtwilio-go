package copilot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/twiebe-va/vtwilio-go"
	"github.com/twiebe-va/vtwilio-go/internal"
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

const (
	servicesURL = "https://messaging.twilio.com/v1/Services"
)

// Copilot implementation
type Copilot struct {
	baseAPI string
	twilio  vtwilio.VTwilio
}

// NewCopilot returns a copilot instance
func NewCopilot(t vtwilio.VTwilio) *Copilot {
	return &Copilot{twilio: t, baseAPI: servicesURL}
}

// NewService creates a new Copilot service
func (c *Copilot) NewService(friendlyName, callbackURL string) (*Service, error) {
	values := url.Values{}
	values.Set("FriendlyName", friendlyName)
	values.Set("StatusCallback", callbackURL)

	en := values.Encode()
	urlStr := fmt.Sprintf("%v.json", c.baseAPI)
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(en))
	if err != nil {
		return nil, err
	}
	internal.SetUpRequest(req, c.twilio.GetAccountSID(), c.twilio.GetAccountAuthToken())
	return handleNewService(req)
}

func handleNewService(req *http.Request) (*Service, error) {
	bodyBytes, err := internal.HandleRequest(req)
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