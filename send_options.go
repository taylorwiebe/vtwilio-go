package vtwilio

type sendConfiguration struct {
	MediaURL string
}

// SendOption is an option for messages being sent
type SendOption func(c *sendConfiguration)

// MediaURL add an image to a twilio message
func MediaURL(url string) SendOption {
	return func(c *sendConfiguration) {
		c.MediaURL = url
	}
}
