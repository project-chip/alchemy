package config

import (
	"runtime/debug"
	"strings"
)

var version string

var Tag string

func Version() string {
	if len(version) > 0 {
		return version
	}
	var revision string
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				revision = setting.Value
				break
			}
		}
	}
	if len(revision) > 7 {
		revision = revision[0:7]
	}
	var ver strings.Builder
	if len(Tag) > 0 {
		ver.WriteString(Tag)
		if len(revision) > 0 {
			ver.WriteString(" (")
			ver.WriteString(revision)
			ver.WriteRune(')')
		}
	} else if len(revision) > 0 {
		ver.WriteString(revision)
	} else {
		ver.WriteString("unknown")
	}
	version = ver.String()
	return version
}
