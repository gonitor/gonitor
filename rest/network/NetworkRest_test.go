package network_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/route"
	"github.com/gonitor/gonitor/service/network"
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

// TestNetworkRestGetInterfaces .
func TestNetworkRestGetInterfaces(test *testing.T) {
	testRouter := SetupRouter()

	url := route.GetRestEndPoint("/network/interfaces")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	_, err := network.ServiceGetInterfaces()
	if err == nil {
		assert.Equal(test, resp.Code, 200)
	} else {
		assert.Equal(test, resp.Code, 400)
	}

}

// NetworkRestGetConnections .
func TestNetworkRestGetConnections(test *testing.T) {
	testRouter := SetupRouter()

	url := route.GetRestEndPoint("/network/connections")
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	_, err := network.ServiceGetConnections()
	if err == nil {
		assert.Equal(test, resp.Code, 200)
	} else {
		assert.Equal(test, resp.Code, 400)
	}
}
