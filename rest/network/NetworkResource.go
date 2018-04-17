package network

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/network"
)

// NetworkRestGetInterfaces .
func NetworkRestGetInterfaces(context *gin.Context) {
	body, err := network.ServiceGetInterfaces()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// NetworkRestGetConnections .
func NetworkRestGetConnections(context *gin.Context) {
	body, err := network.ServiceGetConnections()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
