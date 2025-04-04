package paths

import (
	"context"
	"log/slog"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
)

var bannedPaths map[string]string

func Expand(filepaths []string) ([]string, error) {
	out := make([]string, 0, len(filepaths))
	for _, filepath := range filepaths {
		paths, err := doublestar.FilepathGlob(filepath)
		if err != nil {
			return nil, err
		}
		out = append(out, FilterBanned(paths...)...)
	}
	return out, nil
}

func NewTargeter(paths ...string) func(cxt context.Context) ([]string, error) {
	return func(cxt context.Context) ([]string, error) {
		return Expand(paths)
	}
}

func FilterBanned(paths ...string) []string {
	var filtered = make([]string, 0, len(paths))
	for _, p := range paths {
		var banned bool
		for ban, reason := range bannedPaths {
			if strings.HasSuffix(p, ban) {
				slog.Debug("Skipping excluded", "file", p, "reason", reason)
				banned = true
			}
		}
		if banned {
			continue
		}
		filtered = append(filtered, p)
	}
	return filtered
}
