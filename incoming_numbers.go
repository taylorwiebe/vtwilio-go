package vtwilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/twiebe-va/vtwilio-go/internal"
)

// IncomingPhoneNumber purchase an incoming phone number
func (v *VTwilio) IncomingPhoneNumber(number string, opts ...IncomingPhoneNumberOption) (*IncomingPhoneNumber, error) {
	return v.incomingPhoneNumber(number, "", opts...)
}

// UpdateIncomingPhoneNumber updates an existing phone numbers info
func (v *VTwilio) UpdateIncomingPhoneNumber(number, sid string, opts ...IncomingPhoneNumberOption) (*IncomingPhoneNumber, error) {
	return v.incomingPhoneNumber(number, sid, opts...)
}

func (v *VTwilio) incomingPhoneNumber(number, sid string, opts ...IncomingPhoneNumberOption) (*IncomingPhoneNumber, error) {
	if err := validateNumber(number); err != nil {
		return nil, err
	}
	config := &incomingNumberConfiguration{
		PhoneNumber: number,
	}
	for _, o := range opts {
		o(config)
	}
	en := buildPostValues(config)

	urlStr := buildIncomingPhoneNumber(v.baseAPI, v.accountSID, sid)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(en))
	if err != nil {
		return nil, err
	}
	internal.SetUpRequest(req, v.accountSID, v.authToken)
	return handleIncomingPhoneNumbers(req)
}

func buildIncomingPhoneNumber(api, accountSID, sid string) string {
	urlStr := fmt.Sprintf("%s%s%s", api, accountSID, incomingPhoneNumbersAPI)
	if sid != "" {
		urlStr = fmt.Sprintf("%s/%s", urlStr, sid)
	}
	urlStr = fmt.Sprintf("%s.json", urlStr)
	return urlStr
}

// ReleaseNumber "deletes" a number. This number could be used by someone else.
func (v *VTwilio) ReleaseNumber(sid string) error {
	if sid == "" {
		return fmt.Errorf("invalid sid")
	}

	urlStr := buildIncomingPhoneNumber(v.baseAPI, v.accountSID, sid)
	req, err := http.NewRequest("DELETE", urlStr, nil)
	if err != nil {
		return err
	}
	internal.SetUpRequest(req, v.accountSID, v.authToken)
	return internal.GenericHandler(req)
}

func validateNumber(n string) error {
	if n == "" {
		return fmt.Errorf("phone number is required")
	}
	if n[0] != '+' {
		return fmt.Errorf("phone number must begin with +")
	}
	if strings.Contains(n, "-") {
		return fmt.Errorf("phone number must be numbers")
	}
	if len(n) > 16 {
		return fmt.Errorf("number can only contain 15 digits")
	}

	return nil
}

func buildPostValues(c *incomingNumberConfiguration) string {
	v := reflect.Indirect(reflect.ValueOf(c))
	values := url.Values{}
	for i := 0; i < v.NumField(); i++ {
		val, _ := v.Field(i).Interface().(string)
		if val == "" {
			continue
		}
		values.Set(v.Type().Field(i).Tag.Get(tag), val)
	}
	return values.Encode()
}

func handleIncomingPhoneNumbers(req *http.Request) (*IncomingPhoneNumber, error) {
	bodyBytes, err := internal.HandleRequest(req)
	if err != nil {
		return nil, err
	}

	var data IncomingPhoneNumber
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
