package testscript

import (
	"context"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/testscript/python"
	"github.com/project-chip/alchemy/yaml2python/parse"
)

func Pipeline(cxt context.Context, specRoot string, sdkRoot string, pipelineOptions pipeline.Options, asciiSettings []asciidoc.AttributeName, generatorOptions []python.GeneratorOption, fileOptions files.Options, filePaths []string) (err error) {

	err = sdk.CheckAlchemyVersion(sdkRoot)
	if err != nil {
		return
	}

	errata.LoadErrataConfig(specRoot)

	docParser, err := spec.NewParser(specRoot, asciiSettings)
	if err != nil {
		return
	}

	specFiles, err := pipeline.Start(cxt, spec.Targeter(specRoot))
	if err != nil {
		return
	}

	specDocs, err := pipeline.Parallel(cxt, pipelineOptions, docParser, specFiles)
	if err != nil {
		return
	}

	specBuilder := spec.NewBuilder()
	_, err = pipeline.Collective(cxt, pipelineOptions, &specBuilder, specDocs)
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
		filter := files.NewPathFilter[*spec.Doc](filePaths)
		specDocs, err = pipeline.Collective(cxt, pipelineOptions, filter, specDocs)
		if err != nil {
			return
		}
	}

	scriptGenerator := NewTestPlanGenerator(specBuilder.Spec, sdkRoot, picsLabels)
	testplans, err := pipeline.Parallel(cxt, pipelineOptions, scriptGenerator, specDocs)
	if err != nil {
		return
	}

	generator := python.NewPythonTestGenerator(specBuilder.Spec, sdkRoot, picsLabels, generatorOptions...)
	var scripts pipeline.StringSet
	scripts, err = pipeline.Parallel(cxt, pipelineOptions, generator, testplans)
	if err != nil {
		return
	}

	writer := files.NewWriter[string]("Writing test scripts", fileOptions)
	err = writer.Write(cxt, scripts, pipelineOptions)
	return
}
