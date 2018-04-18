package route

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/config"
)

// SetRoutes sets all routes or paths for API.
func SetRoutes(router *gin.Engine) {
	if config.EnableRestAPI {
		RestV1SetRoutes(router)
	}
	if config.EnableStreamAPI {
		StreamV1SetRoutes(router)
	}
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
