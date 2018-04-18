package runtime

import (
	"runtime"
)

// ServiceGetGoOS gets the operating system (OS) of Go runtime.
func ServiceGetGoOS() string {
	return runtime.GOOS
}
