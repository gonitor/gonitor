package host

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/host"
	"github.com/gonitor/gonitor/util"
)

// HostStreamGetInfo .
func HostStreamGetInfo(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := host.ServiceGetInfo()
		return util.StreamHandleResponse(context, body, err, "HostStreamGetInfo")
	})
}

// HostStreamGetTemperature .
func HostStreamGetTemperature(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := host.ServiceGetTemperature()
		return util.StreamHandleResponse(context, body, err, "HostStreamGetTemperature")
	})
}
