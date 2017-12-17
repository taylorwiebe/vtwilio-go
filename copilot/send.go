package copilot

import (
	"github.com/twiebe-va/vtwilio-go"
)

// SendMessage sends a message to a Twilio service
func (c *Copilot) SendMessage(message, to string, opts ...vtwilio.SendOption) (*vtwilio.Message, error) {
	o := []vtwilio.SendOption{vtwilio.ServiceSID(c.serviceSID)}
	o = append(o, opts...)
	resp, err := c.twilio.SendMessage(message, to, o...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
