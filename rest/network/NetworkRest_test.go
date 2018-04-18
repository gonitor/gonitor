package network_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor/config"
	"github.com/gonitor/gonitor/service/network"
)

// TestNetworkRestGetInterfaces .
func TestNetworkRestGetInterfaces(test *testing.T) {
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/network/interfaces")
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
	testRouter := config.SetupTestRouter()

	url := config.GetRestEndPoint("/network/connections")
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
