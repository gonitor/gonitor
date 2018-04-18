package cpu

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/cpu"
	"github.com/gonitor/gonitor/util"
)

// CpuStreamGetSumPercent streams the percentage summary of cpu usage via GET request.
func CpuStreamGetSumPercent(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := cpu.ServiceGetSummaryPercent()
		return util.StreamHandleResponse(context, body, err, "CpuStreamGetSumPercent")
	})
}

// CpuStreamGetCount streams the CPU counts via GET request.
func CpuStreamGetCount(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := cpu.ServiceGetCount()
		return util.StreamHandleResponse(context, body, err, "CpuStreamGetCount")
	})
}

// CpuStreamGetSumTime streams the time summary via GET request.
func CpuStreamGetSumTime(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := cpu.ServiceGetSummaryTime()
		return util.StreamHandleResponse(context, body, err, "CpuStreamGetSumTime")
	})
}

// CpuStreamGetInfo streams the CPU information via GET request.
func CpuStreamGetInfo(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := cpu.ServiceGetInfo()
		return util.StreamHandleResponse(context, body, err, "CpuStreamGetInfo")
	})
}

// CpuStreamGetPercent streams CPU usage in details via GET request.
func CpuStreamGetPercent(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := cpu.ServiceGetPercent()
		return util.StreamHandleResponse(context, body, err, "CpuStreamGetPercent")
	})
}

// CpuStreamGetTime streams CPU time information in details via GET request.
func CpuStreamGetTime(context *gin.Context) {
	rateLimit, convertErr := util.ConvertStringToTimeDuration(context.Param("rateLimit"))
	context.Stream(func(w io.Writer) bool {
		if convertErr != nil {
			return false
		}
		time.Sleep(rateLimit * time.Second)
		body, err := cpu.ServiceGetTime()
		return util.StreamHandleResponse(context, body, err, "CpuStreamGetTime")
	})
}
