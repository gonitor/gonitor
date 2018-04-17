package cpu

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/cpu"
	"github.com/gonitor/gonitor/util"
)

// CpuRestGetSumPercent .
func CpuRestGetSumPercent(context *gin.Context) {
	body, err := cpu.ServiceGetSummaryPercent()
	util.RestHandleResponse(context, body, err)
}

// CpuRestGetCount .
func CpuRestGetCount(context *gin.Context) {
	body, err := cpu.ServiceGetCount()
	util.RestHandleResponse(context, body, err)
}

// CpuRestGetSumTime .
func CpuRestGetSumTime(context *gin.Context) {
	body, err := cpu.ServiceGetSummaryTime()
	util.RestHandleResponse(context, body, err)
}

// CpuRestGetInfo .
func CpuRestGetInfo(context *gin.Context) {
	body, err := cpu.ServiceGetInfo()
	util.RestHandleResponse(context, body, err)
}

// CpuRestGetPercent .
func CpuRestGetPercent(context *gin.Context) {
	body, err := cpu.ServiceGetPercent()
	util.RestHandleResponse(context, body, err)
}

// CpuRestGetTime .
func CpuRestGetTime(context *gin.Context) {
	body, err := cpu.ServiceGetTime()
	util.RestHandleResponse(context, body, err)
}
