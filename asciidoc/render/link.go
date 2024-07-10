package render

import (
	"github.com/project-chip/alchemy/asciidoc"
)

func renderLink(cxt Target, il *asciidoc.Link) (err error) {
	cxt.StartBlock()
	cxt.WriteString(il.URL.Scheme)
	Elements(cxt, "", il.URL.Path...)

	err = renderAttributes(cxt, il.Attributes(), true)
	cxt.EndBlock()
	return
}

func renderLinkMacro(cxt Target, il *asciidoc.LinkMacro) (err error) {
	cxt.StartBlock()
	cxt.WriteString("link:")
	cxt.WriteString(il.URL.Scheme)
	Elements(cxt, "", il.URL.Path...)

	err = renderAttributes(cxt, il.Attributes(), true)
	cxt.EndBlock()
	return
}

func renderImageBlock(cxt Target, ib *asciidoc.BlockImage) (err error) {
	cxt.FlushWrap()
	cxt.EnsureNewLine()
	_, err = renderSelectAttributes(cxt, ib.Attributes(), AttributeFilterID|AttributeFilterTitle, AttributeFilterNone, false)
	if err != nil {
		return
	}
	cxt.EnsureNewLine()
	cxt.DisableWrap()
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
	cxt.EnableWrap()
	return
}

func renderInlineImage(cxt Target, ib *asciidoc.InlineImage) (err error) {
	cxt.FlushWrap()
	cxt.EnsureNewLine()
	_, err = renderSelectAttributes(cxt, ib.Attributes(), AttributeFilterID|AttributeFilterTitle, AttributeFilterNone, false)
	if err != nil {
		return
	}
	cxt.EnsureNewLine()
	cxt.DisableWrap()
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
	cxt.EnableWrap()
	return
}
