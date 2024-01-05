package utils

import (
	"time"
)

func TimeParsed(dateString string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateString)
}
