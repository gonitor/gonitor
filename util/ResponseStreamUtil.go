package util

import (
	"github.com/gin-gonic/gin"
)

// StreamHandleResponse handles the Stream API response.
func StreamHandleResponse(context *gin.Context, body interface{}, err error, messageName string) bool {
	if err != nil {
		return false
	}
	context.SSEvent(messageName, body)
	return true
}
