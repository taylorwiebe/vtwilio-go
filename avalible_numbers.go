package vtwilio

import "fmt"

// Country type
type Country string

const (
	// US United States Code
	US Country = "US"
	// CA Canada code
	CA Country = "CA"
)

// AvailablePhoneNumbers finds an available phone number
func (v *VTwilio) AvailablePhoneNumbers(country Country, opts ...AvailableOption) (*AvailablePhoneNumbers, error) {
	config := &availableConfiguration{}
	for _, o := range opts {
		o(config)
	}

	urlStr := fmt.Sprintf("%v%v%v.json", baseAPI, v.accountSID, avaliblePhoneNumbersAPI)

	return nil, nil
}

// func (v *VTwilio) sendMessage(message, to string) (*Message, error) {
// 	values := url.Values{}
// 	values.Set("To", to)
// 	values.Set("From", v.twilioNumber)
// 	values.Set("Body", message)
// 	en := values.Encode()

// 	urlStr := fmt.Sprintf("%v%v%v.json", baseAPI, v.accountSID, messageAPI)
// 	req, err := http.NewRequest("POST", urlStr, strings.NewReader(en))
// 	if err != nil {
// 		return nil, err
// 	}
// 	setUpRequest(req, v.accountSID, v.authToken)
// 	return handleMessage(req)
// }
