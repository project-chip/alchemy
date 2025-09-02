package spec

import (
	"context"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
)

func Parse(cxt context.Context, parserOptions ParserOptions, processingOptions pipeline.ProcessingOptions, builderOptions []BuilderOption, attributes []asciidoc.AttributeName) (specification *Specification, specDocs DocSet, err error) {

	var libraries pipeline.Map[string, *pipeline.Data[*Library]]
	libraries, specDocs, err = BuildLibraries(cxt, parserOptions, processingOptions)
	if err != nil {
		return
	}

	specification, specDocs, err = Build(cxt, parserOptions, processingOptions, builderOptions, libraries, attributes)
	return
}

func Build(cxt context.Context, parserOptions ParserOptions, processingOptions pipeline.ProcessingOptions, builderOptions []BuilderOption, libraries LibrarySet, attributes []asciidoc.AttributeName) (specification *Specification, specDocs DocSet, err error) {
	err = PreParse(cxt, parserOptions, processingOptions, libraries, attributes)
	if err != nil {
		return
	}

	specBuilder := NewBuilder(parserOptions.Root, builderOptions...)
	specDocs, err = pipeline.Collective(cxt, processingOptions, &specBuilder, libraries)
	if err != nil {
		return
	}

	specification = specBuilder.Spec
	return
}

func PreParse(cxt context.Context, parserOptions ParserOptions, processingOptions pipeline.ProcessingOptions, libraries LibrarySet, attributes []asciidoc.AttributeName) (err error) {
	var preparser *LibraryParser
	preparser, err = NewLibraryParser(parserOptions.Root, attributes)
	if err != nil {
		return
	}

	_, err = pipeline.Parallel(cxt, pipeline.ProcessingOptions{Serial: true}, preparser, libraries)
	return
}

func BuildLibraries(cxt context.Context, parserOptions ParserOptions, processingOptions pipeline.ProcessingOptions) (libraries LibrarySet, specDocs DocSet, err error) {
	var specReader Reader
	specReader, err = NewReader(parserOptions)
	if err != nil {
		return
	}

	err = errata.LoadErrataConfig(parserOptions.Root)
	if err != nil {
		return
	}

	specTargeter := Targeter(parserOptions.Root)

	var specPaths pipeline.Paths
	specPaths, err = pipeline.Start(cxt, specTargeter)
	if err != nil {
		return
	}

	var docs pipeline.Map[string, *pipeline.Data[*asciidoc.Document]]
	docs, err = pipeline.Parallel(cxt, processingOptions, specReader, specPaths)
	if err != nil {
		return
	}

	libraries, err = pipeline.Collective(cxt, processingOptions, NewLibraryBuilder(parserOptions.Root), docs)
	return
}
