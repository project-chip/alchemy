package cli

import (
	"bytes"
	"fmt"
	"os"

	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/provisional"
)

type Provisional struct {
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`

	spec.FilterOptions `embed:""`

	HeadRoot string `name:"head-root" default:"connectedhomeip-spec" help:"the src root of your clone of CHIP-Specifications/connectedhomeip-spec"  group:"Provisional:"`
	BaseRoot string `name:"base-root" default:"connectedhomeip-spec" help:"the src root of your clone of CHIP-Specifications/connectedhomeip-spec"  group:"Provisional:"`
}

func (c *Provisional) Run(cc *Context) (err error) {

	var out bytes.Buffer
	writer := files.NewPatcher[string]("Generating patch file...", &out)
	writer.Root = c.HeadRoot

	_, err = provisional.Pipeline(cc, c.BaseRoot, c.HeadRoot, nil, c.ProcessingOptions, writer)
	if err != nil {
		return
	}

	fmt.Fprintf(os.Stderr, "patch:\n%s", out.String())
	return
}
