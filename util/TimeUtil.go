package util

import (
	"strconv"
	"time"
)

// ConvertStringToTimeDuration .
func ConvertStringToTimeDuration(input string) (time.Duration, error) {
	output, err := strconv.Atoi(input)
	return time.Duration(output), err
}
