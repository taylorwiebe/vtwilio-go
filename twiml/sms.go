package twiml

import "encoding/xml"

// SMS TwiML data
type SMS struct {
	XMLName        xml.Name `xml:"Sms"`
	To             string   `xml:"to,attr"`
	From           string   `xml:"from,attr"`
	Action         string   `xml:"action,attr"`
	Method         Method   `xml:"method,attr"`
	StatusCallback string   `xml:"statusCallback,attr"`
}
