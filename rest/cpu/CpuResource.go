package cpu

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/cpu"
)

// CpuRestGetSumPercent .
func CpuRestGetSumPercent(context *gin.Context) {
	body, err := cpu.ServiceGetSummaryPercent()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// CpuRestGetCount .
func CpuRestGetCount(context *gin.Context) {
	body, err := cpu.ServiceGetCount()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// CpuRestGetSumTime .
func CpuRestGetSumTime(context *gin.Context) {
	body, err := cpu.ServiceGetSummaryTime()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// CpuRestGetInfo .
func CpuRestGetInfo(context *gin.Context) {
	body, err := cpu.ServiceGetInfo()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// CpuRestGetPercent .
func CpuRestGetPercent(context *gin.Context) {
	body, err := cpu.ServiceGetPercent()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}

// CpuRestGetTime .
func CpuRestGetTime(context *gin.Context) {
	body, err := cpu.ServiceGetTime()
	if err == nil {
		context.JSON(200, body)
	} else {
		context.String(400, err.Error())
	}
}
