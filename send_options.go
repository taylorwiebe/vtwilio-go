package vtwilio

type sendConfiguration struct {
	MediaURL       string
	From           string
	CallbackURL    string
	CallbackMethod Method
}

// SendOption is an option for messages being sent
type SendOption func(c *sendConfiguration)

// MediaURL add an image to a twilio message
func MediaURL(url string) SendOption {
	return func(c *sendConfiguration) {
		c.MediaURL = url
	}
}

// FromNumber is the phone number that the text message will be sent from
// This will override the phone number set on the client for the current text message
func FromNumber(number string) SendOption {
	return func(c *sendConfiguration) {
		c.From = number
	}
}

// Callback adds a callback url with the method that Twilio should call this endpoint
func Callback(url string, method Method) SendOption {
	return func(c *sendConfiguration) {
		c.CallbackURL = url
		c.CallbackMethod = method
	}
}
