package vtwilio

type incomingNumberConfiguration struct {
	PhoneNumber          string `vtwilio:"PhoneNumber"`
	AreaCode             string `vtwilio:"AreaCode"`
	FriendlyName         string `vtwilio:"FriendlyName"`
	VoiceURL             string `vtwilio:"VoiceUrl"`
	VoiceMethod          string `vtwilio:"FriendlyName"`
	VoiceFallbackURL     string `vtwilio:"VoiceFallbackUrl"`
	VoiceFallbackMethod  string `vtwilio:"VoiceFallbackMethod"`
	StatusCallback       string `vtwilio:"StatusCallback"`
	StatusCallbackMethod string `vtwilio:"StatusCallbackMethod"`
	VoiceCallerIDLookup  bool   `vtwilio:"VoiceCallerIdLookup"`
	VoiceApplicationSID  string `vtwilio:"VoiceApplicationSid"`
	TrunkSID             string `vtwilio:"TrunkSid"`
	SMSURL               string `vtwilio:"SmsUrl"`
	SMSMethod            string `vtwilio:"SmsMethod"`
	SMSFallbackURL       string `vtwilio:"SmsFallbackUrl"`
	SMSFallbackMethod    string `vtwilio:"SmsFallbackMethod"`
	SMSApplicationSID    string `vtwilio:"SmsApplicationSid"`
	AddressSID           string `vtwilio:"AddressSid"`
	APIVersion           string `vtwilio:"ApiVersion"`
	AccountSID           string `vtwilio:"AccountSid"`
}

// IncomingPhoneNumberOption options for an incoming phone number purchase
type IncomingPhoneNumberOption func(*incomingNumberConfiguration)

// AreaCode Twilio description:
// The desired area code for your new incoming phone number.
// Any three digit, US or Canada area code is valid. Twilio will provision a random phone number within
// this area code for you. You must include either this or a PhoneNumber parameter to have your POST succeed. (US and Canada only)
func AreaCode(a string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.AreaCode = a
	}
}

// APIVersion Twilio description:
// Calls to this phone number will start a new TwiML session with this API version. Either 2010-04-01 or 2008-08-01.
func APIVersion(v string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.APIVersion = v
	}
}

// FriendlyName Twilio description:
// A human readable description of the new incoming phone number. Maximum 64 characters. Defaults to a formatted version of the number.
func FriendlyName(n string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.FriendlyName = n
	}
}

// VoiceURL Twilio description:
// The URL Twilio will request when this phone number receives a call. The VoiceURL will no longer be used if
//  a VoiceApplicationSid or a TrunkSid is set.
func VoiceURL(url string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.VoiceURL = url
	}
}

// VoiceMethod Twilio description:
// The HTTP method Twilio will use when requesting the above Url. Either GET or POST.
func VoiceMethod(m Method) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.VoiceMethod = m.String()
	}
}

// VoiceFallBackURL Twilio description:
// A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by VoiceUrl.
func VoiceFallBackURL(url string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.VoiceFallbackURL = url
	}
}

// VoiceFallBackMethod Twilio description:
// The HTTP method that should be used to request the VoiceFallbackUrl. Either GET or POST.
func VoiceFallBackMethod(m Method) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.VoiceFallbackMethod = m.String()
	}
}

// StatusCallback Twilio description:
// The URL that Twilio will request to pass status parameters (such as call ended) to your application.
func StatusCallback(url string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.StatusCallback = url
	}
}

// StatusCallbackMethod Twilio description:
// The HTTP method that should be used to request the VoiceFallbackUrl. Either GET or POST.
func StatusCallbackMethod(m Method) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.StatusCallbackMethod = m.String()
	}
}

// VoiceCallerIDLookup Twilio description:
// Do a lookup of a caller's name from the CNAM database and post it to your app. Either true or false.
func VoiceCallerIDLookup(v bool) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.VoiceCallerIDLookup = v
	}
}

// VoiceApplicationSID Twilio description:
// The 34 character sid of the application Twilio should use to handle phone calls to this number.
// If a VoiceApplicationSid is present, Twilio will ignore all of the voice urls above and use those set on the application instead.
// Setting a VoiceApplicationSid will automatically delete your TrunkSid and vice versa.
func VoiceApplicationSID(sid string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.VoiceApplicationSID = sid
	}
}

// TrunkSID Twilio description:
// The 34 character sid of the Trunk Twilio should use to handle phone calls to this number.
// If a TrunkSid is present, Twilio will ignore all of the voice urls and voice applications above and use those set on the Trunk.
// Setting a TrunkSid will automatically delete your VoiceApplicationSid and vice versa.
func TrunkSID(sid string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.TrunkSID = sid
	}
}

// SMSURL Twilio description:
// A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by SmsUrl.
func SMSURL(url string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.SMSURL = url
	}
}

// SMSMethod Twilio description:
// The HTTP method that should be used to request the SmsUrl. Either GET or POST.
func SMSMethod(m Method) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.SMSMethod = m.String()
	}
}

// SMSFallbackURL Twilio description:
// A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by SmsUrl.
func SMSFallbackURL(url string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.SMSFallbackURL = url
	}
}

// SMSFallbackMethod Twilio description:
// The HTTP method that should be used to request the SmsUrl. Either GET or POST.
func SMSFallbackMethod(m Method) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.SMSFallbackMethod = m.String()
	}
}

// SMSApplicationSID Twilio description:
// The 34 character sid of the application Twilio should use to handle SMSs sent to this number.
// If a SmsApplicationSid is present, Twilio will ignore all of the SMS urls above and use those set on the application instead.
func SMSApplicationSID(sid string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.SMSApplicationSID = sid
	}
}

// AccountSID Twilio description:
// The unique 34 character id of the account to which you wish to transfer this phone number. See Exchanging Numbers Between Subaccounts.
func AccountSID(sid string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.AccountSID = sid
	}
}

// AddressSID Twilio description:
// The 34 character sid of the address Twilio should associate with the number. If the number has address restrictions,
// only another address that satisfies the requirement can replace the existing one.
func AddressSID(sid string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.AddressSID = sid
	}
}
