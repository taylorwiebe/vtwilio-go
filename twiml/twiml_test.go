package twiml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/twiebe-va/vtwilio-go/twiml"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name          string
		in            *twiml.TwiML
		expected      string
		expectedError bool
	}{
		{
			name: "message",
			in:   twiml.NewTwiML().Message("My Message"),
			expected: "<Response>\n	<Message>My Message</Message>\n</Response>",
		},
		{
			name: "redirect",
			in:   twiml.NewTwiML().Redirect("/someplace.xml"),
			expected: "<Response>\n	<Redirect>/someplace.xml</Redirect>\n</Response>",
		},
		{
			name: "say",
			in:   twiml.NewTwiML().Say("My message"),
			expected: "<Response>\n	<Say>My message</Say>\n</Response>",
		},
		{
			name: "say loop",
			in:   twiml.NewTwiML().Say("My message", twiml.SayLoop(2)),
			expected: "<Response>\n	<Say loop=\"2\">My message</Say>\n</Response>",
		},
		{
			name: "say language",
			in:   twiml.NewTwiML().Say("My message", twiml.SayLanguage(twiml.EN)),
			expected: "<Response>\n	<Say language=\"en\">My message</Say>\n</Response>",
		},
		{
			name: "say voice",
			in:   twiml.NewTwiML().Say("My message", twiml.SayVoice(twiml.Alice)),
			expected: "<Response>\n	<Say voice=\"Alice\">My message</Say>\n</Response>",
		},
		{
			name: "say all options",
			in:   twiml.NewTwiML().Say("My message", twiml.SayLoop(2), twiml.SayLanguage(twiml.EN), twiml.SayVoice(twiml.Alice)),
			expected: "<Response>\n	<Say voice=\"Alice\" language=\"en\" loop=\"2\">My message</Say>\n</Response>",
		},
		{
			name: "sms",
			in:   twiml.NewTwiML().SMS("sms message"),
			expected: "<Response>\n	<Sms>sms message</Sms>\n</Response>",
		},
		{
			name: "sms to",
			in:   twiml.NewTwiML().SMS("sms message", twiml.SMSTo("+12345678910")),
			expected: "<Response>\n	<Sms to=\"+12345678910\">sms message</Sms>\n</Response>",
		},
		{
			name: "sms from",
			in:   twiml.NewTwiML().SMS("sms message", twiml.SMSFrom("+12345678910")),
			expected: "<Response>\n	<Sms from=\"+12345678910\">sms message</Sms>\n</Response>",
		},
		{
			name: "sms action",
			in:   twiml.NewTwiML().SMS("sms message", twiml.SMSAction("/action.html")),
			expected: "<Response>\n	<Sms action=\"/action.html\">sms message</Sms>\n</Response>",
		},
		{
			name: "sms method GET",
			in:   twiml.NewTwiML().SMS("sms message", twiml.SMSMethod(twiml.GET)),
			expected: "<Response>\n	<Sms method=\"GET\">sms message</Sms>\n</Response>",
		},
		{
			name: "sms method POST",
			in:   twiml.NewTwiML().SMS("sms message", twiml.SMSMethod(twiml.POST)),
			expected: "<Response>\n	<Sms method=\"POST\">sms message</Sms>\n</Response>",
		},
		{
			name: "sms method POST",
			in:   twiml.NewTwiML().SMS("sms message", twiml.SMSStatusCallBack("/callback")),
			expected: "<Response>\n	<Sms statusCallback=\"/callback\">sms message</Sms>\n</Response>",
		},
		{
			name: "sms all options",
			in: twiml.NewTwiML().SMS("sms message",
				twiml.SMSStatusCallBack("/callback"),
				twiml.SMSMethod(twiml.POST),
				twiml.SMSAction("/action.html"),
				twiml.SMSFrom("+12345678910"),
				twiml.SMSTo("+12345678910")),
			expected: "<Response>\n\t<Sms to=\"+12345678910\" from=\"+12345678910\" action=\"/action.html\" method=\"POST\" statusCallback=\"/callback\">sms message</Sms>\n</Response>",
		},
		{
			name: "dial",
			in:   twiml.NewTwiML().Dial(),
			expected: "<Response>\n	<Dial></Dial>\n</Response>",
		},
		{
			name: "dial number",
			in:   twiml.NewTwiML().Dial(twiml.DialNumber("+12345678910")),
			expected: "<Response>\n	<Dial>\n\t\t<Number>+12345678910</Number>\n\t</Dial>\n</Response>",
		},
		{
			name: "dial action",
			in:   twiml.NewTwiML().Dial(twiml.DialAction("/action")),
			expected: "<Response>\n	<Dial action=\"/action\"></Dial>\n</Response>",
		},
		{
			name: "dial method POST",
			in:   twiml.NewTwiML().Dial(twiml.DialMethod(twiml.POST)),
			expected: "<Response>\n	<Dial method=\"POST\"></Dial>\n</Response>",
		},
		{
			name: "dial method GET",
			in:   twiml.NewTwiML().Dial(twiml.DialMethod(twiml.GET)),
			expected: "<Response>\n	<Dial method=\"GET\"></Dial>\n</Response>",
		},
		{
			name: "dial caller id",
			in:   twiml.NewTwiML().Dial(twiml.DialCallerID("+123445678910")),
			expected: "<Response>\n	<Dial callerId=\"+123445678910\"></Dial>\n</Response>",
		},
		{
			name: "dial all options",
			in: twiml.NewTwiML().Dial(twiml.DialCallerID("+123445678910"),
				twiml.DialMethod(twiml.POST),
				twiml.DialAction("/action"),
				twiml.DialNumber("+12345678910")),
			expected: "<Response>\n	<Dial action=\"/action\" method=\"POST\" callerId=\"+123445678910\">\n\t\t<Number>+12345678910</Number>\n\t</Dial>\n</Response>",
		},
		{
			name: "reject",
			in:   twiml.NewTwiML().Reject(),
			expected: "<Response>\n	<Reject></Reject>\n</Response>",
		},
		{
			name: "reject reason busy",
			in:   twiml.NewTwiML().Reject(twiml.RejectReason(twiml.Busy)),
			expected: "<Response>\n	<Reject reason=\"busy\"></Reject>\n</Response>",
		},
		{
			name: "reject reason rejected",
			in:   twiml.NewTwiML().Reject(twiml.RejectReason(twiml.Rejected)),
			expected: "<Response>\n	<Reject reason=\"rejected\"></Reject>\n</Response>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.in.Build()
			if tt.expectedError && err == nil {
				t.Error("expected an error, got none")
			} else if !tt.expectedError && err != nil {
				t.Errorf("did not expect an error, got %v", err)
			}
			assert.Equal(t, tt.expected, string(actual))
		})
	}
}
