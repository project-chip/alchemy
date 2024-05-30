package files

import (
	"context"
	"os"
	"path/filepath"

	"github.com/hasty/alchemy/internal/pipeline"
)

type PathFilter[T any] struct {
	paths []string
}

func NewPathFilter[T any](paths []string) *PathFilter[T] {
	return &PathFilter[T]{paths: paths}
}

func (p PathFilter[T]) Name() string {
	return ""
}

func (p PathFilter[T]) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeCollective
}

func (p PathFilter[T]) Process(cxt context.Context, inputs []*pipeline.Data[T]) (outputs []*pipeline.Data[T], err error) {
	if len(p.paths) == 0 {
		return inputs, nil
	}
	pathMap := make(map[string]struct{})
	for _, p := range p.paths {
		pathMap[filepath.Base(p)] = struct{}{}
	}
	stats := make([]os.FileInfo, 0, len(p.paths))
	for _, p := range p.paths {
		var fi os.FileInfo
		fi, err = os.Stat(p)
		if err != nil {
			return
		}
		stats = append(stats, fi)
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
			}
		}
	}
	return
}
