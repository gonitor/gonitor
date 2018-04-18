package cpu_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/config"
)

//TestCpuRestGetSumPercent .
func TestCpuRestGetSumPercent(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/cpu/sum/percent")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

//TestCpuRestGetCount .
func TestCpuRestGetCount(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/cpu/count")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

//TestCpuRestGetSumTime .
func TestCpuRestGetSumTime(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/cpu/sum/time")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

//TestCpuRestGetInfo .
func TestCpuRestGetInfo(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/cpu/info")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

//TestCpuRestGetPercent .
func TestCpuRestGetPercent(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/cpu/percent")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

//TestCpuRestGetTime .
func TestCpuRestGetTime(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/cpu/time")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}
