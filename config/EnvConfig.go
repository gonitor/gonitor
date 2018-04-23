package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

// EnableProductionMode enables production mode.
var EnableProductionMode = true

// EnableRestAPI enables the REST API.
var EnableRestAPI = true

// EnableStreamAPI enables the Stream API.
var EnableStreamAPI = true

// LoadEnvVariables loads environment vaiables.
func LoadEnvVariables() {
	if os.Getenv("GONI_PROMODE") == "true" || EnableProductionMode {
		gin.SetMode(gin.ReleaseMode)
	}

	if os.Getenv("GONI_REST") == "false" {
		EnableRestAPI = false
	}

	if os.Getenv("GONI_STREAM") == "false" {
		EnableStreamAPI = false
	}
}
