package copilot

// MessageDeliveryStatus is a helper struct for expected requests from Twilio's message status webhooks
type MessageDeliveryStatus struct {
	SID                 string `json:"sid"`
	DateCreated         string `json:"date_created"`
	DateUpdated         string `json:"date_updated"`
	DateSent            string `json:"date_sent"`
	AccountSID          string `json:"account_sid"`
	To                  string `json:"to"`
	From                string `json:"from"`
	Body                string `json:"body"`
	Status              string `json:"status"`
	MessagingServiceSID string `json:"messaging_service_sid"`
	NumSegments         string `json:"num_segments"`
	NumMedia            string `json:"num_media"`
	Direction           string `json:"direction"`
	APIVersion          string `json:"api_version"`
	Price               string `json:"price"`
	PriceUnit           string `json:"price_unit"`
	ErrorCode           string `json:"error_code"`
	ErrorMessage        string `json:"error_message"`
	URI                 string `json:"uri"`
}
