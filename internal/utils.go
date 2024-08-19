package utils

import (
	"fmt"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

func ParseDate(date string) (*time.Time, error) {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	t, err := time.ParseInLocation(TimeFormat, date, location)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date: %v\n", err)
	}

	return &t, nil
}

func HandleDate(date string) (*time.Time, error) {
	if date == "" {
		return nil, nil
	}

	t, err := ParseDate(date)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func HandleDateString(date *time.Time) string {
	if date == nil {
		return ""
	}

	return date.Format(TimeFormat)
}
