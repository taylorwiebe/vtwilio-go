package vtwilio

// Interface for VTwilio
type Interface interface {
	SetPhoneNumber(n string) *VTwilio
	SendMessage(message string, to string, opts ...SendOption) (*Message, error)
	ListMessages(opts ...ListOption) (*List, error)
	GetMessage(messageSID string) (*Message, error)
	AvailablePhoneNumbers(countryCode string, opts ...AvailableOption) (*AvailablePhoneNumbers, error)
	IncomingPhoneNumber(number string, opts ...IncomingPhoneNumberOption) error
}

const (
	baseAPI                  = "https://api.twilio.com/2010-04-01/Accounts/"
	messageAPI               = "/Messages"
	availablePhoneNumbersAPI = "/AvailablePhoneNumbers"
	incomingPhoneNumbersAPI  = "/IncomingPhoneNumbers"
	local                    = "/Local"
	tag                      = "vtwilio"
)

// VTwilio is a structure holding details about a twilio account
type VTwilio struct {
	accountSID   string
	authToken    string
	twilioNumber string
	baseAPI      string
}

// List is a response from a get
type List struct {
	FirstPageURI    string     `json:"first_page_uri"`
	End             int        `json:"end"`
	PreviousPageURI string     `json:"previous_page_uri"`
	Messages        []*Message `json:"messages"`
}

// Media contains media data
type Media struct {
	Media string `json:"media"`
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
	SubresourceURIs     Media  `json:"subresource_uris"`
}

// Capabilities structure
type Capabilities struct {
	Voice bool `json:"voice"`
	SMS   bool `json:"SMS"`
	MMS   bool `json:"MMS"`
}

// AvailablePhoneNumberData is the data for an available phone number
type AvailablePhoneNumberData struct {
	FriendlyName string       `json:"friendly_name"`
	PhoneNumber  string       `json:"phone_number"`
	LATA         string       `json:"lata"`
	RateCenter   string       `json:"rate_center"`
	Latitude     string       `json:"latitude"`
	Longitude    string       `json:"longitude"`
	Region       string       `json:"region"`
	PostalCode   string       `json:"postal_code"`
	ISOCountry   string       `json:"iso_country"`
	Capabilities Capabilities `json:"capabilities"`
	Beta         bool         `json:"beta"`
}

// AvailablePhoneNumbers response form twilio
type AvailablePhoneNumbers struct {
	URI                  string                     `json:"uri"`
	AvailablePhoneNumber []AvailablePhoneNumberData `json:"available_phone_numbers"`
}

// IncomingPhoneNumber data from twilio
type IncomingPhoneNumber struct {
	SID                 string       `json:"sid"`
	AccountSID          string       `json:"account_sid"`
	FriendlyName        string       `json:"friendly_name"`
	PhoneNumber         string       `json:"phone_number"`
	VoiceURL            string       `json:"voice_url"`
	VoiceMethod         string       `json:"voice_method"`
	VoiceFallbackURL    string       `json:"voice_fallback_url"`
	VoiceFallbackMethod string       `json:"voice_fallback_method"`
	DateCreated         string       `json:"date_created"`
	DateUpdated         string       `json:"date_updated"`
	Capabilities        Capabilities `json:"capabilities"`
	Beta                bool         `json:"beta"`
	URI                 string       `json:"uri"`
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
	setDefaults(v)
	return v
}

func setDefaults(v *VTwilio) {
	v.baseAPI = baseAPI
}

// SetPhoneNumber sets the twilio phone number
func (v *VTwilio) SetPhoneNumber(n string) *VTwilio {
	v.twilioNumber = n
	return v
}
