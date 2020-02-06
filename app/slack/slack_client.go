package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type SlackClient interface {
	PostMessage(message string) error
}

type SlackClientImpl struct {
	Token      string
	PostMsgUrl string
	Channel    string
}

type postResult struct {
	OK bool
}

func (c SlackClientImpl) PostMessage(message string) error {

	// build request form
	values := url.Values{}
	values.Add("token", c.Token)
	values.Add("channel", c.Channel)
	values.Add("text", message)

	// prepare http client
	req, _ := http.NewRequest("POST", c.PostMsgUrl, strings.NewReader(values.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := new(http.Client)

	// do request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// defer action
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response postResult
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return fmt.Errorf("failed parse slack api response body: %s", err)
	}

	log.Printf("slack notification result: %v", response.OK)

	return nil
}
