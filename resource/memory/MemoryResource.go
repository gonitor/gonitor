package memory

import (
	"github.com/gin-gonic/gin"
)

// MemoryGetVirtual .
func MemoryGetVirtual(context *gin.Context) {
	body, err := GetVirtualMemory()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// MemoryGetSwap .
func MemoryGetSwap(context *gin.Context) {
	body, err := GetSwapMemory()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
