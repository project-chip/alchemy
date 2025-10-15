package spec

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type LibraryBuilder struct {
	specRoot string
	errata   *errata.Collection
	config   *config.Config
}

func NewLibraryBuilder(specRoot string, config *config.Config, errata *errata.Collection) *LibraryBuilder {
	b := &LibraryBuilder{
		specRoot: specRoot,
		config:   config,
		errata:   errata,
	}
	return b
}

func (lb LibraryBuilder) Name() string {
	return "Grouping spec documents"
}

func (lb *LibraryBuilder) Process(cxt context.Context, inputs []*pipeline.Data[*asciidoc.Document]) (outputs []*pipeline.Data[*Library], err error) {

	docCache := cacheFromPipeline(lb.specRoot, inputs)

	for _, libraryConfig := range lb.config.Libraries {
		root, ok := docCache.cache.Load(libraryConfig.Root)
		if !ok {
			slog.Warn("doc root not found", "root", libraryConfig.Root)
			continue
		}
		outputs = append(outputs, pipeline.NewData(root.Path.Relative, NewLibrary(root, libraryConfig, lb.errata, docCache)))
	}

	return
}
