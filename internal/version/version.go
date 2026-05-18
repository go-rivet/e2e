package version

import (
	"fmt"
)

var (
	Version   = "dev"
	CommitSHA = "unknown"
	BuildTime = "unknown"
)

func GetVersion() string {
	return Version
}

func GetVersionWithBuildInfo() string {
	return fmt.Sprintf("Rivet E2E: %s\nCommit:    %s\nBuilt:     %s", Version, CommitSHA, BuildTime)
}
