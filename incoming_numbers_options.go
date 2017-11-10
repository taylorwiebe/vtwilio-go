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
	VoiceCallerIDLookup  string `vtwilio:"VoiceCallerIdLookup"`
	VoiceApplicationSID  string `vtwilio:"VoiceApplicationSid"`
	TrunkSid             string `vtwilio:"TrunkSid"`
	SMSURL               string `vtwilio:"SmsUrl"`
	SMSMethod            string `vtwilio:"SmsMethod"`
	SMSFallbackURL       string `vtwilio:"SmsFallbackUrl"`
	SMSFallbackMethod    string `vtwilio:"SmsFallbackMethod"`
	SMSApplicationSid    string `vtwilio:"SmsApplicationSid"`
	AddressSid           string `vtwilio:"AddressSid"`
	APIVersion           string `vtwilio:"ApiVersion"`
}

// AreaCode Twilio description:
// The desired area code for your new incoming phone number.
// Any three digit, US or Canada area code is valid. Twilio will provision a random phone number within
// this area code for you. You must include either this or a PhoneNumber parameter to have your POST succeed. (US and Canada only)
func AreaCode(a string) IncomingPhoneNumberOption {
	return func(i *incomingNumberConfiguration) {
		i.AreaCode = a
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

// IncomingPhoneNumberOption options for an incoming phone number purchase
type IncomingPhoneNumberOption func(*incomingNumberConfiguration)
