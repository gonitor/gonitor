package route

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

// SetRestRoutes .
func SetRestRoutes(router *gin.Engine) {
	RestV1SetRoutes(router)
}

// GetRestEndPoint .
func GetRestEndPoint(url string) string {
	root := RestV1GroupEndPoint

	var buffer bytes.Buffer
	buffer.WriteString(root)
	buffer.WriteString(url)
	return buffer.String()
}
