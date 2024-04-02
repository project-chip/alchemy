package render

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func renderParagraph(cxt *Context, p *types.Paragraph, previous *any) (err error) {
	err = renderAttributes(cxt, p, p.Attributes, false)
	if err != nil {
		return
	}
	err = Elements(cxt, "", p.Elements)
	return
}
