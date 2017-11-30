package twiml

import (
	"encoding/xml"
)

// Interface for TwiML
type Interface interface {
	Build()
}

// TwiML response structure
type TwiML struct {
	XMLName     xml.Name `xml:"Response"`
	SayOpt      *Say     `xml:"Say,omitempty"`
	MessageOpt  []string `xml:"Message,omitempty"`
	RedirectOpt string   `xml:"Redirect,omitempty"`
	DialOpt     *Dial    `xml:"Dial,omitempty"`
	RejectOpt   *Reject  `xml:"Reject,omitempty"`
	SMSOpt      *SMS     `xml:"Sms,omitempty"`
}

// Method for http
type Method string

const (
	// POST method
	POST Method = "POST"
	// GET method
	GET Method = "GET"
)

// NewTwiML returns a new instance of TwiML
func NewTwiML() *TwiML {
	return &TwiML{}
}

// Build a TwiML response
func (t *TwiML) Build() ([]byte, error) {
	output, err := xml.MarshalIndent(t, "", "	")
	if err != nil {
		return nil, err
	}
	return output, nil
}
