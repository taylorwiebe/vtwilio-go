package vtwilio

import (
	"strconv"
	"strings"
)

// AvailableOption is an option for available phone numbers
type AvailableOption func(*availableConfiguration)

type availableConfiguration struct {
	NearNumber   string
	NearLatLong  string
	Distance     string
	InPostalCode string
	InLocality   string
	InRegion     string
	InRateCenter string
	InLata       string
}

// NearNumber Twilio Description:
// Given a phone number, find a geographically close number within Distance miles.
// Distance defaults to 25 miles.
func NearNumber(n string) AvailableOption {
	return func(a *availableConfiguration) {
		a.NearNumber = n
	}
}

// NearLatLong Twilio Description:
// Given a latitude/longitude pair lat,long find geographically close numbers within Distance miles.
func NearLatLong(lat, long string) AvailableOption {
	return func(a *availableConfiguration) {
		a.NearLatLong = strings.Join([]string{lat, long}, ",")
	}
}

// Distance Twilio Description:
// Specifies the search radius for a Near- query in miles.
// If not specified this defaults to 25 miles. Maximum searchable distance is 500 miles.
func Distance(d int) AvailableOption {
	return func(a *availableConfiguration) {
		a.Distance = strconv.Itoa(d)
	}
}

// InPostalCode Twilio Description:
// Limit results to a particular postal code.
// Given a phone number, search within the same postal code as that number.
func InPostalCode(p string) AvailableOption {
	return func(a *availableConfiguration) {
		a.InPostalCode = p
	}
}

// InLocality Twilio Description:
// Limit results to a particular locality (i.e. City).
// Given a phone number, search within the same Locality as that number.
func InLocality(l string) AvailableOption {
	return func(a *availableConfiguration) {
		a.InLocality = l
	}
}

// InRegion Twilio Description:
// Limit results to a particular region (i.e. State/Province).
// Given a phone number, search within the same Region as that number.
func InRegion(r string) AvailableOption {
	return func(a *availableConfiguration) {
		a.InRegion = r
	}
}

// InRateCenter Twilio Description:
// Limit results to a specific rate center,
// or given a phone number search within the same rate center as that number.
// Requires InLata to be set as well.
func InRateCenter(r string) AvailableOption {
	return func(a *availableConfiguration) {
		a.InRateCenter = r
	}
}

// InLATA Twilio Description:
// Limit results to a specific Local access and transport area (LATA).
// Given a phone number, search within the same LATA as that number.
func InLATA(l string) AvailableOption {
	return func(a *availableConfiguration) {
		a.InLata = l
	}
}
