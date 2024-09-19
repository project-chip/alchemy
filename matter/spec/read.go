package spec

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/pipeline"
)

func ReadFile(path string, rootPath string) (*Doc, error) {

	contents, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer contents.Close()
	b, err := io.ReadAll(contents)
	if err != nil {
		return nil, err
	}
	return read(b, path, rootPath)
}

func readString(contents string, path string, rootPath string) (doc *Doc, err error) {
	return read([]byte(contents), path, rootPath)
}

func read(b []byte, path string, rootPath string) (doc *Doc, err error) {
	d, err := parse.Bytes(path, b)
	if err != nil {
		return nil, fmt.Errorf("read error in %s: %w", path, err)
	}

	var p asciidoc.Path
	p, err = NewDocPath(path, rootPath)
	if err != nil {
		return nil, err
	}

	doc, err = newDoc(d, p)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

type Reader struct {
	name     string
	rootPath string
}

func NewReader(name string, rootPath string) (Reader, error) {
	if !filepath.IsAbs(rootPath) {
		var err error
		rootPath, err = filepath.Abs(rootPath)
		if err != nil {
			return Reader{}, err
		}
	}
	return Reader{name: name, rootPath: rootPath}, nil
}

func (r Reader) Name() string {
	return r.name
}

func (r Reader) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (r Reader) Process(cxt context.Context, input *pipeline.Data[struct{}], index int32, total int32) (outputs []*pipeline.Data[*Doc], extras []*pipeline.Data[struct{}], err error) {
	var doc *Doc
	doc, err = ReadFile(input.Path, r.rootPath)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[*Doc]{Path: input.Path, Content: doc})
	return
}

type StringReader struct {
	name     string
	rootPath string
}

func NewStringReader(name string, rootPath string) (StringReader, error) {
	if !filepath.IsAbs(rootPath) {
		var err error
		rootPath, err = filepath.Abs(rootPath)
		if err != nil {
			return StringReader{}, err
		}
	}
	return StringReader{name: name, rootPath: rootPath}, nil
}

func (r StringReader) Name() string {
	return r.name
}

func (r StringReader) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (r StringReader) Process(cxt context.Context, input *pipeline.Data[string], index int32, total int32) (outputs []*pipeline.Data[*Doc], extras []*pipeline.Data[string], err error) {
	var doc *Doc
	doc, err = readString(input.Content, input.Path, r.rootPath)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[*Doc]{Path: input.Path, Content: doc})
	return
}
