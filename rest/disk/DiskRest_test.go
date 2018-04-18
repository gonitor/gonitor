package disk_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/config"
)

// TestDiskRestGetUsage .
func TestDiskRestGetUsage(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/disk/usage")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}
