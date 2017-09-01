package vtwilio

import (
	"time"
)

type dateOption int

const (
	_ dateOption = iota
	before
	equal
	after
)

type optionConfiguration struct {
	To        string
	Date      time.Time
	DateRange dateOption
	PageSize  int
	Page      int
}

// Option is a list option
type Option func(o *optionConfiguration)

// To sets who the message is sent to
func To(to string) Option {
	return func(r *optionConfiguration) {
		r.To = to
	}
}

// OnDate get the messages for a specific ay
func OnDate(date time.Time) Option {
	return func(r *optionConfiguration) {
		r.Date = date
		r.DateRange = equal
	}
}

// OnAndBeforeDate get the messages on and before the specified day
func OnAndBeforeDate(date time.Time) Option {
	return func(r *optionConfiguration) {
		r.Date = date
		r.DateRange = before
	}
}

// OnAndAfterDate get the messages on and after the specified day
func OnAndAfterDate(date time.Time) Option {
	return func(r *optionConfiguration) {
		r.Date = date
		r.DateRange = after
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
