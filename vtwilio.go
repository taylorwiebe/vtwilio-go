package vtwilio

// Interface for VTwilio
type Interface interface {
	SetPhoneNumber(n string) *VTwilio
	SendMessage(message string, to string) (*Message, error)
	ListMessages(opts ...ListOption) (*List, error)
	GetMessage(messageSID string) (*Message, error)
	AvailablePhoneNumbers(countryCode string, opts ...AvailableOption) (*AvailablePhoneNumbers, error)
}

const (
	baseAPI                  = "https://api.twilio.com/2010-04-01/Accounts/"
	messageAPI               = "/Messages"
	availablePhoneNumbersAPI = "/AvailablePhoneNumbers"
	incomingPhoneNumbersAPI  = "/IncomingPhoneNumbers"
	local                    = "/Local"
)

// VTwilio is a structure holding details about a twilio account
type VTwilio struct {
	accountSID   string
	authToken    string
	twilioNumber string
}

// List is a response from a get
type List struct {
	FirstPageURI    string     `json:"first_page_uri"`
	End             int        `json:"end"`
	PreviousPageURI string     `json:"previous_page_uri"`
	Messages        []*Message `json:"messages"`
}

// Message is a response from Twilio
type Message struct {
	SID                 string `json:"sid"`
	DateCreated         string `json:"date_created"`
	DateUpdated         string `json:"date_updated"`
	DateSent            string `json:"date_sent"`
	AccountSID          string `json:"account_sid"`
	To                  string `json:"to"`
	From                string `json:"from"`
	MessagingServiceSID string `json:"messaging_service_sid"`
	Body                string `json:"body"`
	Status              string `json:"status"`
	NumSegments         string `json:"num_segments"`
	NumMedia            string `json:"num_media"`
	Direction           string `json:"direction"`
	APIVersion          string `json:"api_version"`
	Price               string `json:"price"`
	PriceUnit           string `json:"price_unit"`
	ErrorCode           string `json:"error_code"`
	ErrorMessage        string `json:"error_message"`
	URI                 string `json:"uri"`
	SubresourceURIs     struct {
		Media string `json:"media"`
	} `json:"subresource_uris"`
}

// AvailablePhoneNumbers response form twilio
type AvailablePhoneNumbers struct {
	URI                   string `json:"uri"`
	AvailablePhoneNumbers []struct {
		FriendlyName string `json:"friendly_name"`
		PhoneNumber  string `json:"phone_number"`
		LATA         string `json:"lata"`
		RateCenter   string `json:"rate_center"`
		Latitude     string `json:"latitude"`
		Longitude    string `json:"longitude"`
		Region       string `json:"region"`
		PostalCode   string `json:"postal_code"`
		ISOCountry   string `json:"iso_country"`
		Capabilities struct {
			Voice bool `json:"voice"`
			SMS   bool `json:"SMS"`
			MMS   bool `json:"MMS"`
		} `json:"capabilities"`
		Beta bool `json:"beta"`
	} `json:"available_phone_numbers"`
}

// Option options for vtwilio
type Option func(*VTwilio)

// TwilioNumber sets the twilio number
func TwilioNumber(n string) Option {
	return func(v *VTwilio) {
		v.twilioNumber = n
	}
}

// NewVTwilio returns a new NewVTwilio instance
func NewVTwilio(accountSID, authToken string, opts ...Option) *VTwilio {
	v := &VTwilio{accountSID: accountSID, authToken: authToken}
	for _, o := range opts {
		o(v)
	}
	return v
}

// SetPhoneNumber sets the twilio phone number
func (v *VTwilio) SetPhoneNumber(n string) *VTwilio {
	v.twilioNumber = n
	return v
}
