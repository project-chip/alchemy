package files

import (
	"context"
	"fmt"
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
	return pipeline.ProcessorTypeSerial
}

func (p PathFilter[T]) Process(cxt context.Context, input *pipeline.Data[T], index int32, total int32) (outputs []*pipeline.Data[T], extras []*pipeline.Data[T], err error) {
	err = fmt.Errorf("path filtering must be done serially")
	return
}

func (p PathFilter[T]) ProcessAll(cxt context.Context, inputs []*pipeline.Data[T]) (outputs []*pipeline.Data[T], err error) {
	if len(p.paths) == 0 {
		return inputs, nil
	}
	pathMap := make(map[string]struct{})
	for _, p := range p.paths {
		pathMap[filepath.Base(p)] = struct{}{}
	}
	for _, d := range inputs {
		p := filepath.Base(d.Path)
		if _, ok := pathMap[p]; ok {
			outputs = append(outputs, d)
			delete(pathMap, p)
		}
	}
	return
}
