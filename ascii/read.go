package ascii

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
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
		contents = doorLockPattern.ReplaceAllString(contents, " ")
	}
	return contents, nil
}

func ReadFile(path string, settings ...configuration.Setting) (*Doc, error) {

	contents, err := readFile(path)
	if err != nil {
		return nil, err
	}
	return Read(contents, path)
}

func Read(contents string, path string) (doc *Doc, err error) {

	config := configuration.NewConfiguration(configuration.WithFilename(path))
	config.IgnoreIncludes = true

	var d *types.Document

	d, err = ParseDocument(strings.NewReader(contents), config, parser.MaxExpressions(2000000))

	if err != nil {
		return nil, fmt.Errorf("failed parse: %w", err)
	}

	doc, err = NewDoc(d)
	if err != nil {
		return nil, err
	}
	doc.Path = path

	PatchUnrecognizedReferences(doc)

	return doc, nil
}

type Reader struct {
	name          string
	asciiSettings []configuration.Setting
	options       pipeline.Options
}

func NewReader(name string, options pipeline.Options, asciiSettings ...configuration.Setting) Reader {
	return Reader{name: name, options: options, asciiSettings: asciiSettings}
}

func (r Reader) Name() string {
	return r.name
}

func (r Reader) Type() pipeline.ProcessorType {
	return r.options.DefaultProcessorType()
}

func (r Reader) Process(cxt context.Context, input *pipeline.Data[struct{}], index int32, total int32) (outputs []*pipeline.Data[*Doc], extras []*pipeline.Data[struct{}], err error) {
	var doc *Doc
	doc, err = ReadFile(input.Path, r.asciiSettings...)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[*Doc]{Path: input.Path, Content: doc})
	return
}

func (r Reader) ProcessAll(cxt context.Context, inputs []*pipeline.Data[struct{}]) (outputs []*pipeline.Data[*Doc], err error) {
	for _, input := range inputs {
		var doc *Doc
		fmt.Fprintf(os.Stderr, "Reading %s...\n", input.Path)
		doc, err = ReadFile(input.Path, r.asciiSettings...)
		if err != nil {
			return
		}
		outputs = append(outputs, &pipeline.Data[*Doc]{Path: input.Path, Content: doc})
	}
	return
}
