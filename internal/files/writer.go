package files

import (
	"context"
	"os"

	"github.com/project-chip/alchemy/internal/pipeline"
)

type Writer[T string | []byte] interface {
	pipeline.Processor
	SetName(name string)
}

type writer struct {
	name string
}

func (w *writer) Name() string {
	return w.name
}

func (w *writer) SetName(name string) {
	w.name = name
}

func NewWriter[T string | []byte](name string, options Options) Writer[T] {
	if options.DryRun {
		return &DryRun[T]{writer: writer{name: name}}
	}
	if options.Patch {
		return NewPatcher[T](name, os.Stdout)
	}
	return &FileWriter[T]{writer: writer{name: name}}
}

type FileWriter[T string | []byte] struct {
	writer
}

func (sp *FileWriter[T]) Name() string {
	return sp.name
}

func (sp *FileWriter[T]) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (sp *FileWriter[T]) Process(cxt context.Context, input *pipeline.Data[T], index int32, total int32) (outputs []*pipeline.Data[struct{}], extras []*pipeline.Data[T], err error) {
	err = os.WriteFile(input.Path, []byte(input.Content), os.ModeAppend|0644)
	return
}

var _ pipeline.IndividualProcessor[string, struct{}] = &FileWriter[string]{}
