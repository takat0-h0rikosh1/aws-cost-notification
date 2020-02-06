package main

import (
	"aws-cost-notification/app/aws"
	"aws-cost-notification/app/config"
	"aws-cost-notification/app/gaitameonline"
	"aws-cost-notification/app/service"
	"aws-cost-notification/app/slack"
	"fmt"
	"log"
)

func init() {
	config.InitConfig()
}

func main() {
	fmt.Println("START.")

	period, prefix, channel := parseArgs()

	costNotificationService := service.NewCostNotificationService(
		prefix,
		aws.CostExplorerClientImpl{},
		gaitameonline.GaitameOnlineClientImpl{Url: config.C.ExchangeRateURL},
		slack.SlackClientImpl{
			Token:      config.C.SlackToken,
			PostMsgUrl: config.C.SlackPostMessageURL,
			Channel:    channel,
		},
	)
	if err := costNotificationService.Notify(&period); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done.")
}
