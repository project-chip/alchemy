package spec

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/pipeline"
)

func ReadFile(path string) (*Doc, error) {

	contents, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer contents.Close()
	b, err := io.ReadAll(contents)
	if err != nil {
		return nil, err
	}
	return read(b, path)
}

func Read(contents string, path string) (doc *Doc, err error) {
	return read([]byte(contents), path)
}

func read(b []byte, path string) (doc *Doc, err error) {
	d, err := parse.Bytes(path, b)
	if err != nil {
		return nil, fmt.Errorf("read error in %s: %w", path, err)
	}

	var p Path
	p, err = NewPath(path)
	if err != nil {
		return nil, err
	}

	doc, err = NewDoc(d, p)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

type Reader struct {
	name string
}

func NewReader(name string) Reader {
	return Reader{name: name}
}

func (r Reader) Name() string {
	return r.name
}

func (r Reader) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (r Reader) Process(cxt context.Context, input *pipeline.Data[struct{}], index int32, total int32) (outputs []*pipeline.Data[*Doc], extras []*pipeline.Data[struct{}], err error) {
	var doc *Doc
	doc, err = ReadFile(input.Path)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[*Doc]{Path: input.Path, Content: doc})
	return
}

type StringReader struct {
	name string
}

func NewStringReader(name string) StringReader {
	return StringReader{name: name}
}

func (r StringReader) Name() string {
	return r.name
}

func (r StringReader) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (r StringReader) Process(cxt context.Context, input *pipeline.Data[string], index int32, total int32) (outputs []*pipeline.Data[*Doc], extras []*pipeline.Data[string], err error) {
	var doc *Doc
	doc, err = Read(input.Content, input.Path)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[*Doc]{Path: input.Path, Content: doc})
	return
}
