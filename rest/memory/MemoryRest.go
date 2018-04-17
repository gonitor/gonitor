package memory

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/memory"
	"github.com/gonitor/gonitor/util"
)

// MemoryRestGetVirtual .
func MemoryRestGetVirtual(context *gin.Context) {
	body, err := memory.ServiceGetVirtualMemory()
	util.RestHandleResponse(context, body, err)
}

// MemoryRestGetSwap .
func MemoryRestGetSwap(context *gin.Context) {
	body, err := memory.ServiceGetSwapMemory()
	util.RestHandleResponse(context, body, err)
}
