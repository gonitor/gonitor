package host

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/host"
	"github.com/gonitor/gonitor/util"
)

// HostRestGetInfo .
func HostRestGetInfo(context *gin.Context) {
	body, err := host.ServiceGetInfo()
	util.RestHandleResponse(context, body, err)
}

// HostRestGetTemperature .
func HostRestGetTemperature(context *gin.Context) {
	body, err := host.ServiceGetTemperature()
	util.RestHandleResponse(context, body, err)
}
