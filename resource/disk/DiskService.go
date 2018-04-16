package disk

import (
	"github.com/shirou/gopsutil/disk"
)

// GetUsage .
func GetUsage() (*disk.UsageStat, error) {
	return disk.Usage("/")
}
