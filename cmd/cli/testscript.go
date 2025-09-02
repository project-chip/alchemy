package cli

import (
	"github.com/project-chip/alchemy/cmd/common"
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

	var specification *spec.Specification
	var specDocs spec.DocSet
	specification, _, err = spec.Parse(cc, cmd.ParserOptions, cmd.ProcessingOptions, nil, cmd.ASCIIDocAttributes.ToList())

	err = spec.PatchSpecForSdk(specification)

	if err != nil {
		return
	}

	picsLabels, err := parse.LoadPICSLabels(cmd.SDKOptions.SdkRoot)

	if err != nil {
		return
	}

	specDocs, err = filterSpecDocs(cc, specDocs, specification, cmd.FilterOptions, cmd.ProcessingOptions)

	if err != nil {
		return
	}

	specDocs, err = filterSpecDocs(cc, specDocs, specification, cmd.FilterOptions, cmd.ProcessingOptions)

	if err != nil {
		return
	}

	err = checkSpecErrors(cc, specification, cmd.FilterOptions, specDocs)

	if err != nil {
		return
	}

	scriptGenerator := testscript.NewTestScriptGenerator(specification, cmd.SDKOptions.SdkRoot, picsLabels)
	testplans, err := pipeline.Parallel(cc, cmd.ProcessingOptions, scriptGenerator, specDocs)

	if err != nil {
		return
	}

	generator := python.NewPythonTestRenderer(specification, cmd.SDKOptions.SdkRoot, picsLabels, cmd.GeneratorOptions.ToOptions()...)
	var scripts pipeline.StringSet
	scripts, err = pipeline.Parallel(cc, cmd.ProcessingOptions, generator, testplans)

	if err != nil {
		return
	}

	writer := files.NewWriter[string]("Writing test scripts", cmd.OutputOptions)
	err = writer.Write(cc, scripts, cmd.ProcessingOptions)
	return

}
