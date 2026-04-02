package render

import (
	"github.com/project-chip/alchemy/asciidoc"
)

func renderLink(cxt Target, il *asciidoc.Link) (err error) {
	cxt.StartBlock()
	cxt.WriteString(il.URL.Scheme)
	err = Elements(cxt, "", il.URL.Path...)
	if err != nil {
		return
	}

	err = renderAttributes(cxt, il.Attributes(), attributeRenderTypeInline)
	cxt.EndBlock()
	return
}

func renderLinkMacro(cxt Target, il *asciidoc.LinkMacro) (err error) {
	cxt.StartBlock()
	cxt.WriteString("link:")
	cxt.WriteString(il.URL.Scheme)
	err = Elements(cxt, "", il.URL.Path...)
	if err != nil {
		return
	}

	err = renderAttributes(cxt, il.Attributes(), attributeRenderTypeInline)
	cxt.EndBlock()
	return
}

func renderImageBlock(cxt Target, ib *asciidoc.BlockImage) (err error) {
	cxt.FlushWrap()
	cxt.EnsureNewLine()
	_, err = renderSelectAttributes(cxt, ib.Attributes(), AttributeFilterID|AttributeFilterTitle, AttributeFilterNone, attributeRenderTypeBlock)
	if err != nil {
		return
	}
	cxt.EnsureNewLine()
	cxt.DisableWrap()
	cxt.WriteString("image::")
	err = Elements(cxt, "", ib.ImagePath...)
	if err != nil {
		return
	}
	var count int
	count, err = renderSelectAttributes(cxt, ib.Attributes(), AttributeFilterAll, AttributeFilterID|AttributeFilterTitle|AttributeFilterCols, attributeRenderTypeInline)
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
	_, err = renderSelectAttributes(cxt, ib.Attributes(), AttributeFilterID|AttributeFilterTitle, AttributeFilterNone, attributeRenderTypeBlock)
	if err != nil {
		return
	}
	cxt.EnsureNewLine()
	cxt.DisableWrap()
	cxt.WriteString("image:")
	err = Elements(cxt, "", ib.ImagePath...)
	if err != nil {
		return
	}
	var count int
	count, err = renderSelectAttributes(cxt, ib.Attributes(), AttributeFilterAll, AttributeFilterID|AttributeFilterTitle|AttributeFilterCols, attributeRenderTypeInline)
	if err != nil {
		return
	}
	if count == 0 {
		cxt.WriteString("[]")
	}
	cxt.EnableWrap()
	return
}
