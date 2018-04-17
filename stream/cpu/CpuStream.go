package cpu

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/service/cpu"
	"github.com/gonitor/gonitor/util"
)

// CpuStreamGetSumPercent .
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

// CpuStreamGetCount .
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

// CpuStreamGetSumTime .
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

// CpuStreamGetInfo .
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

// CpuStreamGetPercent .
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

// CpuStreamGetTime .
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
