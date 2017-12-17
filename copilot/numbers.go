package copilot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/twiebe-va/vtwilio-go/internal"
)

// Number response structure
type Number struct {
	SID          string   `json:"sid"`
	AccountSID   string   `json:"account_sid"`
	ServiceSID   string   `json:"service_sid"`
	DateCreated  string   `json:"date_created"`
	DateUpdated  string   `json:"date_updated"`
	PhoneNumber  string   `json:"phone_number"`
	CountryCode  string   `json:"country_code"`
	Capabilities []string `json:"capabilities"`
	URL          string   `json:"url"`
}

// ListMeta is meta data for the list
type ListMeta struct {
	Page            int    `json:"page"`
	PageSize        int    `json:"page_size"`
	FirstPageURL    string `json:"first_page_url"`
	PreviousPageURL string `json:"previous_page_url"`
	NextPageURL     string `json:"next_page_url"`
	Key             string `json:"key"`
	URL             string `json:"url"`
}

// ServiceNumbers list
type ServiceNumbers struct {
	Meta         ListMeta `json:"meta"`
	PhoneNumbers []Number `json:"phone_numbers"`
}

// AddPhoneNumber adds a phone number to a service
func (c *Copilot) AddPhoneNumber(phoneNumberSID string) (*Number, error) {
	if c.serviceSID == "" {
		return nil, fmt.Errorf("service SID not set")
	}

	values := url.Values{}
	values.Set("PhoneNumberSid", phoneNumberSID)

	en := values.Encode()

	urlStr := fmt.Sprintf("%v/%v/PhoneNumbers", c.baseAPI, c.serviceSID)
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(en))
	if err != nil {
		return nil, err
	}

	internal.SetUpRequest(req, c.twilio.GetAccountSID(), c.twilio.GetAccountAuthToken())
	return handleAddPhoneNumber(req)
}

func handleAddPhoneNumber(req *http.Request) (*Number, error) {
	bodyBytes, err := internal.HandleRequest(req)
	if err != nil {
		return nil, err
	}

	var data Number
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ListPhoneNumber lists numbers a service has
func (c *Copilot) ListPhoneNumber(page, pageSize int) (*ServiceNumbers, error) {
	if c.serviceSID == "" {
		return nil, fmt.Errorf("service SID not set")
	}

	if pageSize == 0 {
		pageSize = 50
	}

	urlStr := fmt.Sprintf("%v/%v/PhoneNumbers?PageSize=%d&Page=%d", c.baseAPI, c.serviceSID, pageSize, page)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	internal.SetUpRequest(req, c.twilio.GetAccountSID(), c.twilio.GetAccountAuthToken())
	return handleListPhoneNumber(req)
}

func handleListPhoneNumber(req *http.Request) (*ServiceNumbers, error) {
	bodyBytes, err := internal.HandleRequest(req)
	if err != nil {
		return nil, err
	}

	var data ServiceNumbers
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// RemoveNumber removes a phone number from a service. This does not remove it from the account.
func (c *Copilot) RemoveNumber(numberSID string) error {
	if numberSID == "" {
		return fmt.Errorf("invalid sid")
	}

	urlStr := fmt.Sprintf("%v/%v/PhoneNumbers/%v", c.baseAPI, c.serviceSID, numberSID)
	req, err := http.NewRequest("DELETE", urlStr, nil)
	if err != nil {
		return err
	}
	internal.SetUpRequest(req, c.twilio.GetAccountSID(), c.twilio.GetAccountAuthToken())
	return internal.GenericHandler(req)
}
