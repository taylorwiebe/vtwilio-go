package vtwilio

type sendConfiguration struct {
	ServiceSID string
	MediaURL   string
}

// SendOption is an option for messages being sent
type SendOption func(c *sendConfiguration)

// MediaURL add an image to a twilio message
func MediaURL(url string) SendOption {
	return func(c *sendConfiguration) {
		c.MediaURL = url
	}
}

// ServiceSID is when a service should be used over a specific phone number
func ServiceSID(sid string) SendOption {
	return func(c *sendConfiguration) {
		c.ServiceSID = sid
	}
}
