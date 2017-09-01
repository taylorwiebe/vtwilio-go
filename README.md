# VTwilio - Go
### Version 0.0.1
Call the twilio rest api in go.

## Examples
### Send a message
```
func sendTwilioMessage() {
	t := vtwilio.NewVTwilio(accountSID, authToken, twilioNumber)
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
	t := vtwilio.NewVTwilio(accountSID, authToken, twilioNumber)
	message, err := t.GetMessage(messageSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
}
```

### Get a list of messages
#### List options
- `PageSize(int)` - current page, defaults to 10
- `Page(int)` - size of the page, defaults to 0
- `OnDate(time.Time)` - Get messages on a date
- `OnAndBeforeDate(time.Time)` - Get message on and before a given date
- `OnAndAfterDate(time.Time)` - Get messages on and after given date
```
func ListTwilioMessages() {
	t := vtwilio.NewVTwilio(accountSID, authToken, twilioNumber)
	messages, err := t.ListMessages(vtwilio.PageSize(1), vtwilio.Page(0))
	if err != nil {
		panic(err)
	}
	fmt.Println(messages)
}
```

## Change Log
### v0.0.1
- Currently still a work in progress
