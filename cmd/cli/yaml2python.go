package cli

import (
	"context"
	"path/filepath"

	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/testscript"
	"github.com/project-chip/alchemy/testscript/python"
	"github.com/project-chip/alchemy/testscript/yaml"
	"github.com/project-chip/alchemy/testscript/yaml/parse"
)

type Yaml2Python struct {
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
	spec.ParserOptions         `embed:""`
	render.RenderOptions       `embed:""`
	sdk.SDKOptions             `embed:""`
	python.GeneratorOptions    `embed:""`

	Paths []string `arg:""`
}

func (c *Yaml2Python) Run(cc *Context) (err error) {

	err = sdk.CheckAlchemyVersion(c.SDKOptions.SdkRoot)
	if err != nil {
		return
	}

	var inputs pipeline.Paths
	inputs, err = pipeline.Start(cc, paths.NewTargeter(c.Paths...))
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

	inputs, err = pipeline.Collective(cc, c.ProcessingOptions, pipeline.CollectiveFunc("Filtering YAML tests", picsYamlFilter), inputs)
	if err != nil {
		return
	}

	specParser, err := spec.NewParser(c.ASCIIDocAttributes.ToList(), c.ParserOptions)
	if err != nil {
		return
	}

	err = errata.LoadErrataConfig(c.ParserOptions.Root)
	if err != nil {
		return
	}

	var parser parse.TestYamlParser
	parser, err = parse.NewTestYamlParser(c.SDKOptions.SdkRoot)
	if err != nil {
		return
	}

	var tests pipeline.Map[string, *pipeline.Data[*parse.Test]]
	tests, err = pipeline.Parallel(cc, c.ProcessingOptions, parser, inputs)
	if err != nil {
		return
	}

	specFiles, err := pipeline.Start(cc, specParser.Targets)
	if err != nil {
		return
	}

	specDocs, err := pipeline.Parallel(cc, c.ProcessingOptions, specParser, specFiles)
	if err != nil {
		return
	}

	specBuilder := spec.NewBuilder(c.ParserOptions.Root)
	_, err = pipeline.Collective(cc, c.ProcessingOptions, &specBuilder, specDocs)
	if err != nil {
		return
	}

	err = spec.PatchSpecForSdk(specBuilder.Spec)
	if err != nil {
		return
	}

	picsLabels, err := parse.LoadPICSLabels(c.SDKOptions.SdkRoot)
	if err != nil {
		return
	}

	converter := yaml.NewYamlTestConverter(specBuilder.Spec, c.SDKOptions.SdkRoot, picsLabels)

	testplans, err := pipeline.Parallel(cc, c.ProcessingOptions, converter, tests)
	if err != nil {
		return
	}

	generator := testscript.NewTestScriptConverter(specBuilder.Spec, c.SDKOptions.SdkRoot, picsLabels)

	var testscripts pipeline.Map[string, *pipeline.Data[*testscript.Test]]
	testscripts, err = pipeline.Parallel(cc, c.ProcessingOptions, generator, testplans)
	if err != nil {
		return
	}

	renderer := python.NewPythonTestRenderer(specBuilder.Spec, c.SDKOptions.SdkRoot, picsLabels, c.GeneratorOptions.ToOptions()...)
	var scripts pipeline.StringSet
	scripts, err = pipeline.Parallel(cc, c.ProcessingOptions, renderer, testscripts)
	if err != nil {
		return
	}

	writer := files.NewWriter[string]("Writing test scripts", c.OutputOptions)
	err = writer.Write(cc, scripts, c.ProcessingOptions)
	return
}
