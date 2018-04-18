package route

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

// SetRoutes sets all routes or paths for API.
func SetRoutes(router *gin.Engine) {
	RestV1SetRoutes(router)
	StreamV1SetRoutes(router)
}

// GetRestEndPoint concatenates REST API endpoint with each given URL.
func GetRestEndPoint(url string) string {
	root := RestV1GroupEndPoint

	var buffer bytes.Buffer
	buffer.WriteString(root)
	buffer.WriteString(url)
	return buffer.String()
}

// GetStreamEndPoint concatenates Stream API endpoint with each given URL.
func GetStreamEndPoint(url string) string {
	root := StreamV1GroupEndPoint

	var buffer bytes.Buffer
	buffer.WriteString(root)
	buffer.WriteString(url)
	return buffer.String()
}
