package python

import (
	"context"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/testscript"
	"github.com/project-chip/alchemy/testscript/yaml/parse"
)

func Pipeline(cxt context.Context, sdkRoot string, pipelineOptions pipeline.Options, parserOptions []spec.ParserOption, asciiSettings []asciidoc.AttributeName, generatorOptions []GeneratorOption, fileOptions files.Options, filePaths []string) (err error) {

	specParser, err := spec.NewParser(asciiSettings, parserOptions...)
	if err != nil {
		return
	}

	err = sdk.CheckAlchemyVersion(sdkRoot)
	if err != nil {
		return
	}

	err = errata.LoadErrataConfig(specParser.Root)
	if err != nil {
		return
	}

	specFiles, err := pipeline.Start(cxt, specParser.Targets)
	if err != nil {
		return
	}

	specDocs, err := pipeline.Parallel(cxt, pipelineOptions, specParser, specFiles)
	if err != nil {
		return
	}

	specBuilder := spec.NewBuilder(specParser.Root)
	specDocs, err = pipeline.Collective(cxt, pipelineOptions, &specBuilder, specDocs)
	if err != nil {
		return
	}

	err = spec.PatchSpecForSdk(specBuilder.Spec)
	if err != nil {
		return
	}

	picsLabels, err := parse.LoadPICSLabels(sdkRoot)
	if err != nil {
		return
	}

	if len(filePaths) > 0 { // Filter the spec by whatever extra args were passed
		filter := paths.NewFilter[*spec.Doc](specParser.Root, filePaths)
		specDocs, err = pipeline.Collective(cxt, pipelineOptions, filter, specDocs)
		if err != nil {
			return
		}
	}

	scriptGenerator := testscript.NewTestScriptGenerator(specBuilder.Spec, sdkRoot, picsLabels)
	testplans, err := pipeline.Parallel(cxt, pipelineOptions, scriptGenerator, specDocs)
	if err != nil {
		return
	}

	generator := NewPythonTestRenderer(specBuilder.Spec, sdkRoot, picsLabels, generatorOptions...)
	var scripts pipeline.StringSet
	scripts, err = pipeline.Parallel(cxt, pipelineOptions, generator, testplans)
	if err != nil {
		return
	}

	writer := files.NewWriter[string]("Writing test scripts", fileOptions)
	err = writer.Write(cxt, scripts, pipelineOptions)
	return
}
