package main

import (
	"aws-cost-notification/app/model"
	"aws-cost-notification/app/util"
	"flag"
	"log"
	"time"
)

func parseArgs() (model.Period, string, string) {
	period := (&model.Period{}).Default()

	startDate := flag.String("from", period.From.Format(util.ExactlyMonthFormat), "start date to get costs")
	endDate := flag.String("to", period.To.Format(util.ExactlyMonthFormat), "end date to get costs")
	prefix := flag.String("prefix", "no prefix", "view account name when post slack message")
	channel := flag.String("channel", "", "slack post target channel")
	flag.Parse()

	parseFaileMsg := "iligal agument(correct is yyyy/mm)"
	parsedStartDate, err := time.Parse(util.ExactlyMonthFormat.String(), *startDate)
	if err != nil {
		log.Fatalf("%s: %v", parseFaileMsg, *startDate)
	}
	parsedEndDate, err := time.Parse(util.ExactlyMonthFormat.String(), *endDate)
	if err != nil {
		log.Fatalf("%s: %v", parseFaileMsg, *endDate)
	}
	if *channel == "" {
		log.Fatal("`-channel` is required")
	}

	period.From = model.From(parsedStartDate)
	period.To = model.To(parsedEndDate)

	return *period, *prefix, *channel
}
