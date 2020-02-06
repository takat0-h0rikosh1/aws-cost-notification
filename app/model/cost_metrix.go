package model

type CostMetrics string

// https://docs.aws.amazon.com/ja_jp/aws-cost-management/latest/APIReference/API_GetCostAndUsage.html#API_GetCostAndUsage_RequestSyntax
// NetAmortizedCost, NetUnblendedCost はマスターアカウントによる許可が必要なので対象外とする
const (
	AmortizedCost         = CostMetrics("AmortizedCost")
	BlendedCost           = CostMetrics("BlendedCost")
	UnblendedCost         = CostMetrics("UnblendedCost")
	NormalizedUsageAmount = CostMetrics("NormalizedUsageAmount")
	UsageQuantity         = CostMetrics("UsageQuantity")
)

func (cm CostMetrics) String() string {
	return string(cm)
}

var CostMetrixAllValues = []string{
	AmortizedCost.String(),
	BlendedCost.String(),
	UnblendedCost.String(),
	NormalizedUsageAmount.String(),
	UsageQuantity.String(),
}
