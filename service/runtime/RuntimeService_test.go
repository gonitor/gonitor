package runtime_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/service/runtime"
)

// TestServiceGetGoOS .
func TestServiceGetGoOS(test *testing.T) {
	result := runtime.ServiceGetGoOS()
	assert.Equal(test, len(result) > 0, true)
}
