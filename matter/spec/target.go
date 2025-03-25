package spec

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/project-chip/alchemy/internal/paths"
)

func getSpecPaths(specRoot string) (specPaths []string, err error) {
	srcRoot := filepath.Join(specRoot, "/src/")
	err = filepath.WalkDir(srcRoot, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".adoc" && !strings.HasSuffix(path, "-draft.adoc") {
			specPaths = append(specPaths, path)
		}
		return nil
	})
	specPaths = paths.FilterBanned(specPaths...)
	return
}

func Targeter(specRoot string) func(cxt context.Context) ([]string, error) {
	return func(cxt context.Context) ([]string, error) {
		return getSpecPaths(specRoot)
	}
}
