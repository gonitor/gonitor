package load_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/service/load"
)

// TestServiceGetAverage .
func TestServiceGetAverage(test *testing.T) {
	result, err := load.ServiceGetAverage()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, result != nil, true)
}

// TestGetMisc .
func TestGetMisc(test *testing.T) {
	result, err := load.GetMisc()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, result != nil, true)
}
