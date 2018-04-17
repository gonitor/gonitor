package disk_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/service/disk"
)

// TestServiceGetUsage .
func TestServiceGetUsage(test *testing.T) {
	result, err := disk.ServiceGetUsage()
	assert.Equal(test, err == nil, true)
	assert.Equal(test, result != nil, true)
}
