package parse

import (
	"context"
	"os"
	"path/filepath"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type ElementsReader struct {
	name string
	Root string
}

func NewElementsReader(name string, rootPath string) (ElementsReader, error) {
	if !filepath.IsAbs(rootPath) {
		var err error
		rootPath, err = filepath.Abs(rootPath)
		if err != nil {
			return ElementsReader{}, err
		}
	}
	return ElementsReader{name: name, Root: rootPath}, nil
}

func (r ElementsReader) Name() string {
	return r.name
}

func (r ElementsReader) Process(cxt context.Context, input *pipeline.Data[struct{}], index int32, total int32) (outputs []*pipeline.Data[asciidoc.Elements], extras []*pipeline.Data[struct{}], err error) {
	var elements asciidoc.Elements
	var file *os.File
	file, err = os.Open(input.Path)
	if err != nil {
		return
	}
	defer file.Close()
	elements, err = Elements(input.Path, file)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[asciidoc.Elements]{Path: input.Path, Content: elements})
	return
}
