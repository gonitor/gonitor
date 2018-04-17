package memory

import (
	"github.com/shirou/gopsutil/mem"
)

// ServiceGetVirtualMemory .
func ServiceGetVirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

// ServiceGetSwapMemory .
func ServiceGetSwapMemory() (*mem.SwapMemoryStat, error) {
	return mem.SwapMemory()
}
