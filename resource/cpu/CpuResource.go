package cpu

import (
	"github.com/gin-gonic/gin"
)

// CpuGetSumPercent .
func CpuGetSumPercent(context *gin.Context) {
	body, err := GetSummaryPercent()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// CpuGetCount .
func CpuGetCount(context *gin.Context) {
	body, err := GetCount()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// CpuGetSumTime .
func CpuGetSumTime(context *gin.Context) {
	body, err := GetSummaryTime()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// CpuGetInfo .
func CpuGetInfo(context *gin.Context) {
	body, err := GetInfo()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// CpuGetPercent .
func CpuGetPercent(context *gin.Context) {
	body, err := GetPercent()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// CpuGetTime .
func CpuGetTime(context *gin.Context) {
	body, err := GetTime()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
