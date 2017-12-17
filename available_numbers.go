package vtwilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/twiebe-va/vtwilio-go/internal"
)

// AvailablePhoneNumbers finds an available phone number
func (v *VTwilio) AvailablePhoneNumbers(countryCode string, opts ...AvailableOption) (*AvailablePhoneNumbers, error) {
	config := &availableConfiguration{}
	for _, o := range opts {
		o(config)
	}

	val := buildValues(config)

	urlStr := fmt.Sprintf("%s%s%s/%s%s.json?%s", v.baseAPI, v.accountSID, availablePhoneNumbersAPI, countryCode, local, val)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	internal.SetUpRequest(req, v.accountSID, v.authToken)
	return handleAvailability(req)
}

func buildValues(c *availableConfiguration) string {
	v := reflect.Indirect(reflect.ValueOf(c))
	values := []string{}
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).Interface().(string)
		if val == "" {
			continue
		}
		values = append(values, fmt.Sprintf("%s=%s", v.Type().Field(i).Name, val))
	}
	return strings.Join(values, "&")
}

func handleAvailability(req *http.Request) (*AvailablePhoneNumbers, error) {
	bodyBytes, err := internal.HandleRequest(req)
	if err != nil {
		return nil, err
	}

	var data AvailablePhoneNumbers
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
