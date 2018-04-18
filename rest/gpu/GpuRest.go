package gpu

import (
	"github.com/gin-gonic/gin"

	"github.com/gonitor/gonitor/service/gpu"
	"github.com/gonitor/gonitor/util"
)

// GpuRestGetInfo gets the GPU information via GET request.
func GpuRestGetInfo(context *gin.Context) {
	body, err := gpu.ServiceGetInfo()
	util.RestHandleResponse(context, body, err)
}
