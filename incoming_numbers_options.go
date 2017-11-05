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

// IncomingPhoneNumberOption options for an incoming phone number purchase
type IncomingPhoneNumberOption func(*incomingNumberConfiguration)
