package cli

import (
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/testscript"
	"github.com/project-chip/alchemy/testscript/python"
	"github.com/project-chip/alchemy/testscript/yaml/parse"
)

type TestScript struct {
	sdk.SDKOptions             `embed:""`
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	spec.ParserOptions         `embed:""`
	spec.FilterOptions         `embed:""`
	python.GeneratorOptions    `embed:""`
	files.OutputOptions        `embed:""`
}

func (cmd *TestScript) Run(cc *Context) (err error) {

	err = sdk.CheckAlchemyVersion(cmd.SDKOptions.SdkRoot)
	if err != nil {
		return
	}

	err = errata.LoadErrataConfig(cmd.ParserOptions.Root)
	if err != nil {
		return
	}

	specParser, err := spec.NewParser(cmd.ASCIIDocAttributes.ToList(), cmd.ParserOptions)
	if err != nil {
		return
	}

	specPaths, err := pipeline.Start(cc, specParser.Targets)
	if err != nil {
		return
	}

	specDocs, err := pipeline.Parallel(cc, cmd.ProcessingOptions, specParser, specPaths)
	if err != nil {
		return
	}

	specBuilder := spec.NewBuilder(cmd.ParserOptions.Root)
	specDocs, err = pipeline.Collective(cc, cmd.ProcessingOptions, &specBuilder, specDocs)
	if err != nil {
		return
	}

	err = spec.PatchSpecForSdk(specBuilder.Spec)
	if err != nil {
		return
	}

	picsLabels, err := parse.LoadPICSLabels(cmd.SDKOptions.SdkRoot)
	if err != nil {
		return
	}

	specDocs, err = filterSpecDocs(cc, specDocs, specBuilder.Spec, cmd.FilterOptions, cmd.ProcessingOptions)
	if err != nil {
		return
	}

	specDocs, err = filterSpecDocs(cc, specDocs, specBuilder.Spec, cmd.FilterOptions, cmd.ProcessingOptions)
	if err != nil {
		return
	}

	err = checkSpecErrors(cc, specBuilder.Spec, cmd.FilterOptions, specDocs)
	if err != nil {
		return
	}

	scriptGenerator := testscript.NewTestScriptGenerator(specBuilder.Spec, cmd.SDKOptions.SdkRoot, picsLabels)
	testplans, err := pipeline.Parallel(cc, cmd.ProcessingOptions, scriptGenerator, specDocs)
	if err != nil {
		return
	}

	generator := python.NewPythonTestRenderer(specBuilder.Spec, cmd.SDKOptions.SdkRoot, picsLabels, cmd.GeneratorOptions.ToOptions()...)
	var scripts pipeline.StringSet
	scripts, err = pipeline.Parallel(cc, cmd.ProcessingOptions, generator, testplans)
	if err != nil {
		return
	}

	writer := files.NewWriter[string]("Writing test scripts", cmd.OutputOptions)
	err = writer.Write(cc, scripts, cmd.ProcessingOptions)
	return
}
