package cli

import (
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type Validate struct {
	common.ASCIIDocAttributes  `embed:""`
	spec.ParserOptions         `embed:""`
	pipeline.ProcessingOptions `embed:""`
}

func (c *Validate) Run(cc *Context) (err error) {

	var specification *spec.Specification
	specification, _, err = spec.Parse(cc, c.ParserOptions, c.ProcessingOptions, c.ASCIIDocAttributes.ToList())
	if err != nil {
		return
	}
	spec.Validate(specification)
	return
}
