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

	route.SetRoutes(router)

	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}

// TestLoadGetAverage .
func TestLoadGetAverage(test *testing.T) {
	testRouter := SetupRouter()

	url := "/api/v1/load/average"
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}

// TestLoadGetMisc .
func TestLoadGetMisc(test *testing.T) {
	testRouter := SetupRouter()

	url := "/api/v1/load/misc"
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}
