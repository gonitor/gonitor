package network

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/network"
	"github.com/gonitor/gonitor/util"
)

// NetworkRestGetInterfaces gets the network interaces via GET request.
func NetworkRestGetInterfaces(context *gin.Context) {
	body, err := network.ServiceGetInterfaces()
	util.RestHandleResponse(context, body, err)
}

// NetworkRestGetConnections gets the network connections via GET request.
func NetworkRestGetConnections(context *gin.Context) {
	body, err := network.ServiceGetConnections()
	util.RestHandleResponse(context, body, err)
}
