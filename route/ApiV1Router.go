package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/resource/cpu"
	"github.com/gonitor/gonitor/resource/disk"
	"github.com/gonitor/gonitor/resource/gpu"
	"github.com/gonitor/gonitor/resource/home"
	"github.com/gonitor/gonitor/resource/host"
	"github.com/gonitor/gonitor/resource/load"
	"github.com/gonitor/gonitor/resource/memory"
	"github.com/gonitor/gonitor/resource/network"
	"github.com/gonitor/gonitor/resource/runtime"
)

// ApiV1SetRoutes .
func ApiV1SetRoutes(router *gin.Engine) {

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/home/index", home.HomeGetIndex)
		// Runtime
		apiV1.GET("/runtime/goos", runtime.RuntimeGetGoOs)

		// Cpu
		apiV1.GET("/cpu/sum/percent", cpu.CpuGetSumPercent)
		apiV1.GET("/cpu/count", cpu.CpuGetCount)
		apiV1.GET("/cpu/sum/time", cpu.CpuGetSumTime)
		apiV1.GET("/cpu/info", cpu.CpuGetInfo)
		apiV1.GET("/cpu/percent", cpu.CpuGetPercent)
		apiV1.GET("/cpu/time", cpu.CpuGetTime)

		// Disk
		apiV1.GET("/disk/usage", disk.DiskGetUsage)

		// Host
		apiV1.GET("/host/info", host.HostGetInfo)
		apiV1.GET("/host/temperature", host.HostGetTemperature)

		// Gpu
		apiV1.GET("/gpu/info", gpu.GpuGetInfo)

		// Memory
		apiV1.GET("/memory/virtual", memory.MemoryGetVirtual)
		apiV1.GET("/memory/swap", memory.MemoryGetSwap)

		// Load
		apiV1.GET("/load/average", load.LoadGetAverage)
		apiV1.GET("/load/misc", load.LoadGetMisc)

		// Network
		apiV1.GET("/network/interfaces", network.NetworkGetInterfaces)
		apiV1.GET("/network/connections", network.NetworkGetConnections)
	}
}
