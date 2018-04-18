package runtime_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/config"
)

//TestHomeGetIndex .
func TestRuntimeRestGetGoOS(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/runtime/goos")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

// TestCleanUp .
func TestCleanUp(t *testing.T) {

}
