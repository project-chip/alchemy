package cli

import (
	"github.com/project-chip/alchemy/errdiff"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type ErrDiff struct {
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`

	spec.FilterOptions `embed:""`

	HeadRoot   string `name:"head-root" default:"connectedhomeip-spec" help:"the src root of your clone of CHIP-Specifications/connectedhomeip-spec"`
	BaseRoot   string `name:"base-root" default:"connectedhomeip-spec" help:"the src root of your clone of CHIP-Specifications/connectedhomeip-spec"`
}

func (c *ErrDiff) Run(cc *Context) (err error) {
	_, err = errdiff.Pipeline(cc, c.BaseRoot, c.HeadRoot, c.Paths, c.ProcessingOptions)

	return
}
