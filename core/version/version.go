package version

import (
	"runtime"
)

// Build information. Populated at build-time.
var (
	Version   string
	Branch    string
	BuildUser string
	BuildDate string
	GoVersion = runtime.Version()
)
