package cpu_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/route"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	route.SetRoutes(router)

	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}

//TestCpuGetSumPercent .
func TestCpuGetSumPercent(test *testing.T) {
	testRouter := SetupRouter()

	url := "/api/v1/cpu/sum/percent"
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

//TestCpuGetCount .
func TestCpuGetCount(test *testing.T) {
	testRouter := SetupRouter()

	url := "/api/v1/cpu/count"
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

//TestCpuGetSumTime .
func TestCpuGetSumTime(test *testing.T) {
	testRouter := SetupRouter()

	url := "/api/v1/cpu/sum/time"
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

//TestCpuGetInfo .
func TestCpuGetInfo(test *testing.T) {
	testRouter := SetupRouter()

	url := "/api/v1/cpu/info"
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

//TestCpuGetPercent .
func TestCpuGetPercent(test *testing.T) {
	testRouter := SetupRouter()

	url := "/api/v1/cpu/percent"
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

//TestCpuGetTime .
func TestCpuGetTime(test *testing.T) {
	testRouter := SetupRouter()

	url := "/api/v1/cpu/time"
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}
