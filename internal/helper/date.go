package helper

import (
	"time"
)

var format = "02/01/2006"

func StrToDate(value string) *time.Time {
	if value == "" {
		return nil
	}

	date, err := time.Parse(format, value)
	if err != nil {
		panic(err)
	}

	return &date
}

func DateToStr(value *time.Time) string {
	if value == nil {
		return ""
	}

	return value.Format(format)
}
