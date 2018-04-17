package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor/stream/cpu"
	"github.com/gonitor/gonitor/stream/disk"
	"github.com/gonitor/gonitor/stream/gpu"
	"github.com/gonitor/gonitor/stream/host"
	"github.com/gonitor/gonitor/stream/load"
	"github.com/gonitor/gonitor/stream/memory"
	"github.com/gonitor/gonitor/stream/network"
	"github.com/gonitor/gonitor/stream/runtime"
)

// StreamV1GroupEndPoint .
var StreamV1GroupEndPoint = "/stream/v1"

// StreamV1SetRoutes .
func StreamV1SetRoutes(router *gin.Engine) {

	streamV1 := router.Group(StreamV1GroupEndPoint)
	{
		// Runtime
		streamV1.GET("/runtime/goos/:rateLimit", runtime.RuntimeStreamGetGoOS)

		// Cpu
		streamV1.GET("/cpu/sum/percent/:rateLimit", cpu.CpuStreamGetSumPercent)
		streamV1.GET("/cpu/count/:rateLimit", cpu.CpuStreamGetCount)
		streamV1.GET("/cpu/sum/time/:rateLimit", cpu.CpuStreamGetSumTime)
		streamV1.GET("/cpu/info/:rateLimit", cpu.CpuStreamGetInfo)
		streamV1.GET("/cpu/percent/:rateLimit", cpu.CpuStreamGetPercent)
		streamV1.GET("/cpu/time/:rateLimit", cpu.CpuStreamGetTime)

		// Disk
		streamV1.GET("/disk/usage/:rateLimit", disk.DiskStreamGetUsage)

		// Host
		streamV1.GET("/host/info/:rateLimit", host.HostStreamGetInfo)
		streamV1.GET("/host/temperature/:rateLimit", host.HostStreamGetTemperature)

		// Gpu
		streamV1.GET("/gpu/info/:rateLimit", gpu.GpuStreamGetInfo)

		// Memory
		streamV1.GET("/memory/virtual/:rateLimit", memory.MemoryStreamGetVirtual)
		streamV1.GET("/memory/swap/:rateLimit", memory.MemoryStreamGetSwap)

		// Load
		streamV1.GET("/load/average/:rateLimit", load.LoadStreamGetAverage)
		streamV1.GET("/load/misc/:rateLimit", load.LoadStreamGetMisc)

		// Network
		streamV1.GET("/network/interfaces/:rateLimit", network.NetworkStreamGetInterfaces)
		streamV1.GET("/network/connections/:rateLimit", network.NetworkStreamGetConnections)
	}
}
