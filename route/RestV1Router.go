package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/rest/cpu"
	"github.com/gonitor/gonitor/rest/disk"
	"github.com/gonitor/gonitor/rest/gpu"
	"github.com/gonitor/gonitor/rest/host"
	"github.com/gonitor/gonitor/rest/load"
	"github.com/gonitor/gonitor/rest/memory"
	"github.com/gonitor/gonitor/rest/network"
	"github.com/gonitor/gonitor/rest/runtime"
)

// RestV1GroupEndPoint .
var RestV1GroupEndPoint = "/rest/v1"

// RestV1SetRoutes .
func RestV1SetRoutes(router *gin.Engine) {

	restV1 := router.Group(RestV1GroupEndPoint)
	{
		// Runtime
		restV1.GET("/runtime/goos", runtime.RuntimeRestGetGoOS)

		// Cpu
		restV1.GET("/cpu/sum/percent", cpu.CpuRestGetSumPercent)
		restV1.GET("/cpu/count", cpu.CpuRestGetCount)
		restV1.GET("/cpu/sum/time", cpu.CpuRestGetSumTime)
		restV1.GET("/cpu/info", cpu.CpuRestGetInfo)
		restV1.GET("/cpu/percent", cpu.CpuRestGetPercent)
		restV1.GET("/cpu/time", cpu.CpuRestGetTime)

		// Disk
		restV1.GET("/disk/usage", disk.DiskRestGetUsage)

		// Host
		restV1.GET("/host/info", host.HostRestGetInfo)
		restV1.GET("/host/temperature", host.HostRestGetTemperature)

		// Gpu
		restV1.GET("/gpu/info", gpu.GpuRestGetInfo)

		// Memory
		restV1.GET("/memory/virtual", memory.MemoryRestGetVirtual)
		restV1.GET("/memory/swap", memory.MemoryRestGetSwap)

		// Load
		restV1.GET("/load/average", load.LoadRestGetAverage)
		restV1.GET("/load/misc", load.LoadRestGetMisc)

		// Network
		restV1.GET("/network/interfaces", network.NetworkRestGetInterfaces)
		restV1.GET("/network/connections", network.NetworkRestGetConnections)
	}
}
