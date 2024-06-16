package spec

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/internal/pipeline"
)

func readFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	contents := string(file)

	if filepath.Base(path) == "DoorLock.adoc" {
		var doorLockPattern = regexp.MustCompile(`\n+\s*[^&\n]+&#8224;\s+`)
		contents = doorLockPattern.ReplaceAllString(contents, "\n")
	}
	return contents, nil
}

func ReadFile(path string, attributes ...asciidoc.AttributeName) (*Doc, error) {

	contents, err := readFile(path)
	if err != nil {
		return nil, err
	}
	return Read(contents, path, attributes...)
}

func Read(contents string, path string, attributes ...asciidoc.AttributeName) (doc *Doc, err error) {

	var d *asciidoc.Document

	d, err = ParseDocument(strings.NewReader(contents), path, attributes...)

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
