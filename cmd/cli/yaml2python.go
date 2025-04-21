package cli

import (
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/testscript/python"
	"github.com/project-chip/alchemy/testscript/yaml"
)

type Yaml2Python struct {
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
	spec.ParserOptions         `embed:""`
	render.RenderOptions       `embed:""`
	sdk.SDKOptions             `embed:""`
	python.GeneratorOptions    `embed:""`

	Paths []string `arg:""`
}

func (c *Yaml2Python) Run(cc *Context) (err error) {

	err = yaml.Pipeline(cc,
		c.SDKOptions,
		c.ProcessingOptions,
		c.ParserOptions,
		c.ASCIIDocAttributes.ToList(),
		c.GeneratorOptions.ToOptions(),
		c.OutputOptions,
		c.Paths)

	return
}
