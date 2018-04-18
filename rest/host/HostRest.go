package host

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/host"
	"github.com/gonitor/gonitor/util"
)

// HostRestGetInfo gets the host information via GET request.
func HostRestGetInfo(context *gin.Context) {
	body, err := host.ServiceGetInfo()
	util.RestHandleResponse(context, body, err)
}

// HostRestGetTemperature gets the host temperature via GET request.
func HostRestGetTemperature(context *gin.Context) {
	body, err := host.ServiceGetTemperature()
	util.RestHandleResponse(context, body, err)
}
