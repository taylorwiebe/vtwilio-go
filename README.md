# VTwilio - Go
#####  Note: breaking changes may happen in releases before v1.0.0
### Version 0.1.0
Call the twilio rest api in go.

## Examples
### Send a message
```
func sendTwilioMessage() {
	t := vtwilio.NewVTwilio(sid, token, vtwilio.TwilioNumber(twilioNumber))
	message, err := t.SendMessage("Hello world", "12345678910")
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
}
```

### Get a single message
```
func GetATwilioMessage() {
	t := vtwilio.NewVTwilio(sid, token, vtwilio.TwilioNumber(twilioNumber))
	message, err := t.GetMessage(messageSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
}
```

### Get a List of Messages
#### List Options
- `PageSize(int)` - current page, defaults to 10
- `Page(int)` - size of the page, defaults to 0
- `OnDate(time.Time)` - Get messages on a date
- `OnAndBeforeDate(time.Time)` - Get message on and before a given date
- `OnAndAfterDate(time.Time)` - Get messages on and after given date
```
func ListTwilioMessages() {
	t := vtwilio.NewVTwilio(accountSID, authToken, vtwilio.TwilioNumber(twilioNumber))
	messages, err := t.ListMessages(vtwilio.PageSize(1), vtwilio.Page(0))
	if err != nil {
		panic(err)
	}
	fmt.Println(messages)
}
```

### Get Available Numbers
#### Available Number Options
- `NearNumber`
- `NearLatLong`
- `Distance`
- `InPostalCode`
- `InLocality`
- `InRegion`
- `InRateCenter`
- `InLATA`

```
func GetAvailableNumbers() {
	t := vtwilio.NewVTwilio(sid, token)
	numbers, err := t.AvailablePhoneNumbers("US", vtwilio.InRegion("CA"))
	if err != nil {
		panic(err)
	}
	fmt.Println(numbers)
}
```
### Incoming Phone Number
After using the `AvailablePhoneNumbers` method, a number can be chosen and purchased from Twilio using the `IncomingPhoneNumber` method.

Note: This will not work with a Twilio trial account.
#### Incoming Phone Number Options
- `AreaCode`
- `APIVersion`
- `FriendlyName`
- `VoiceURL`
- `VoiceMethod`
- `VoiceFallBackURL`
- `VoiceFallBackMethod`
- `StatusCallback`
- `StatusCallbackMethod`
- `VoiceCallerIDLookup`
- `VoiceApplicationSID`
- `TrunkSID`
- `SMSURL`
- `SMSMethod`
- `SMSFallbackURL`
- `SMSFallbackMethod`
- `SMSApplicationSID`
- `AccountSID`
- `AddressSID`
#### Purchase a number example
```
func PurchaseNumber(number string) (*vtwilio.IncomingPhoneNumber, error) {
	t := vtwilio.NewVTwilio(sid, token)
	num, err := t.IncomingPhoneNumber(number)
	if err != nil {
		return nil, err
	}
	return num, nil
}
```
#### Update an existing number
Updating a number uses the same options as purchasing one.
```
func UpdateNumber() error {
	t := vtwilio.NewVTwilio(sid, token)
	_, err := t.UpdateIncomingPhoneNumber("+12345678910", "MESSAGE_SID",
		vtwilio.FriendlyName("New Friendly Name"))
	if err != nil {
		return err
	}
	return nil
}
```
#### Release a number
This will delete a Twilio number from your account, and allow someone else to potentially buy this number.
```
func ReleaseNumber() error {
	t := vtwilio.NewVTwilio(sid, token)
	err := t.ReleaseNumber("NUMBER_SID")
	if err != nil {
		return err
	}
	return nil
}
```

### TwiML
[TwiML Docs](./twiml/README.md)

## Change Log
### v0.1.1
- Fix typo
### v0.1.0
- Add from number option for send method to override the client's default
- Add callback url option to send method
### v0.0.3
- TwiML support
- Incoming phone numbers
- Update phone number settings
- Release a phone number
### v0.0.2
- Refactoring code
- Add lookup for available phone numbers
#### Breaking Changes
- No longer required to pass a number to VTwilio in the case that you want to look up a number.
This is now an option

### v0.0.1
- Currently still a work in progress
