package vtwilio

import (
	"time"
)

type optionConfiguration struct {
	To       string
	Date     time.Time
	PageSize int
	Page     int
}

// Option is a list option
type Option func(o *optionConfiguration)

// To sets who the message is sent to
func To(to string) Option {
	return func(r *optionConfiguration) {
		r.To = to
	}
}

// Date sets the time range
func Date(date time.Time) Option {
	return func(r *optionConfiguration) {
		r.Date = date
	}
}

// PageSize sets the page size
func PageSize(pageSize int) Option {
	return func(r *optionConfiguration) {
		r.PageSize = pageSize
	}
}

// Page sets the page number
func Page(page int) Option {
	return func(r *optionConfiguration) {
		r.Page = page
	}
}
