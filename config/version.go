package config

import "runtime/debug"

var version string

func Version() string {
	if len(version) > 0 {
		return version
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				version = setting.Value
				return version
			}
		}
	}
	return "unknown"
}
