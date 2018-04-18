package config

import (
	"github.com/gin-gonic/gin"
)

// SetupTestRouter sets up the test router.
func SetupTestRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	SetRoutes(router)

	return router
}
