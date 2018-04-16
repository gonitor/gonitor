package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/route"
)

//CORSMiddleware ...
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
	router := gin.Default()

	router.Use(CORSMiddleware())

	route.SetRoutes(router)

	router.LoadHTMLGlob("./view/*")

	router.Static("/view", "./view")

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.NoRoute(func(context *gin.Context) {
		context.HTML(404, "404.html", gin.H{})
	})

	router.Run(":9000")
}
