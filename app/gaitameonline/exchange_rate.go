package gaitameonline

import "fmt"

type ExchangeRate struct {
	Quotes []Quote
}

type Quote struct {
	High             string
	Open             string
	Bid              string
	CurrencyPairCode string
	Ask              string
	Low              string
}

func (q *Quote) IsEmpty() bool {
	return q.Open == ""
}

func (o *ExchangeRate) USDJPY() (Quote, error) {
	var result Quote
	for _, v := range o.Quotes {
		if v.CurrencyPairCode == "USDJPY" {
			result = v
			break
		}
	}
	if result.IsEmpty() {
		return result, fmt.Errorf("not found quote", o)
	} else {
		return result, nil
	}
}
