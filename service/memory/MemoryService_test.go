package memory_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/service/memory"
)

// TestServiceGetVirtualMemory .
func TestServiceGetVirtualMemory(test *testing.T) {
	result, err := memory.ServiceGetVirtualMemory()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, result != nil, true)
}

// TestServiceGetSwapMemory .
func TestServiceGetSwapMemory(test *testing.T) {
	result, err := memory.ServiceGetSwapMemory()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, result != nil, true)
}
