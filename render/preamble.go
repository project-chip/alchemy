package render

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderPreamble(cxt *output.Context, p *types.Preamble) {
	RenderElements(cxt, "", p.Elements)
}
