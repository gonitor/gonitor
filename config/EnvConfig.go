package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

// EnableRestAPI enables the REST API.
var EnableRestAPI = true

// EnableStreamAPI enables the Stream API.
var EnableStreamAPI = true

// LoadEnvVariables loads environment vaiables.
func LoadEnvVariables() {
	if os.Getenv("GONI_PROMODE") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	if os.Getenv("GONI_REST") == "false" {
		EnableRestAPI = false
	}

	if os.Getenv("GONI_STREAM") == "false" {
		EnableStreamAPI = false
	}
}
