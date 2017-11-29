package vtwilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type errorMessage struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

func handleRequest(req *http.Request) ([]byte, error) {
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
		var err errorMessage
		json.Unmarshal(bodyBytes, &err)
		return nil, fmt.Errorf("Error: %v", err.Message)
	}

	return bodyBytes, nil
}

func setUpRequest(req *http.Request, accountSID, authToken string) {
	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}

func handleMessage(req *http.Request) (*Message, error) {
	bodyBytes, err := handleRequest(req)
	if err != nil {
		return nil, err
	}

	var data Message
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func handleListMessages(req *http.Request) (*List, error) {
	bodyBytes, err := handleRequest(req)
	if err != nil {
		return nil, err
	}

	var data List
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func handleAvailability(req *http.Request) (*AvailablePhoneNumbers, error) {
	bodyBytes, err := handleRequest(req)
	if err != nil {
		return nil, err
	}

	var data AvailablePhoneNumbers
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func handleIncomingPhoneNumbers(req *http.Request) (*IncomingPhoneNumber, error) {
	bodyBytes, err := handleRequest(req)
	if err != nil {
		return nil, err
	}

	var data IncomingPhoneNumber
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func genericHandler(req *http.Request) error {
	if _, err := handleRequest(req); err != nil {
		return err
	}
	return nil
}
