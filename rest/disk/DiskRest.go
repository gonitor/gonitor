package disk

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/disk"
	"github.com/gonitor/gonitor/util"
)

// DiskRestGetUsage .
func DiskRestGetUsage(context *gin.Context) {
	body, err := disk.ServiceGetUsage()
	util.RestHandleResponse(context, body, err)
}
