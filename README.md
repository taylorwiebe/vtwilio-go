# VTwilio - Go
### Version 0.0.1
Call the twilio rest api in go.

## Example
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
func sendTwilioMessage() {
	t := vtwilio.NewVTwilio(accountSID, authToken, twilioNumber)
	message, err := t.GetMessage(messageSID)
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
}
```

### Get a list of messages
```
func sendTwilioMessage() {
	t := vtwilio.NewVTwilio(accountSID, authToken, twilioNumber)
	messages, err := t.ListMessages(vtwilio.PageSize(1), vtwilio.Page(0))
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
}
```

## Change Log
### v0.0.1
- Currently still a work in progress
