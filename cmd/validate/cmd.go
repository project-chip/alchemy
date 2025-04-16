package validate

import (
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/spec/validate"
)

type Command struct {
	common.ASCIIDocAttributes  `embed:""`
	spec.ParserOptions         `embed:""`
	pipeline.ProcessingOptions `embed:""`
}

func (c *Command) Run(alchemy *cli.Alchemy) (err error) {

	specParser, err := spec.NewParser(c.ASCIIDocAttributes.ToList(), c.ParserOptions.ToOptions()...)
	if err != nil {
		return err
	}

	err = errata.LoadErrataConfig(specParser.Root)
	if err != nil {
		return
	}

	specFiles, err := pipeline.Start(alchemy, specParser.Targets)
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Parallel(alchemy, c.ProcessingOptions, specParser, specFiles)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder(specParser.Root)
	_, err = pipeline.Collective(alchemy, c.ProcessingOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	validate.Validate(specBuilder.Spec)
	return
}
