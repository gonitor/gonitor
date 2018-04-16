package gpu_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/resource/gpu"
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

//TestGpuGetInfo .
func TestGpuGetInfo(test *testing.T) {
	testRouter := SetupRouter()

	url := "/api/v1/gpu/info"
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	_, err := gpu.GetInfo()
	if err == nil {
		assert.Equal(test, resp.Code, 200)
	} else {
		assert.Equal(test, resp.Code, 400)
	}
}
