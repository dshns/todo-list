package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, date string, repeat string) (string, error) {
	taskDate, err := time.Parse(DateFormat, date)
	if err != nil {
		return "", fmt.Errorf("invalid date format: %v", err)
	}

	if repeat == "" {
		return "", nil
	}

	parts := strings.Split(repeat, " ")
	if len(parts) == 0 {
		return "", fmt.Errorf("empty repeat rule")
	}
	rule := parts[0]
	if rule == "d" {
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid 'd' rule format")
		}

		days, err := strconv.Atoi(parts[1])
		if err != nil || days < 1 || days > 400 {
			return "", fmt.Errorf("invalid number of days")
		}

		return nextDateByDays(taskDate, now, days), nil
	} else if rule == "y" {
		if len(parts) != 1 {
			return "", fmt.Errorf("invalid 'y' rule format")
		}
		return nextDateByYear(taskDate, now), nil
	} else {
		return "", fmt.Errorf("unsupported format")
	}
}

func nextDateByDays(taskDate, now time.Time, days int) string {
	for {
		taskDate = taskDate.AddDate(0, 0, days)
		if taskDate.After(now) {
			return taskDate.Format(DateFormat)
		}
	}
}

func nextDateByYear(taskDate, now time.Time) string {
	for {
		taskDate = taskDate.AddDate(1, 0, 0)
		if taskDate.After(now) {
			return taskDate.Format(DateFormat)
		}
	}
}
