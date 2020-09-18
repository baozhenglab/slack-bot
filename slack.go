package slackbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"time"

	goservice "github.com/baozhenglab/go-sdk"
)

type slackService struct {
	webhook string
}

var client = http.Client{
	Timeout: 5 * time.Second,
}

const (
	KeyService = "slack-bot"
)

func (slack *slackService) Name() string {
	return KeyService
}

func (slack *slackService) GetPrefix() string {
	return KeyService
}

func (slack *slackService) InitFlags() {
	prefix := fmt.Sprintf("%s-", slack.Name())
	flag.StringVar(&slack.webhook, prefix+"webhook-url", "", "Webhook url of slack bot")
}

func (slack *slackService) Get() interface{} {
	return slack
}

func NewSlackBot() goservice.PrefixConfigure {
	return new(slackService)
}

func (slack *slackService) SendMessage(form map[string]string) error {
	endPointRequest := slack.webhook
	jsonValue, err := json.Marshal(form)
	if err != nil {
		return nil
	}
	req, err := http.NewRequest("POST", endPointRequest, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	var parse map[string]interface{}
	json.NewDecoder(response.Body).Decode(&parse)
	if parse["ok"] != true {
		return errors.New(parse["error"].(string))
	}
	return nil
}
