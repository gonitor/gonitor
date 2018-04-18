package load

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/load"
	"github.com/gonitor/gonitor/util"
)

// LoadRestGetAverage gets the load average via GET request.
func LoadRestGetAverage(context *gin.Context) {
	body, err := load.ServiceGetAverage()
	util.RestHandleResponse(context, body, err)
}

// LoadRestGetMisc gets the load misc via GET request.
func LoadRestGetMisc(context *gin.Context) {
	body, err := load.GetMisc()
	util.RestHandleResponse(context, body, err)
}
