package aws

import (
	"aws-cost-notification/app/model"
	"aws-cost-notification/app/util"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"time"
)

type CostExplorerClient interface {
	GetCostAndUsage(period *model.Period) (*costexplorer.GetCostAndUsageOutput, error)
}

type CostExplorerClientImpl struct{}

func (c CostExplorerClientImpl) GetCostAndUsage(period *model.Period) (*costexplorer.GetCostAndUsageOutput, error) {
	// 先月日を出すため from の月から1を引いて取得条件にする
	from := time.Time(period.From).AddDate(0, -1, 0).Format(util.DefaultDateFormat.String())
	end := time.Time(period.To).Format(util.DefaultDateFormat.String())
	dateInterval := costexplorer.DateInterval{Start: &from, End: &end}

	// コストの取得粒度を月単位に指定
	granularity := "MONTHLY"

	// 全メトリクスを取得対象とする
	var metrics []*string
	for i, _ := range model.CostMetrixAllValues {
		metrics = append(metrics, &model.CostMetrixAllValues[i])
	}

	input := costexplorer.GetCostAndUsageInput{
		Granularity: &granularity,
		Metrics:     metrics,
		TimePeriod:  &dateInterval}

	// 問い合わせ
	sess := session.Must(session.NewSession())
	svc := costexplorer.New(sess)
	result, err := svc.GetCostAndUsage(&input)
	if err != nil {
		return nil, fmt.Errorf("failed getting cost information from aws cost exproler: %s", err)
	}

	return result, nil
}
