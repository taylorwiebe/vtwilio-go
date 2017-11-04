package vtwilio

import (
	"fmt"
	"net/http"
)

// ListMessages returns a list if the messages you have sent
func (v *VTwilio) ListMessages(opts ...Option) (*List, error) {
	c := &optionConfiguration{
		PageSize: 10,
		Page:     0,
	}

	for _, o := range opts {
		o(c)
	}

	return v.listMessages(c)
}

func (v *VTwilio) listMessages(config *optionConfiguration) (*List, error) {
	urlStr := fmt.Sprintf("%v%v%v.json?PageSize=%v&Page=%v", baseAPI, v.accountSID, messageAPI, config.PageSize, config.Page)
	if !config.Date.IsZero() {
		urlStr = fmt.Sprintf("%v&%v", urlStr, handleDateRange(config.Date, config.DateRange))
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(v.accountSID, v.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return handleListRequest(req)
}
