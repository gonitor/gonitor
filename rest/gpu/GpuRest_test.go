package gpu_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/route"
	"github.com/gonitor/gonitor/service/gpu"
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

//TestGpuRestGetInfo .
func TestGpuRestGetInfo(test *testing.T) {
	testRouter := SetupRouter()

	url := route.GetRestEndPoint("/gpu/info")
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
