package runtime

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/runtime"
)

// RuntimeRestGetGoOS .
func RuntimeRestGetGoOS(context *gin.Context) {

	body := runtime.ServiceGetGoOS()
	context.JSON(200, body)
}
