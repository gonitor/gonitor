package runtime

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/runtime"
	"github.com/gonitor/gonitor/util"
)

// RuntimeRestGetGoOS .
func RuntimeRestGetGoOS(context *gin.Context) {
	body := runtime.ServiceGetGoOS()
	util.RestHandleResponse(context, body, nil)
}
