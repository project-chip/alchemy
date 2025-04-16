package disco

import (
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/disco"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type Command struct {
	disco.DiscoOptions `embed:""`

	spec.ParserOptions         `embed:""`
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	render.RenderOptions       `embed:""`
	files.OutputOptions        `embed:""`

	Paths []string `arg:"" optional:"" help:"The paths of AsciiDoc files to disco-ball. If not specified, all files will be disco-balled."`
}

func (d *Command) Run(alchemy *cli.Alchemy) (err error) {
	writer := files.NewWriter[string]("Writing disco-balled docs", d.OutputOptions)

	err = disco.Pipeline(alchemy, d.SpecRoot, d.Paths, d.ProcessingOptions, d.DiscoOptions, d.RenderOptions.ToOptions(), writer)
	return
}
