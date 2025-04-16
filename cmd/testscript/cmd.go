package testscript

import (
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/testscript/python"
)

type Command struct {
	sdk.SDKOptions             `embed:""`
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	spec.ParserOptions         `embed:""`
	python.GeneratorOptions    `embed:""`
	files.OutputOptions        `embed:""`

	Paths []string `arg:""`
}

func (cmd *Command) Run(alchemy *cli.Alchemy) (err error) {

	err = python.Pipeline(alchemy,
		cmd.SdkRoot,
		cmd.ProcessingOptions,
		cmd.ParserOptions.ToOptions(),
		cmd.ASCIIDocAttributes.ToList(),
		cmd.GeneratorOptions.ToOptions(),
		cmd.OutputOptions,
		cmd.Paths)

	return
}
