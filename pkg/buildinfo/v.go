package buildinfo

import "time"

var (
	GitVersion string
	BuildDate  time.Time
)

const (
	DevVersion = "v0.0.1-draft"
)
