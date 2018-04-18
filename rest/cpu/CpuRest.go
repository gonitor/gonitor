package cpu

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/cpu"
	"github.com/gonitor/gonitor/util"
)

// CpuRestGetSumPercent gets the percentage summary of cpu usage via GET request.
func CpuRestGetSumPercent(context *gin.Context) {
	body, err := cpu.ServiceGetSummaryPercent()
	util.RestHandleResponse(context, body, err)
}

// CpuRestGetCount gets the CPU counts via GET request.
func CpuRestGetCount(context *gin.Context) {
	body, err := cpu.ServiceGetCount()
	util.RestHandleResponse(context, body, err)
}

// CpuRestGetSumTime gets the time summary via GET request.
func CpuRestGetSumTime(context *gin.Context) {
	body, err := cpu.ServiceGetSummaryTime()
	util.RestHandleResponse(context, body, err)
}

// CpuRestGetInfo gets the CPU information via GET request.
func CpuRestGetInfo(context *gin.Context) {
	body, err := cpu.ServiceGetInfo()
	util.RestHandleResponse(context, body, err)
}

// CpuRestGetPercent gets CPU usage in details via GET request.
func CpuRestGetPercent(context *gin.Context) {
	body, err := cpu.ServiceGetPercent()
	util.RestHandleResponse(context, body, err)
}

// CpuRestGetTime gets CPU time information in details via GET request.
func CpuRestGetTime(context *gin.Context) {
	body, err := cpu.ServiceGetTime()
	util.RestHandleResponse(context, body, err)
}
