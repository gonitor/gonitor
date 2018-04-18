package memory

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/memory"
	"github.com/gonitor/gonitor/util"
)

// MemoryRestGetVirtual gets the virtual memory via GET request.
func MemoryRestGetVirtual(context *gin.Context) {
	body, err := memory.ServiceGetVirtualMemory()
	util.RestHandleResponse(context, body, err)
}

// MemoryRestGetSwap gets the swap memory via GET request.
func MemoryRestGetSwap(context *gin.Context) {
	body, err := memory.ServiceGetSwapMemory()
	util.RestHandleResponse(context, body, err)
}
