package twiml

import "encoding/xml"

// SMS TwiML data
type SMS struct {
	XMLName        xml.Name `xml:"Sms,omitempty"`
	To             string   `xml:"to,attr,omitempty"`
	From           string   `xml:"from,attr,omitempty"`
	Action         string   `xml:"action,attr,omitempty"`
	Method         Method   `xml:"method,attr,omitempty"`
	StatusCallback string   `xml:"statusCallback,attr,omitempty"`
	Value          string   `xml:",chardata"`
}

// SMS send an sms message
func (t *TwiML) SMS(message string, opts ...SMSOption) *TwiML {
	sms := &SMS{Value: message}
	for _, o := range opts {
		o(sms)
	}

	t.SMSOpt = sms
	return t
}

// SMSOption options for sms
type SMSOption func(s *SMS)

// SMSTo the number the message is to
func SMSTo(t string) SMSOption {
	return func(s *SMS) {
		s.To = t
	}
}

// SMSFrom the number the message is from
func SMSFrom(f string) SMSOption {
	return func(s *SMS) {
		s.From = f
	}
}

// SMSAction to take for the sms
func SMSAction(a string) SMSOption {
	return func(s *SMS) {
		s.Action = a
	}
}

// SMSStatusCallBack call back url
func SMSStatusCallBack(c string) SMSOption {
	return func(s *SMS) {
		s.StatusCallback = c
	}
}

// SMSMethod method of the callback
func SMSMethod(m Method) SMSOption {
	return func(s *SMS) {
		s.Method = m
	}
}
