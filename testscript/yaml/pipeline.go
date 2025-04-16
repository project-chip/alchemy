package yaml

import (
	"context"
	"path/filepath"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/testscript"
	"github.com/project-chip/alchemy/testscript/python"
	"github.com/project-chip/alchemy/testscript/yaml/parse"
)

func Pipeline(cxt context.Context, sdkRoot string, pipelineOptions pipeline.ProcessingOptions, parserOptions []spec.ParserOption, asciiSettings []asciidoc.AttributeName, generatorOptions []python.GeneratorOption, fileOptions files.OutputOptions, filePaths []string) (err error) {

	err = sdk.CheckAlchemyVersion(sdkRoot)
	if err != nil {
		return
	}

	var inputs pipeline.Paths
	inputs, err = pipeline.Start(cxt, paths.NewTargeter(filePaths...))
	if err != nil {
		return
	}

	picsYamlFilter := func(cxt context.Context, inputs []*pipeline.Data[struct{}]) (outputs []*pipeline.Data[struct{}], err error) {
		for _, input := range inputs {
			switch filepath.Base(input.Path) {
			case "PICS.yaml":
			default:
				outputs = append(outputs, input)
			}
		}
		return
	}

	inputs, err = pipeline.Collective(cxt, pipelineOptions, pipeline.CollectiveFunc("Filtering YAML tests", picsYamlFilter), inputs)
	if err != nil {
		return
	}

	specParser, err := spec.NewParser(asciiSettings, parserOptions...)
	if err != nil {
		return
	}

	err = errata.LoadErrataConfig(specParser.Root)
	if err != nil {
		return
	}

	var parser parse.TestYamlParser
	parser, err = parse.NewTestYamlParser(sdkRoot)
	if err != nil {
		return
	}

	var tests pipeline.Map[string, *pipeline.Data[*parse.Test]]
	tests, err = pipeline.Parallel(cxt, pipelineOptions, parser, inputs)
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

	converter := NewYamlTestConverter(specBuilder.Spec, sdkRoot, picsLabels)

	testplans, err := pipeline.Parallel(cxt, pipelineOptions, converter, tests)
	if err != nil {
		return
	}

	generator := testscript.NewTestScriptConverter(specBuilder.Spec, sdkRoot, picsLabels)

	var testscripts pipeline.Map[string, *pipeline.Data[*testscript.Test]]
	testscripts, err = pipeline.Parallel(cxt, pipelineOptions, generator, testplans)
	if err != nil {
		return
	}

	renderer := python.NewPythonTestRenderer(specBuilder.Spec, sdkRoot, picsLabels, generatorOptions...)
	var scripts pipeline.StringSet
	scripts, err = pipeline.Parallel(cxt, pipelineOptions, renderer, testscripts)
	if err != nil {
		return
	}

	writer := files.NewWriter[string]("Writing test scripts", fileOptions)
	err = writer.Write(cxt, scripts, pipelineOptions)
	return
}
