package disk

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/disk"
)

// DiskRestGetUsage .
func DiskRestGetUsage(context *gin.Context) {
	body, err := disk.ServiceGetUsage()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
