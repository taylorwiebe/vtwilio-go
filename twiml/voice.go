package twiml

import "encoding/xml"

// Language for voices
type Language string

const (
	// EN English
	EN Language = "en"
	// ENGB English Great Britain
	ENGB Language = "en-gb"
	// ES English
	ES Language = "es"
	// ENAU English Australian
	ENAU Language = "en-AU"
)

// Voice are Twilio voices
type Voice string

const (
	// Man voice
	Man Voice = "man"
	// Woman voice
	Woman Voice = "woman"
	// Alice voice
	Alice Voice = "Alice"
)

// Say message options
type Say struct {
	XMLName  xml.Name `xml:"Say"`
	Voice    Voice    `xml:"voice,attr"`
	Language Language `xml:"language,attr"`
	Loop     int      `xml:"loop,attr"`
}

// Reason for rejection
type Reason string

const (
	// Busy so call is rejected
	Busy Reason = "busy"
	// Rejected default
	Rejected Reason = "rejected"
)

// Reject an incoming call
type Reject struct {
	XMLName xml.Name `xml:"Reject"`
	Reason  Reason   `xml:"reason,attr"`
}

// Dial is the TwiML dial structure
type Dial struct {
	XMLName  xml.Name `xml:"Dial"`
	Action   string   `xml:"action,attr"`
	Method   string   `xml:"method,attr"`
	CallerID string   `xml:"callerId,attr"`
	Number   string   `xml:",omitempty"`
}
