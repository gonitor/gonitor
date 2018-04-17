package load_test

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

	route.SetRestRoutes(router)

	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}

// TestLoadRestGetAverage .
func TestLoadRestGetAverage(test *testing.T) {
	testRouter := SetupRouter()

	url := route.GetRestEndPoint("/load/average")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

// TestLoadRestGetMisc .
func TestLoadRestGetMisc(test *testing.T) {
	testRouter := SetupRouter()

	url := route.GetRestEndPoint("/load/misc")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}
