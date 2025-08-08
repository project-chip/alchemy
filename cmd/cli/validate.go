package cli

import (
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type Validate struct {
	common.ASCIIDocAttributes  `embed:""`
	spec.ParserOptions         `embed:""`
	pipeline.ProcessingOptions `embed:""`
}

func (c *Validate) Run(cc *Context) (err error) {

	specParser, err := spec.NewParser(c.ASCIIDocAttributes.ToList(), c.ParserOptions)
	if err != nil {
		return err
	}

	err = errata.LoadErrataConfig(c.ParserOptions.Root)
	if err != nil {
		return
	}

	specPaths, err := pipeline.Start(cc, specParser.Targets)
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Parallel(cc, c.ProcessingOptions, specParser, specPaths)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder(c.ParserOptions.Root)
	_, err = pipeline.Collective(cc, c.ProcessingOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	spec.Validate(specBuilder.Spec)
	return
}
