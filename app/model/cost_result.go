package model

import (
	"aws-cost-notification/app/util"
	"fmt"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"sort"
	"strconv"
	"time"
)

type CostResult struct {
	CurrentMonthCost      float64   // とある月のコスト
	LastMonthCost         float64   // とある月の前の月のコスト
	DiffCost              float64   // とある月とその前の月とコストの差額
	CurrentMonthStartDate time.Time // とある月の最初の日（メッセージ出力用なので最初の日であることに特に意味はない
}

// とある月のコストとその前の月のコストを受け取って CostResult を作る
func NewCostResult(current *costexplorer.ResultByTime, last *costexplorer.ResultByTime) (*CostResult, error) {

	currentMonthCost, err := strconv.ParseFloat(*current.Total[AmortizedCost.String()].Amount, 32)
	if err != nil {
		return nil, fmt.Errorf("failed convert to float: %s", err)
	}

	lastMonthCost, err := strconv.ParseFloat(*last.Total[AmortizedCost.String()].Amount, 32)
	if err != nil {
		return nil, fmt.Errorf("failed convert to float: %s", err)
	}

	t, err := time.Parse(util.DefaultDateFormat.String(), *current.TimePeriod.Start)
	if err != nil {
		return nil, fmt.Errorf("parse failed string date: %s", err)
	}

	return &CostResult{
		CurrentMonthCost:      currentMonthCost,
		LastMonthCost:         lastMonthCost,
		DiffCost:              currentMonthCost - lastMonthCost,
		CurrentMonthStartDate: t,
	}, nil
}

// コストエクスプローラーの問い合わせ結果を受け取って CostResult の配列を組み立てる
func CreateCostResult(costAndUsage costexplorer.GetCostAndUsageOutput) ([]*CostResult, error) {

	// 月毎のコストの配列を取得
	resultsByMonth := costAndUsage.ResultsByTime

	// 当月と先月で比較するので事前にソートしておく
	sort.Slice(resultsByMonth, func(i, j int) bool { return *resultsByMonth[i].TimePeriod.Start > *resultsByMonth[j].TimePeriod.Start })

	// 集計
	var results []*CostResult
	for i, v := range resultsByMonth[:len(resultsByMonth)-1] {
		costResult, err := NewCostResult(v, resultsByMonth[i+1])
		if err != nil {
			return nil, err
		}
		results = append(results, costResult)
	}
	return results, nil
}
