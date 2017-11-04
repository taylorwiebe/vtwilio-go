package vtwilio

// Interface for VTwilio
type Interface interface {
	SendMessage(message string, to string) (*Message, error)
	ListMessages(opts ...Option) (*List, error)
	GetMessage(messageSID string) (*Message, error)
}

const (
	baseAPI    = "https://api.twilio.com/2010-04-01/Accounts/"
	messageAPI = "/Messages"
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

// NewVTwilio returns a new NewVTwilio instance
func NewVTwilio(accountSID, authToken, twilioNumber string) *VTwilio {
	return &VTwilio{accountSID, authToken, twilioNumber}
}
