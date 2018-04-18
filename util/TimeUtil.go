package util

import (
	"strconv"
	"time"
)

// ConvertStringToTimeDuration converts the input string into time duration.
func ConvertStringToTimeDuration(input string) (time.Duration, error) {
	output, err := strconv.Atoi(input)
	return time.Duration(output), err
}
