package disk

import (
	"github.com/shirou/gopsutil/disk"
)

// ServiceGetUsage gets the disk usage.
func ServiceGetUsage() (*disk.UsageStat, error) {
	return disk.Usage("/")
}
