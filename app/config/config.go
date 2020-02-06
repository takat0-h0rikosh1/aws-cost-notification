package config

import (
	"os"
)

var C Config

type Config struct {
	ExchangeRateURL     string
	CostExplorerURL     string
	SlackPostMessageURL string
	SlackToken          string
}

func InitConfig() {
	C = Config{
		ExchangeRateURL:     "https://www.gaitameonline.com/rateaj/getrate",
		CostExplorerURL:     "https://console.aws.amazon.com/cost-reports/home",
		SlackPostMessageURL: "https://slack.com/api/chat.postMessage",
		SlackToken:          os.Getenv("SLACK_TOKEN"),
	}
}
