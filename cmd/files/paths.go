package files

import (
	"log/slog"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
)

var bannedPaths map[string]string = map[string]string{
	"namespaces/Namespace-Common-Position.adoc":         "parser does not support nested tables",
	"service_device_management/PowerSourceCluster.adoc": "parser gets stuck parsing",
	"secure_channel/Discovery.adoc":                     "parser gets stuck parsing",
}

func Paths(filepaths []string) ([]string, error) {
	filtered := make([]string, 0, len(filepaths))
	for _, filepath := range filepaths {
		paths, err := doublestar.FilepathGlob(filepath)
		if err != nil {
			return nil, err
		}
		for _, p := range paths {
			var banned bool
			for ban, reason := range bannedPaths {
				if strings.HasSuffix(p, ban) {
					slog.Warn("Skipping excluded", "file", p, "reason", reason)
					banned = true
				}
			}
			if banned {
				continue
			}
			filtered = append(filtered, p)
		}

	}
	return filtered, nil
}
