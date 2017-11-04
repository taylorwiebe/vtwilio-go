package vtwilio

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// AvailablePhoneNumbers finds an available phone number
func (v *VTwilio) AvailablePhoneNumbers(countryCode string, opts ...AvailableOption) (*AvailablePhoneNumbers, error) {
	config := &availableConfiguration{}
	for _, o := range opts {
		o(config)
	}

	val, err := buildValues(config)
	if err != nil {
		return nil, err
	}

	urlStr := fmt.Sprintf("%s%s%s/%s%s.json?%s", baseAPI, v.accountSID, availablePhoneNumbersAPI, countryCode, local, val)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	setUpRequest(req, v.accountSID, v.authToken)
	return handleAvailability(req)
}

func buildValues(c *availableConfiguration) (string, error) {
	v := reflect.Indirect(reflect.ValueOf(c))
	values := []string{}
	for i := 0; i < v.NumField(); i++ {
		val, ok := v.Field(i).Interface().(string)
		if !ok {
			return "", fmt.Errorf("unable to build values")
		}
		if val == "" {
			continue
		}
		values = append(values, fmt.Sprintf("%s=%s", v.Type().Field(i).Name, val))
	}
	return strings.Join(values, "&"), nil
}
