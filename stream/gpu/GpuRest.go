package gpu

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gonitor/gonitor/service/gpu"
	"github.com/gonitor/gonitor/util"
)

// GpuStreamGetInfo streams the GPU information via GET request.
func GpuStreamGetInfo(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := gpu.ServiceGetInfo()
		return util.StreamHandleResponse(context, body, err, "GpuStreamGetInfo")
	})
}
