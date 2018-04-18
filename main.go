package main

import (
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/config"
)

// CORSMiddleware configures the CORS middleware.
func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if context.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			context.AbortWithStatus(200)
		} else {
			context.Next()
		}
	}
}

func main() {
	config.LoadEnvVariables()

	router := gin.Default()

	router.Use(CORSMiddleware())

	config.SetRoutes(router)

	router.Use(static.Serve("/", static.LocalFile("./view", true)))

	router.Run(":9000")
}
