package paths

import (
	"context"
	"log/slog"
	"os"
	"slices"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type Filter[T any] struct {
	specRoot string
	paths    []string
}

func NewFilter[T any](specRoot string, paths []string) *Filter[T] {
	return &Filter[T]{specRoot: specRoot, paths: paths}
}

func (p Filter[T]) Name() string {
	return ""
}

func (p *Filter[T]) Process(cxt context.Context, inputs []*pipeline.Data[T]) (outputs []*pipeline.Data[T], err error) {
	if len(p.paths) == 0 {
		return inputs, nil
	}
	stats := make([]os.FileInfo, 0, len(p.paths))
	filteredFiles := make(map[os.FileInfo]string, len(p.paths))
	for _, p := range p.paths {
		var expandedPaths []string
		expandedPaths, err = doublestar.FilepathGlob(p)
		if err != nil {
			return nil, err
		}
		for _, ep := range expandedPaths {
			var fi os.FileInfo
			fi, err = os.Stat(ep)
			if err != nil {
				return
			}
			stats = append(stats, fi)
			filteredFiles[fi] = ep
		}
	}
	for _, d := range inputs {
		var fi os.FileInfo
		fi, err = os.Stat(d.Path)
		if err != nil {
			return
		}
		for i, ofi := range stats {
			if os.SameFile(fi, ofi) {
				outputs = append(outputs, d)
				if len(stats) <= 1 {
					stats = nil
					return
				}
				stats[i] = stats[len(stats)-1]
				stats = stats[:len(stats)-1]
				delete(filteredFiles, ofi)
			}
		}
	}

	var ignoredFiles []string
	for _, path := range filteredFiles {
		specPath, err := asciidoc.NewPath(path, p.specRoot)
		if err == nil {
			if errata.GetSpec(specPath.Relative).UtilityInclude {
				continue
			}
		}
		switch specPath.Relative { // We can ignore these cover files, as they never link to anything
		case "src/cover-main.adoc",
			"src/cover-appclusters.adoc",
			"src/cover-device_library.adoc",
			"src/cover-standard_namespaces.adoc":
			continue
		}
		ignoredFiles = append(ignoredFiles, specPath.Relative)
	}
	if len(ignoredFiles) > 0 {
		slices.Sort(ignoredFiles)
		docRootLogs := make([]slog.Attr, 0, len(errata.DocRoots))
		for _, p := range errata.DocRoots {
			docRootLogs = append(docRootLogs, slog.String("path", p))
		}
		ignoreLogs := make([]slog.Attr, 0, len(ignoredFiles))
		for _, p := range ignoredFiles {
			ignoreLogs = append(ignoreLogs, slog.String("path", p))
		}
		slog.Warn("Some files were ignored since they were not referenced by any document roots", slog.Any("docRoots", slog.GroupValue(docRootLogs...)), slog.Any("ignoredPaths", slog.GroupValue(ignoreLogs...)))
	}
	return
}
