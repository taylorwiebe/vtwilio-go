package twiml

import (
	"encoding/xml"
	"errors"
)

// Interface for TwiML
type Interface interface {
	Build()
}

// TwiML response structure
type TwiML struct {
	XMLName     xml.Name `xml:"Response"`
	SayOpt      Say      `xml:"Say,omitempty"`
	MessageOpt  []string `xml:"Message,omitempty"`
	RedirectOpt string   `xml:"Redirect,omitempty"`
	DialOpt     Dial     `xml:"Dial,omitempty"`
	RejectOpt   Reject   `xml:"Reject,omitempty"`
	SMSOpt      SMS      `xml:"Sms,omitempty"`
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
	return nil, errors.New("not implemented")
}

// Say adds a message that will be said to the user
func (t *TwiML) Say(voice Voice, language Language, loop int) *TwiML {
	t.SayOpt = Say{Voice: voice, Language: language, Loop: loop}
	return t
}

// Message sets a message to be sent to the user
func (t *TwiML) Message(messages ...string) *TwiML {
	t.MessageOpt = append(t.MessageOpt, messages...)
	return t
}

// Redirect twilio
func (t *TwiML) Redirect(url string) *TwiML {
	t.RedirectOpt = url
	return t
}

// Dial to a number
func (t *TwiML) Dial(action, method, callerID, number string) *TwiML {
	t.DialOpt = Dial{Action: action, Method: method, CallerID: callerID, Number: number}
	return t
}

// Reject a call
func (t *TwiML) Reject(reason Reason) *TwiML {
	t.RejectOpt = Reject{Reason: reason}
	return t
}

// SMS send an sms message
func (t *TwiML) SMS(to, from, action, statusCallBack string, method Method) *TwiML {
	t.SMSOpt = SMS{To: to, From: from, Action: action, Method: method, StatusCallback: statusCallBack}
	return t
}
