package spec

import (
	"context"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
)

func Parse(cxt context.Context, parserOptions ParserOptions, processingOptions pipeline.ProcessingOptions, builderOptions []BuilderOption, attributes []asciidoc.AttributeName) (specification *Specification, specDocs DocSet, err error) {

	var docGroups pipeline.Map[string, *pipeline.Data[*DocGroup]]
	docGroups, specDocs, err = BuildDocumentGroups(cxt, parserOptions, processingOptions)
	if err != nil {
		return
	}

	specification, specDocs, err = Build(cxt, parserOptions, processingOptions, builderOptions, docGroups, attributes)
	return
}

func Build(cxt context.Context, parserOptions ParserOptions, processingOptions pipeline.ProcessingOptions, builderOptions []BuilderOption, docGroups DocGroupSet, attributes []asciidoc.AttributeName) (specification *Specification, specDocs DocSet, err error) {
	err = PreParse(cxt, parserOptions, processingOptions, docGroups, attributes)
	if err != nil {
		return
	}

	specBuilder := NewBuilder(parserOptions.Root, builderOptions...)
	specDocs, err = pipeline.Collective(cxt, processingOptions, &specBuilder, docGroups)
	if err != nil {
		return
	}

	specification = specBuilder.Spec
	return
}

func PreParse(cxt context.Context, parserOptions ParserOptions, processingOptions pipeline.ProcessingOptions, docGroups DocGroupSet, attributes []asciidoc.AttributeName) (err error) {
	var preparser *PreParser
	preparser, err = NewPreParser(parserOptions.Root, attributes)
	if err != nil {
		return
	}

	_, err = pipeline.Parallel(cxt, processingOptions, preparser, docGroups)
	return
}

func BuildDocumentGroups(cxt context.Context, parserOptions ParserOptions, processingOptions pipeline.ProcessingOptions) (docGroups DocGroupSet, specDocs DocSet, err error) {
	var specParser Parser
	specParser, err = NewParser(parserOptions)
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

	specDocs, err = pipeline.Parallel(cxt, processingOptions, specParser, specPaths)
	if err != nil {
		return
	}

	docGroups, err = pipeline.Collective(cxt, processingOptions, NewDocumentGrouper(parserOptions.Root), specDocs)
	return
}
