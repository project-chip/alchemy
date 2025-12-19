package regen

import (
	"context"
	"io/fs"
	"path/filepath"
)

func Targeter(sdkRoot string) func(cxt context.Context) ([]string, error) {
	return func(cxt context.Context) ([]string, error) {
		return getZapPaths(sdkRoot)
	}
}

func getZapPaths(sdkRoot string) (zapPaths []string, err error) {
	zapPaths = append(zapPaths, filepath.Join(sdkRoot, "src/controller/data_model/controller-clusters.zap"))
	srcRoot := filepath.Join(sdkRoot, "/examples/")
	err = filepath.WalkDir(srcRoot, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".zap" {
			zapPaths = append(zapPaths, path)
		}
		return nil
	})
	return
}
