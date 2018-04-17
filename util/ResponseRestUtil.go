package util

import "github.com/gin-gonic/gin"

// RestHandleResponse .
func RestHandleResponse(context *gin.Context, body interface{}, err error) {
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
