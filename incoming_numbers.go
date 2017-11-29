package vtwilio

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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

	urlStr := fmt.Sprintf("%s%s%s", v.baseAPI, v.accountSID, incomingPhoneNumbersAPI)
	if sid != "" {
		urlStr = fmt.Sprintf("%s/%s", urlStr, sid)
	}
	urlStr = fmt.Sprintf("%s.json", urlStr)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(en))
	if err != nil {
		return nil, err
	}
	setUpRequest(req, v.accountSID, v.authToken)
	return handleIncomingPhoneNumbers(req)
}

// ReleaseNumber "deletes" a number. This number could be used by someone else.
func (v *VTwilio) ReleaseNumber(sid string) error {
	return errors.New("not implemented")
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
