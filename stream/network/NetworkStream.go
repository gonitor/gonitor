package network

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/network"
	"github.com/gonitor/gonitor/util"
)

// NetworkStreamGetInterfaces .
func NetworkStreamGetInterfaces(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := network.ServiceGetInterfaces()
		return util.StreamHandleResponse(context, body, err, "NetworkStreamGetInterfaces")
	})
}

// NetworkStreamGetConnections .
func NetworkStreamGetConnections(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := network.ServiceGetConnections()
		return util.StreamHandleResponse(context, body, err, "NetworkStreamGetConnections")
	})
}
