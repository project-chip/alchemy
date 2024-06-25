package spec

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/project-chip/alchemy/internal/files"
)

func getSpecPaths(specRoot string) (paths []string, err error) {
	srcRoot := filepath.Join(specRoot, "/src/")
	err = filepath.WalkDir(srcRoot, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".adoc" && !strings.HasSuffix(path, "-draft.adoc") {
			paths = append(paths, path)
		}
		return nil
	})
	paths = files.FilterBannedPaths(paths...)
	return
}

func Targeter(specRoot string) func(cxt context.Context) ([]string, error) {
	return func(cxt context.Context) ([]string, error) {
		return getSpecPaths(specRoot)
	}
}
