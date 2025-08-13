package spec

import (
	"context"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
)

func Read(cxt context.Context, processingOptions pipeline.ProcessingOptions, builderOptions []BuilderOption, specRoot string, docPaths []string) (specification *Specification, specDocs DocSet, err error) {

	specTargeter := Targeter(specRoot)

	var inputs pipeline.Paths
	inputs, err = pipeline.Start(cxt, specTargeter)
	if err != nil {
		return
	}

	var specReader Reader
	specReader, err = NewReader("Reading spec docs", specRoot)
	if err != nil {
		return
	}

	specDocs, err = pipeline.Parallel(cxt, processingOptions, specReader, inputs)
	if err != nil {
		return
	}

	docGroups, err := pipeline.Collective(cxt, processingOptions, NewDocumentGrouper(specRoot), specDocs)
	if err != nil {
		return
	}

	specBuilder := NewBuilder(specReader.Root, builderOptions...)
	_, err = pipeline.Collective(cxt, processingOptions, &specBuilder, docGroups)
	if err != nil {
		return
	}

	if len(docPaths) > 0 {
		filter := paths.NewIncludeFilter[*Doc](specRoot, docPaths)
		specDocs, err = pipeline.Collective(cxt, processingOptions, filter, specDocs)
		if err != nil {
			return
		}
	}

	specification = specBuilder.Spec
	return
}

func Parse(cxt context.Context, parserOptions ParserOptions, processingOptions pipeline.ProcessingOptions, builderOptions []BuilderOption, attributes []asciidoc.AttributeName) (specification *Specification, specDocs DocSet, err error) {

	var specParser Parser
	specParser, err = NewParser(parserOptions)
	if err != nil {
		return
	}

	err = errata.LoadErrataConfig(parserOptions.Root)
	if err != nil {
		return
	}

	var specPaths pipeline.Paths
	specPaths, err = pipeline.Start(cxt, specParser.Targets)
	if err != nil {
		return
	}

	specDocs, err = pipeline.Parallel(cxt, processingOptions, specParser, specPaths)
	if err != nil {
		return
	}

	var docGroups pipeline.Map[string, *pipeline.Data[*DocGroup]]
	docGroups, err = pipeline.Collective(cxt, processingOptions, NewDocumentGrouper(parserOptions.Root), specDocs)
	if err != nil {
		return
	}

	var preparser *PreParser
	preparser, err = NewPreParser(parserOptions.Root, attributes)
	if err != nil {
		return
	}

	_, err = pipeline.Parallel(cxt, processingOptions, preparser, docGroups)
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
