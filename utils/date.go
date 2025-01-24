package utils

import (
	"time"
	"fmt"
)

func ParseDate(dateStr string, format string) (time.Time, error) {
	parsedDate, err := time.Parse(format, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date: %v", err)
	}
	return parsedDate, nil
}

func FormatDate(date time.Time, format string) string {
	return date.Format(format)
}

func GetCurrentDate(format string) string {
	return time.Now().Format(format)
}

func ConvertDateFormat(dateStr string, fromFormat string, toFormat string) (string, error) {
	parsedDate, err := ParseDate(dateStr, fromFormat)
	if err != nil {
		return "", err
	}
	return FormatDate(parsedDate, toFormat), nil
}

func IsValidDate(dateStr string, format string) bool {
	_, err := ParseDate(dateStr, format)
	return err == nil
}
