package network

import (
	"github.com/gin-gonic/gin"
)

// NetworkGetInterfaces .
func NetworkGetInterfaces(context *gin.Context) {
	body, err := GetInterfaces()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// NetworkGetConnections .
func NetworkGetConnections(context *gin.Context) {
	body, err := GetConnections()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
