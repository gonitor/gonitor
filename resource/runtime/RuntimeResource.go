package runtime

import (
	"github.com/gin-gonic/gin"
)

// RuntimeGetGoOs .
func RuntimeGetGoOs(context *gin.Context) {

	body := GetGOOS()
	context.JSON(200, body)
}
