package spec

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type LibraryBuilder struct {
	specRoot string
}

func NewLibraryBuilder(specRoot string) *LibraryBuilder {
	b := &LibraryBuilder{
		specRoot: specRoot,
	}
	return b
}

func (lb LibraryBuilder) Name() string {
	return "Grouping spec documents"
}

func (lb *LibraryBuilder) Process(cxt context.Context, inputs []*pipeline.Data[*asciidoc.Document]) (outputs []*pipeline.Data[*Library], err error) {

	docCache := cacheFromPipeline(lb.specRoot, inputs)

	for _, docRoot := range errata.DocRoots {
		root, ok := docCache.cache.Load(docRoot)
		if !ok {
			slog.Warn("doc root not found", "root", docRoot)
			continue
		}
		outputs = append(outputs, pipeline.NewData(root.Path.Relative, NewLibrary(root, docCache)))
	}
	return
}
