package vtwilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Interface for VTwilio
type Interface interface {
	SendMessage(message string, to string) (*Message, error)
	GetMessages(pageSize int, page int) (*Response, error)
}

const (
	baseAPI    = "https://api.twilio.com/2010-04-01/Accounts/"
	messageAPI = "/Messages.json"
)

// VTwilio is a structure holding details about a twilio account
type VTwilio struct {
	accountSID   string
	authToken    string
	twilioNumber string
}

// Response is a response from a get
type Response struct {
	FirstPageURI    string     `json:"first_page_uri"`
	End             int        `json:"end"`
	PreviousPageURI string     `json:"previous_page_uri"`
	Messages        []*Message `json:"messages"`
}

// Message is a response from Twilio
type Message struct {
	SID                 string `json:"sid"`
	DateCreated         string `json:"date_created"`
	DateUpdated         string `json:"date_updated"`
	DateSent            string `json:"date_sent"`
	AccountSID          string `json:"account_sid"`
	To                  string `json:"to"`
	From                string `json:"from"`
	MessagingServiceSID string `json:"messaging_service_sid"`
	Body                string `json:"body"`
	Status              string `json:"status"`
	NumSegments         string `json:"num_segments"`
	NumMedia            string `json:"num_media"`
	Direction           string `json:"direction"`
	APIVersion          string `json:"api_version"`
	Price               string `json:"price"`
	PriceUnit           string `json:"price_unit"`
	ErrorCode           string `json:"error_code"`
	ErrorMessage        string `json:"error_message"`
	URI                 string `json:"uri"`
	SubresourceURIs     struct {
		Media string `json:"media"`
	} `json:"subresource_uris"`
}

// NewVTwilio returns a new NewVTwilio instance
func NewVTwilio(accountSID, authToken, twilioNumber string) *VTwilio {
	return &VTwilio{accountSID, authToken, twilioNumber}
}

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

	urlStr := strings.Join([]string{baseAPI, v.accountSID, messageAPI}, "")
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(en))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(v.accountSID, v.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Error, with status code: %v", resp.Status)
	}

	var data Message
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMessages returns a list if the messages you have sent
func (v *VTwilio) GetMessages(pageSize int, page int) (*Response, error) {
	return nil, fmt.Errorf("not implemented")
}
