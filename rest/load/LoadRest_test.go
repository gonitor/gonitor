package load_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/config"
)

// TestLoadRestGetAverage .
func TestLoadRestGetAverage(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/load/average")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

// TestLoadRestGetMisc .
func TestLoadRestGetMisc(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/load/misc")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}
