package common

import "time"

func StrToDate(value string) *time.Time {
	if value == "" {
		return nil
	}

	date, err := time.Parse(DATE_FORMAT, value)
	if err != nil {
		panic(err)
	}

	return &date
}

func DateToStr(value *time.Time) string {
	if value == nil {
		return ""
	}

	return value.Format(DATE_FORMAT)
}
