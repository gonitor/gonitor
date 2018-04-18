package disk

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/disk"
	"github.com/gonitor/gonitor/util"
)

// DiskStreamGetUsage streams the disk usage via GET request.
func DiskStreamGetUsage(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := disk.ServiceGetUsage()
		return util.StreamHandleResponse(context, body, err, "DiskStreamGetUsage")
	})
}
