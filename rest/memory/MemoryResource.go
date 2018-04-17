package memory

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/memory"
)

// MemoryRestGetVirtual .
func MemoryRestGetVirtual(context *gin.Context) {
	body, err := memory.ServiceGetVirtualMemory()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// MemoryRestGetSwap .
func MemoryRestGetSwap(context *gin.Context) {
	body, err := memory.ServiceGetSwapMemory()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
