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
	Voice    Voice    `xml:"voice,attr,omitempty"`
	Language Language `xml:"language,attr,omitempty"`
	Loop     int      `xml:"loop,attr,omitempty"`
	Value    string   `xml:",chardata"`
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
	Reason  Reason   `xml:"reason,attr,omitempty"`
}

// Dial is the TwiML dial structure
type Dial struct {
	XMLName  xml.Name `xml:"Dial"`
	Action   string   `xml:"action,attr,omitempty"`
	Method   string   `xml:"method,attr,omitempty"`
	CallerID string   `xml:"callerId,attr,omitempty"`
	Number   string   `xml:",omitempty"`
}

// SayOption option for what a phone call will say
type SayOption func(s *Say)

// SayVoice to take when dial happens
func SayVoice(v Voice) SayOption {
	return func(s *Say) {
		s.Voice = v
	}
}

// SayLanguage voice language
func SayLanguage(l Language) SayOption {
	return func(s *Say) {
		s.Language = l
	}
}

// SayLoop number of times the message loops
func SayLoop(l int) SayOption {
	return func(s *Say) {
		s.Loop = l
	}
}

// Say adds a message that will be said to the user
func (t *TwiML) Say(message string, opts ...SayOption) *TwiML {
	s := &Say{Value: message}
	for _, o := range opts {
		o(s)
	}
	t.SayOpt = s
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

// DialOption option for a phone call
type DialOption func(d *Dial)

// DialAction to take when dial happens
func DialAction(a string) DialOption {
	return func(d *Dial) {
		d.Action = a
	}
}

// DialMethod method of the dial callback
func DialMethod(m string) DialOption {
	return func(d *Dial) {
		d.Method = m
	}
}

// DialCallerID what number to show
func DialCallerID(c string) DialOption {
	return func(d *Dial) {
		d.CallerID = c
	}
}

// DialNumber number to dial
func DialNumber(n string) DialOption {
	return func(d *Dial) {
		d.Number = n
	}
}

// Dial to a number
func (t *TwiML) Dial(opts ...DialOption) *TwiML {
	d := &Dial{}
	for _, o := range opts {
		o(d)
	}
	t.DialOpt = d
	return t
}

// RejectOption option for rejecting a call
type RejectOption func(r *Reject)

// RejectReason to reject a call
func RejectReason(reason Reason) RejectOption {
	return func(r *Reject) {
		r.Reason = reason
	}
}

// Reject a call
func (t *TwiML) Reject(opts ...RejectOption) *TwiML {
	r := &Reject{}
	for _, o := range opts {
		o(r)
	}

	t.RejectOpt = r
	return t
}
