package render

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/output"
)

func renderParagraph(cxt *output.Context, p *types.Paragraph, previous *interface{}) (err error) {
	err = renderAttributes(cxt, p, p.Attributes, false)
	if err != nil {
		return
	}
	err = RenderElements(cxt, "", p.Elements)
	return
}
