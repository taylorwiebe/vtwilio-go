# VTwilio - Go
### Version 0.0.1
Call the twilio rest api in go.

## Example
```
func sendTwilioMessage() {
	t := vtwilio.NewVTwilio(accountSID, authToken, twilioNumber)
	resp, err := t.SendMessage("Hello world", "12345678910")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
```

## Change Log
### v0.0.1
- Currently still a work in progress
