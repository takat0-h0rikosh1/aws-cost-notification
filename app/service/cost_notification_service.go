package service

import (
	"aws-cost-notification/app/aws"
	"aws-cost-notification/app/config"
	"aws-cost-notification/app/gaitameonline"
	"aws-cost-notification/app/model"
	"aws-cost-notification/app/slack"
	"aws-cost-notification/app/util"
	"fmt"
	"github.com/dustin/go-humanize"
	"strconv"
	"strings"
)

type CostNotificationService struct {
	MsgPrefix           string
	CostExplorerClient  aws.CostExplorerClient
	GaitameOnlineClient gaitameonline.GaitameOnlineClient
	SlackClient         slack.SlackClient
}

func NewCostNotificationService(msgPrefix string, costExplorerClient aws.CostExplorerClient, gaitameOnlineClient gaitameonline.GaitameOnlineClient, slackClient slack.SlackClient) *CostNotificationService {
	return &CostNotificationService{
		MsgPrefix:           msgPrefix,
		CostExplorerClient:  costExplorerClient,
		GaitameOnlineClient: gaitameOnlineClient,
		SlackClient:         slackClient,
	}
}

func (s *CostNotificationService) Notify(period *model.Period) error {
	costAndUsage, err := s.CostExplorerClient.GetCostAndUsage(period)
	if err != nil {
		return fmt.Errorf("faild to get cost and usage from aws cost explorer: %s", err)
	}

	exchangeRate, err := s.GaitameOnlineClient.GetCurrentExchangeRate()
	if err != nil {
		return fmt.Errorf("faild to get current exchange rate from gaitameonline: %s", err)
	}

	costResults, err := model.CreateCostResult(*costAndUsage)
	if err != nil {
		return fmt.Errorf("failed to create cost results: %s", err)
	}

	msg := CreateMessage(s.MsgPrefix, costResults, exchangeRate)

	err = s.SlackClient.PostMessage(msg)
	if err != nil {
		return fmt.Errorf("failed slack post message: %s", err)
	}

	return nil
}

func CreateMessage(prefix string, costResults []*model.CostResult, exchangeRate *gaitameonline.ExchangeRate) string {

	quote, _ := exchangeRate.USDJPY()
	yenPerUSD, _ := strconv.ParseFloat(quote.Open, 32)

	// build message
	costUnit := "円"
	msgArr := []string{fmt.Sprintf("【%v】", prefix)}
	for _, v := range costResults {
		msgArr = append(msgArr,
			fmt.Sprintf("%vの償却原価は `約%v%v` でした（先月比: `約%v%v`)。",
				util.FormatToMonthString(v.CurrentMonthStartDate),
				humanize.Comma(int64(v.CurrentMonthCost*yenPerUSD)),
				costUnit,
				humanize.Comma(int64(v.DiffCost*yenPerUSD)),
				costUnit,
			),
		)
	}
	msgArr = append(msgArr, config.C.CostExplorerURL)
	return strings.Join(msgArr, "\n")
}
