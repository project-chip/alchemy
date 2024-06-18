package spec

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hasty/alchemy/asciidoc/parse"
	"github.com/hasty/alchemy/internal/pipeline"
)

func ReadFile(path string) (*Doc, error) {

	contents, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer contents.Close()
	return read(contents, path)
}

func Read(contents string, path string) (doc *Doc, err error) {
	return read(strings.NewReader(contents), path)
}

func read(r io.Reader, path string) (doc *Doc, err error) {
	d, err := parse.Reader(path, r)
	if err != nil {
		return nil, fmt.Errorf("read error in %s: %w", path, err)
	}

	doc, err = NewDoc(d, path)
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
