package host

import (
	"github.com/gin-gonic/gin"
)

// HostGetInfo .
func HostGetInfo(context *gin.Context) {
	body, err := GetInfo()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// HostGetTemperature .
func HostGetTemperature(context *gin.Context) {
	body, err := GetTemperature()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
