package memory_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/config"
)

// TestMemoryRestGetVirtual .
func TestMemoryRestGetVirtual(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/memory/virtual")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)

}

// TestMemoryRestGetSwap .
func TestMemoryRestGetSwap(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/memory/swap")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)

}
