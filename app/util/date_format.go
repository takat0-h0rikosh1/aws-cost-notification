package util

type DateFormat string

const (
	DefaultDateFormat  = DateFormat("2006-01-02")
	ExactlyMonthFormat = DateFormat("2006/01")
)

func (df DateFormat) String() string {
	return string(df)
}
