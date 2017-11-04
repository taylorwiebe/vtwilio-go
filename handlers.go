package vtwilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func handleMessageRequest(req *http.Request) (*Message, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Error, with status code: %v", resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
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

func handleListRequest(req *http.Request) (*List, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Error, with status code: %v", resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
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

func handleDateRange(date time.Time, opt dateOption) string {
	t := date.Format("2006-01-02T15:04:05.999999-07:00")
	t = strings.Split(t, "T")[0]
	if opt == before {
		return fmt.Sprintf("DateSent<=%v", t)
	} else if opt == after {
		return fmt.Sprintf("DateSent>=%v", t)
	} else if opt == equal {
		return fmt.Sprintf("DateSent=%v", t)
	}
	return ""
}
