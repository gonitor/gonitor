package gpu_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/config"
	"github.com/gonitor/gonitor/service/gpu"
)

//TestGpuRestGetInfo .
func TestGpuRestGetInfo(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/gpu/info")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	_, err := gpu.ServiceGetInfo()
	if err == nil {
		assert.Equal(test, resp.Code, 200)
	} else {
		assert.Equal(test, resp.Code, 400)
	}
}
