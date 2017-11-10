package vtwilio

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// IncomingPhoneNumber purchase an incoming phone number
func (v *VTwilio) IncomingPhoneNumber(number string, opts ...IncomingPhoneNumberOption) (*IncomingPhoneNumber, error) {
	if err := validateNumber(number); err != nil {
		return nil, err
	}
	config := &incomingNumberConfiguration{
		PhoneNumber: number,
	}
	for _, o := range opts {
		o(config)
	}

	en, err := buildPostValues(config)
	if err != nil {
		return nil, err
	}

	urlStr := fmt.Sprintf("%s%s%s.json", baseAPI, v.accountSID, incomingPhoneNumbersAPI)
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(en))
	if err != nil {
		return nil, err
	}
	setUpRequest(req, v.accountSID, v.authToken)
	return handleIncomingPhoneNumbers(req)
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

func buildPostValues(c *incomingNumberConfiguration) (string, error) {
	v := reflect.Indirect(reflect.ValueOf(c))
	values := url.Values{}
	for i := 0; i < v.NumField(); i++ {
		val, ok := v.Field(i).Interface().(string)
		if !ok {
			return "", fmt.Errorf("unable to build values")
		}
		if val == "" {
			continue
		}
		values.Set(v.Type().Field(i).Tag.Get(tag), val)
	}
	return values.Encode(), nil
}
