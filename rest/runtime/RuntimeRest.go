package runtime

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/runtime"
	"github.com/gonitor/gonitor/util"
)

// RuntimeRestGetGoOS gets the operating system (OS) of Go runtime via GET request.
func RuntimeRestGetGoOS(context *gin.Context) {
	body := runtime.ServiceGetGoOS()
	util.RestHandleResponse(context, body, nil)
}
