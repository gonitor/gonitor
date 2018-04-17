package memory

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/memory"
	"github.com/gonitor/gonitor/util"
)

// MemoryStreamGetVirtual .
func MemoryStreamGetVirtual(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := memory.ServiceGetVirtualMemory()
		return util.StreamHandleResponse(context, body, err, "MemoryStreamGetVirtual")
	})
}

// MemoryStreamGetSwap .
func MemoryStreamGetSwap(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := memory.ServiceGetSwapMemory()
		return util.StreamHandleResponse(context, body, err, "MemoryStreamGetSwap")
	})
}
