package cpu_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/service/cpu"
)

// TestServiceGetSummaryPercent .
func TestServiceGetSummaryPercent(test *testing.T) {
	result, err := cpu.ServiceGetSummaryPercent()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, len(result) > 0, true)
}

// TestServiceGetCount .
func TestServiceGetCount(test *testing.T) {
	result, err := cpu.ServiceGetCount()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, result > 0, true)
}

// TestServiceGetSummaryTime .
func TestServiceGetSummaryTime(test *testing.T) {
	result, err := cpu.ServiceGetSummaryTime()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, len(result) > 0, true)
}

// TestServiceGetInfo .
func TestServiceGetInfo(test *testing.T) {
	result, err := cpu.ServiceGetInfo()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, len(result) > 0, true)
}

// TestServiceGetPercent .
func TestServiceGetPercent(test *testing.T) {
	result, err := cpu.ServiceGetPercent()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, len(result) > 0, true)
}

// TestServiceGetTime .
func TestServiceGetTime(test *testing.T) {
	result, err := cpu.ServiceGetTime()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, len(result) > 0, true)
}
