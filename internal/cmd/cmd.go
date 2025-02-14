package cmd

import (
	"runtime/debug"
)

// BuildInfo for cmd.
func BuildInfo() *debug.BuildInfo {
	info, _ := debug.ReadBuildInfo()

	return info
}
