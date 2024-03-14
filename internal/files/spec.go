package files

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter"
)

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
	return pipeline.ProcessorTypeCollective
}

func (sp *SpecParser) Process(cxt context.Context, inputs []*pipeline.Data[*ascii.Doc]) (outputs []*pipeline.Data[*ascii.Doc], err error) {
	docs := make([]*ascii.Doc, 0, len(inputs))
	for _, i := range inputs {
		docs = append(docs, i.Content)
	}
	sp.Spec, err = ascii.BuildSpec(docs)
	outputs = inputs
	return
}
