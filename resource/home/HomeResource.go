package home

import (
	"github.com/gin-gonic/gin"
)

//HomeGetIndex .
func HomeGetIndex(c *gin.Context) {
	c.JSON(200, gin.H{"data": "hello"})
}
