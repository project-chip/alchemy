package config

import (
	"runtime/debug"
	"strings"
)

var version string

// Set by build flag, most commonly in .github/workflows/main.yml
var Tag string

func Version() string {
	if len(version) > 0 {
		return version
	}
	info, ok := debug.ReadBuildInfo()
	if ok {
		if len(info.Main.Version) > 0 {
			return info.Main.Version
		}
		var revision string
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				revision = setting.Value
				break
			}
		}
		if len(revision) > 7 {
			revision = revision[0:7]
		}
		if len(Tag) > 0 {
			var ver strings.Builder
			ver.WriteString(Tag)
			if len(revision) > 0 {
				ver.WriteString(" (")
				ver.WriteString(revision)
				ver.WriteRune(')')
			}
			return ver.String()
		} else if len(revision) > 0 {
			return revision
		}
	}
	return "unknown"
}
