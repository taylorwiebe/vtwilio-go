package vtwilio

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// Country type
type Country int

const (
	// US United States Code
	US Country = iota
	// CA Canada code
	CA
)

// ToCode converts a country to a country code
func (c Country) ToCode() (string, error) {
	switch c {
	case US:
		return "/US", nil
	case CA:
		return "/CA", nil
	}
	return "", fmt.Errorf("invalid country")
}

// AvailablePhoneNumbers finds an available phone number
func (v *VTwilio) AvailablePhoneNumbers(country Country, opts ...AvailableOption) (*AvailablePhoneNumbers, error) {
	config := &availableConfiguration{}
	for _, o := range opts {
		o(config)
	}

	c, err := country.ToCode()
	if err != nil {
		return nil, err
	}
	val, err := buildValues(config)
	if err != nil {
		return nil, err
	}

	urlStr := fmt.Sprintf("%s%s%s%s%s.json?%s", baseAPI, v.accountSID, availablePhoneNumbersAPI, c, local, val)
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
