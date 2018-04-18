package load

import (
	"github.com/shirou/gopsutil/load"
)

// ServiceGetAverage gets the load average.
func ServiceGetAverage() (*load.AvgStat, error) {
	return load.Avg()
}

// GetMisc gets the load misc.
func GetMisc() (*load.MiscStat, error) {
	return load.Misc()
}
