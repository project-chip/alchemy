package render

import (
	"github.com/project-chip/alchemy/asciidoc"
)

func renderLink(cxt *Context, il *asciidoc.Link) (err error) {
	if len(il.URL.Scheme) > 0 {
		cxt.WriteString(il.URL.Scheme)
	} else {
		cxt.WriteString("link:")
	}
	Elements(cxt, "", il.URL.Path...)

	return renderAttributes(cxt, il.Attributes(), true)
}

func renderImageBlock(cxt *Context, ib *asciidoc.BlockImage) (err error) {
	cxt.EnsureNewLine()
	_, err = renderSelectAttributes(cxt, ib.Attributes(), AttributeFilterID|AttributeFilterTitle, AttributeFilterNone, false)
	if err != nil {
		return
	}
	cxt.EnsureNewLine()
	cxt.WriteString("image::")
	Elements(cxt, "", ib.Path...)
	var count int
	count, err = renderSelectAttributes(cxt, ib.Attributes(), AttributeFilterAll, AttributeFilterID|AttributeFilterTitle|AttributeFilterCols, true)
	if err != nil {
		return
	}
	if count == 0 {
		cxt.WriteString("[]")
	}
	cxt.EnsureNewLine()
	return
}

func renderInlineImage(cxt *Context, ib *asciidoc.InlineImage) (err error) {
	cxt.EnsureNewLine()
	_, err = renderSelectAttributes(cxt, ib.Attributes(), AttributeFilterID|AttributeFilterTitle, AttributeFilterNone, false)
	if err != nil {
		return
	}
	cxt.EnsureNewLine()
	cxt.WriteString("image:")
	Elements(cxt, "", ib.Path...)
	var count int
	count, err = renderSelectAttributes(cxt, ib.Attributes(), AttributeFilterAll, AttributeFilterID|AttributeFilterTitle|AttributeFilterCols, true)
	if err != nil {
		return
	}
	if count == 0 {
		cxt.WriteString("[]")
	}
	return
}
