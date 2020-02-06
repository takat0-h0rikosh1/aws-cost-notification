package util

import "time"

// time.Time を yyyy/mm へ変換する
func FormatToMonthString(t time.Time) string {
	return t.Format(ExactlyMonthFormat.String())
}
