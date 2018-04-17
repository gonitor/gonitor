package load

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/load"
)

// LoadRestGetAverage .
func LoadRestGetAverage(context *gin.Context) {
	body, err := load.ServiceGetAverage()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// LoadRestGetMisc .
func LoadRestGetMisc(context *gin.Context) {
	body, err := load.GetMisc()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
