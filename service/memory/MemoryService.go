package memory

import (
	"github.com/shirou/gopsutil/mem"
)

// ServiceGetVirtualMemory gets the virtual memory.
func ServiceGetVirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

// ServiceGetSwapMemory gets the swap memory.
func ServiceGetSwapMemory() (*mem.SwapMemoryStat, error) {
	return mem.SwapMemory()
}
