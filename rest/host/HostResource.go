package host

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/host"
)

// HostRestGetInfo .
func HostRestGetInfo(context *gin.Context) {
	body, err := host.ServiceGetInfo()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// HostRestGetTemperature .
func HostRestGetTemperature(context *gin.Context) {
	body, err := host.ServiceGetTemperature()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
