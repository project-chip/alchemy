package cli

import (
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/disco"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type Disco struct {
	disco.DiscoOptions `embed:""`

	spec.ParserOptions         `embed:""`
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	render.RenderOptions       `embed:""`
	files.OutputOptions        `embed:""`

	Paths []string `arg:"" optional:"" help:"The paths of AsciiDoc files to disco-ball. If not specified, all files will be disco-balled."`
}

func (d *Disco) Run(cc *Context) (err error) {
	writer := files.NewWriter[string]("Writing disco-balled docs", d.OutputOptions)

	err = disco.Pipeline(cc, d.ParserOptions, d.Paths, d.ProcessingOptions, d.DiscoOptions, d.ASCIIDocAttributes.ToList(), d.RenderOptions.ToOptions(), writer)
	return
}
