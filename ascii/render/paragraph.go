package render

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func renderParagraph(cxt *Context, p *types.Paragraph, previous *interface{}) (err error) {
	err = renderAttributes(cxt, p, p.Attributes, false)
	if err != nil {
		return
	}
	err = RenderElements(cxt, "", p.Elements)
	return
}
