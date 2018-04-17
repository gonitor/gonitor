package host_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/service/host"
)

// TestServiceGetInfo .
func TestServiceGetInfo(test *testing.T) {
	result, err := host.ServiceGetInfo()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, result != nil, true)
}

// TestServiceGetTemperature .
func TestServiceGetTemperature(test *testing.T) {
	_, err := host.ServiceGetTemperature()
	assert.Equal(test, err == nil, true)
}
