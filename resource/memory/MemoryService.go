package memory

import (
	"github.com/shirou/gopsutil/mem"
)

// GetVirtualMemory .
func GetVirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

// GetSwapMemory .
func GetSwapMemory() (*mem.SwapMemoryStat, error) {
	return mem.SwapMemory()
}
