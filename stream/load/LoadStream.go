package load

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/load"
	"github.com/gonitor/gonitor/util"
)

// LoadStreamGetAverage .
func LoadStreamGetAverage(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := load.ServiceGetAverage()
		return util.StreamHandleResponse(context, body, err, "LoadStreamGetAverage")
	})
}

// LoadStreamGetMisc .
func LoadStreamGetMisc(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := load.GetMisc()
		return util.StreamHandleResponse(context, body, err, "LoadStreamGetMisc")
	})
}
