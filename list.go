package vtwilio

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// ListMessages returns a list if the messages you have sent
func (v *VTwilio) ListMessages(opts ...ListOption) (*List, error) {
	c := &listOptionConfiguration{
		PageSize: 10,
		Page:     0,
	}

	for _, o := range opts {
		o(c)
	}

	return v.listMessages(c)
}

func (v *VTwilio) listMessages(config *listOptionConfiguration) (*List, error) {
	urlStr := fmt.Sprintf("%s%s%s.json?PageSize=%v&Page=%v", v.baseAPI, v.accountSID, messageAPI, config.PageSize, config.Page)
	values := buildListValues(config)
	if values != "" {
		urlStr = fmt.Sprintf("%v&%v", urlStr, values)
	}

	if !config.Date.IsZero() {
		urlStr = fmt.Sprintf("%s&%s", urlStr, handleDateRange(config.Date, config.DateRange))
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	setUpRequest(req, v.accountSID, v.authToken)
	return handleListMessages(req)
}

func buildListValues(c *listOptionConfiguration) string {
	v := reflect.Indirect(reflect.ValueOf(c))
	values := []string{}
	for i := 0; i < v.NumField(); i++ {
		val, ok := v.Field(i).Interface().(string)
		if val == "" || !ok {
			continue
		}
		values = append(values, fmt.Sprintf("%s=%s", v.Type().Field(i).Name, val))
	}
	return strings.Join(values, "&")
}
