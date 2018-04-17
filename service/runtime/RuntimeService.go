package runtime

import (
	"runtime"
)

// ServiceGetGoOS .
func ServiceGetGoOS() string {
	return runtime.GOOS
}
