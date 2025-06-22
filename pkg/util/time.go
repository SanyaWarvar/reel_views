package util

import (
	"errors"
	"fmt"
	"time"
)

func GetCurrentMskTime() time.Time {
	return time.Now().UTC().Add(3 * time.Hour)
}

func GetCurrentUTCTime() time.Time {
	return time.Now().UTC()
}

func ConvertStringToTime(timeStr string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cant convert string {%s} to time", timeStr))
	}
	return &t, nil
}
