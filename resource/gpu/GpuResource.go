package gpu

import (
	"github.com/gin-gonic/gin"

	"github.com/gonitor/gonitor/util"
)

// GpuGetInfo .
func GpuGetInfo(context *gin.Context) {
	body, err := GetInfo()
	if err == nil {
		context.Header(util.GetJSONHeader())
		context.String(200, body)
	} else {
		context.String(400, err.Error())
	}
}
