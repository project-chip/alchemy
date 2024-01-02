package render

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func renderPreamble(cxt *Context, p *types.Preamble) error {
	return RenderElements(cxt, "", p.Elements)
}
