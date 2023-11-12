package render

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/output"
)

func renderPreamble(cxt *output.Context, p *types.Preamble) error {
	return RenderElements(cxt, "", p.Elements)
}
