package config

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/route"
)

// SetRoutes sets all routes or paths for API.
func SetRoutes(router *gin.Engine) {
	if EnableRestAPI {
		route.RestV1SetRoutes(router)
	}
	if EnableStreamAPI {
		route.StreamV1SetRoutes(router)
	}
}

// GetRestEndPoint concatenates REST API endpoint with each given URL.
func GetRestEndPoint(url string) string {
	root := route.RestV1GroupEndPoint

	var buffer bytes.Buffer
	buffer.WriteString(root)
	buffer.WriteString(url)
	return buffer.String()
}

// GetStreamEndPoint concatenates Stream API endpoint with each given URL.
func GetStreamEndPoint(url string) string {
	root := route.StreamV1GroupEndPoint

	var buffer bytes.Buffer
	buffer.WriteString(root)
	buffer.WriteString(url)
	return buffer.String()
}
