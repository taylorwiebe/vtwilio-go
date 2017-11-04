# VTwilio - Go
#####  Note: breaking changes may happen in releases before v1.0.0
### Version 0.0.2
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
`NearNumber`
`NearLatLong`
`Distance`
`InPostalCode`
`InLocality`
`InRegion`
`InRateCenter`
`InLATA`

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

## Change Log
### v0.0.2
- Refactoring code
- Add lookup for available phone numbers
#### Breaking Changes
- No longer required to pass a number to VTwilio in the case that you want to look up a number.
This is now an option

### v0.0.1
- Currently still a work in progress
