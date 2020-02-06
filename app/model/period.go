package model

import (
	"aws-cost-notification/app/util"
	"time"
)

type Period struct {
	From From
	To   To
}

type From time.Time
type To time.Time

func (p *Period) Default() *Period {
	now := time.Now().In(util.FixedTZLocation)
	today := now.Day()
	p.From = From(now.AddDate(0, -1, -(today - 1)))
	p.To = To(now.AddDate(0, 0, -(today - 1)))
	return p
}

func (f *From) Format(df util.DateFormat) string {
	return time.Time(*f).Format(df.String())
}

func (t *To) Format(df util.DateFormat) string {
	return time.Time(*t).Format(df.String())
}
