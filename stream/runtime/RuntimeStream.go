package runtime

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/runtime"
	"github.com/gonitor/gonitor/util"
)

// RuntimeStreamGetGoOS streams the operating system (OS) of Go runtime via GET request.
func RuntimeStreamGetGoOS(context *gin.Context) {
	rateLimit, err := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if err != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body := runtime.ServiceGetGoOS()
		return util.StreamHandleResponse(context, body, nil, "RuntimeStreamGetGoOS")
	})
}
