package spec

import (
	"context"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
)

func Parse(cxt context.Context, parserOptions ParserOptions, processingOptions pipeline.ProcessingOptions, builderOptions []BuilderOption, attributes []asciidoc.AttributeName) (specification *Specification, specDocs DocSet, err error) {

	specDocs, err = LoadSpecDocs(cxt, parserOptions, processingOptions)
	if err != nil {
		return
	}

	specification, specDocs, err = Build(cxt, parserOptions, processingOptions, builderOptions, specDocs, attributes)
	return
}

func Build(cxt context.Context, parserOptions ParserOptions, processingOptions pipeline.ProcessingOptions, builderOptions []BuilderOption, docs DocSet, attributes []asciidoc.AttributeName) (specification *Specification, specDocs DocSet, err error) {

	var ec *errata.Collection
	ec, err = errata.LoadErrataConfig(parserOptions.Root)
	if err != nil {
		return
	}
	var libraries LibrarySet
	libraries, err = pipeline.Collective(cxt, processingOptions, NewLibraryBuilder(parserOptions.Root, ec), docs)
	if err != nil {
		return
	}
	var preparser *LibraryParser
	preparser, err = NewLibraryParser(parserOptions.Root, attributes)
	if err != nil {
		return
	}

	_, err = pipeline.Parallel(cxt, pipeline.ProcessingOptions{Serial: true}, preparser, libraries)
	if err != nil {
		return
	}

	specBuilder := NewBuilder(parserOptions.Root, ec, builderOptions...)
	specDocs, err = pipeline.Collective(cxt, processingOptions, &specBuilder, libraries)
	if err != nil {
		return
	}

	unusedDocs := make(map[*asciidoc.Document]int)
	libraries.Range(func(key string, library *pipeline.Data[*Library]) bool {
		library.Content.cache.cache.Range(func(key string, value *asciidoc.Document) bool {
			if !library.Content.IsUtilityPath(value.Path) {
				unusedDocs[value] = unusedDocs[value] + 1
			}
			return true
		})
		return true
	})
	libraryCount := libraries.Size()
	for doc, unusedCount := range unusedDocs {
		if unusedCount == libraryCount {
			specBuilder.Spec.UnusedDocs = append(specBuilder.Spec.UnusedDocs, doc)
		}
	}
	specification = specBuilder.Spec
	return
}

func LoadSpecDocs(cxt context.Context, parserOptions ParserOptions, processingOptions pipeline.ProcessingOptions) (specDocs DocSet, err error) {
	var specReader Reader
	specReader, err = NewReader(parserOptions)
	if err != nil {
		return
	}

	specTargeter := Targeter(parserOptions.Root)

	var specPaths pipeline.Paths
	specPaths, err = pipeline.Start(cxt, specTargeter)
	if err != nil {
		return
	}

	specDocs, err = pipeline.Parallel(cxt, processingOptions, specReader, specPaths)
	return
}
