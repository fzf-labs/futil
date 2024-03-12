package osutil

import (
	"runtime"
)

// GoVersion 获取go版本。例如：“1.18.2”
func GoVersion() string {
	return runtime.Version()[2:]
}
