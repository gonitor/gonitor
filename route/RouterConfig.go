package route

import "github.com/gin-gonic/gin"

// SetRoutes .
func SetRoutes(router *gin.Engine) {
	ApiV1SetRoutes(router)
}
