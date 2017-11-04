package vtwilio

import (
	"fmt"
	"net/http"
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
	urlStr := fmt.Sprintf("%v%v%v.json?PageSize=%v&Page=%v", baseAPI, v.accountSID, messageAPI, config.PageSize, config.Page)
	if !config.Date.IsZero() {
		urlStr = fmt.Sprintf("%v&%v", urlStr, handleDateRange(config.Date, config.DateRange))
	}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	setUpRequest(req, v.accountSID, v.authToken)
	return handleListMessages(req)
}
