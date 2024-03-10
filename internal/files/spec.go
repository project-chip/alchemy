package files

import (
	"context"
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"strings"
	"sync"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter"
)

func LoadSpec(cxt context.Context, specRoot string, filesOptions Options, asciiSettings []configuration.Setting) (spec *matter.Spec, docs []*ascii.Doc, err error) {
	var lock sync.Mutex
	asciiSettings = append(ascii.GithubSettings(), asciiSettings...)

	var specPaths []string
	specPaths, err = getSpecPaths(specRoot)
	if err != nil {
		return
	}

	err = Process(cxt, specPaths, func(cxt context.Context, path string, index, total int) error {

		doc, err := ascii.ParseFile(path, asciiSettings...)
		if err != nil {
			return err
		}

		lock.Lock()
		docs = append(docs, doc)
		lock.Unlock()
		if filesOptions.Serial {
			slog.InfoContext(cxt, "Parsed spec adoc", "file", path)
		}
		return nil
	}, filesOptions)

	if err != nil {
		return
	}

	slog.InfoContext(cxt, "Building spec...")
	spec, err = ascii.BuildSpec(docs)

	return
}

func getSpecPaths(specRoot string) (paths []string, err error) {
	srcRoot := filepath.Join(specRoot, "/src/")
	err = filepath.WalkDir(srcRoot, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".adoc" && !strings.HasSuffix(path, "-draft.adoc") {
			paths = append(paths, path)
		}
		return nil
	})
	return
}

func SplitSpec(docs []*ascii.Doc) (map[matter.DocType][]*ascii.Doc, error) {
	byType := make(map[matter.DocType][]*ascii.Doc)
	for _, d := range docs {
		docType, err := d.DocType()
		if err != nil {
			return nil, err
		}
		byType[docType] = append(byType[docType], d)
	}
	return byType, nil
}

func SpecTargeter(specRoot string) func(cxt context.Context) ([]string, error) {
	return func(cxt context.Context) ([]string, error) {
		return getSpecPaths(specRoot)
	}
}

type SpecParser struct {
	Spec *matter.Spec
}

func (sp SpecParser) Name() string {
	return "Loading spec"
}

func (sp SpecParser) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeSerial
}

func (sp *SpecParser) Process(cxt context.Context, input *pipeline.Data[*ascii.Doc], index int32, total int32) (outputs []*pipeline.Data[*ascii.Doc], extras []*pipeline.Data[*ascii.Doc], err error) {
	err = fmt.Errorf("spec loading must be done serially")
	return
}

func (sp *SpecParser) ProcessAll(cxt context.Context, inputs []*pipeline.Data[*ascii.Doc]) (outputs []*pipeline.Data[*ascii.Doc], err error) {
	docs := make([]*ascii.Doc, 0, len(inputs))
	for _, i := range inputs {
		docs = append(docs, i.Content)
	}
	sp.Spec, err = ascii.BuildSpec(docs)
	outputs = inputs
	return
}
