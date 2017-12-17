package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ErrorMessage response structure
type ErrorMessage struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

// HandleRequest handles a request to twilio
func HandleRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var err ErrorMessage
		json.Unmarshal(bodyBytes, &err)
		return nil, fmt.Errorf("Error: %v", err.Message)
	}

	return bodyBytes, nil
}

// SetUpRequest sets up basic auth on a request
func SetUpRequest(req *http.Request, accountSID, authToken string) {
	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}

// GenericHandler is a generic http request handler
func GenericHandler(req *http.Request) error {
	if _, err := HandleRequest(req); err != nil {
		return err
	}
	return nil
}
