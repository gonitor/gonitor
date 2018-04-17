package network_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/service/network"
)

// TestServiceGetInterfaces .
func TestServiceGetInterfaces(test *testing.T) {
	result, err := network.ServiceGetInterfaces()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, result != nil, true)
}
