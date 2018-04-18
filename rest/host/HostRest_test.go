package host_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/config"
)

// TestHostRestGetInfo .
func TestHostRestGetInfo(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/host/info")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

// TestHostRestGetTemperature .
func TestHostRestGetTemperature(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/host/temperature")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}
