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

type listOptionConfiguration struct {
	To        string
	From      string
	Date      time.Time
	DateRange dateOption
	PageSize  int
	Page      int
}

// ListOption is a list option
type ListOption func(o *listOptionConfiguration)

// To sets who the message is sent to
func To(to string) ListOption {
	return func(r *listOptionConfiguration) {
		r.To = to
	}
}

// From sets who the message was from
func From(from string) ListOption {
	return func(r *listOptionConfiguration) {
		r.From = from
	}
}

// OnDate get the messages for a specific ay
func OnDate(date time.Time) ListOption {
	return func(r *listOptionConfiguration) {
		r.Date = date
		r.DateRange = equal
	}
}

// OnAndBeforeDate get the messages on and before the specified day
func OnAndBeforeDate(date time.Time) ListOption {
	return func(r *listOptionConfiguration) {
		r.Date = date
		r.DateRange = before
	}
}

// OnAndAfterDate get the messages on and after the specified day
func OnAndAfterDate(date time.Time) ListOption {
	return func(r *listOptionConfiguration) {
		r.Date = date
		r.DateRange = after
	}
}

// PageSize sets the page size
func PageSize(pageSize int) ListOption {
	return func(r *listOptionConfiguration) {
		r.PageSize = pageSize
	}
}

// Page sets the page number
func Page(page int) ListOption {
	return func(r *listOptionConfiguration) {
		r.Page = page
	}
}
