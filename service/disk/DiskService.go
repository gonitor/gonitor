package disk

import (
	"github.com/shirou/gopsutil/disk"
)

// ServiceGetUsage .
func ServiceGetUsage() (*disk.UsageStat, error) {
	return disk.Usage("/")
}
