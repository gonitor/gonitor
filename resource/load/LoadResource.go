package load

import (
	"github.com/gin-gonic/gin"
)

// LoadGetAverage .
func LoadGetAverage(context *gin.Context) {
	body, err := GetAverage()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// LoadGetMisc .
func LoadGetMisc(context *gin.Context) {
	body, err := GetMisc()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
