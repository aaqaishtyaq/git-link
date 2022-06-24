package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// Version would be set at build time using -ldFlag
var (
	Version = "dev"
	Commit  = ""
	Date    = ""
)

const website = "https://github.com/aaqaishtyaq/git-link"

// BuildVersion return version info
func BuildVersion() string {
	result := Version
	if Commit != "" {
		result = fmt.Sprintf("%s\nCommit: %s", result, Commit)
	}
	if Date != "" {
		result = fmt.Sprintf("%s\nBuilt at: %s", result, Date)
	}

	result = fmt.Sprintf("%s\nGOOS: %s\nGOARCH: %s", result, runtime.GOOS, runtime.GOARCH)

	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		result = fmt.Sprintf("%s\nModule version: %s, Checksum: %s", result, info.Main.Version, info.Main.Sum)
	}

	return result + "\n\nFor contribution, Please visit " + website
}
