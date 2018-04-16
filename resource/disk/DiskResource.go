package disk

import (
	"github.com/gin-gonic/gin"
)

// DiskGetUsage .
func DiskGetUsage(context *gin.Context) {
	body, err := GetUsage()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
