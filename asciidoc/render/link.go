package render

import (
	"fmt"

	"github.com/hasty/alchemy/asciidoc"
)

func renderLink(cxt *Context, il *asciidoc.Link) (err error) {
	if len(il.URL.Scheme) > 0 {
		cxt.WriteString(il.URL.Scheme)
	} else {
		cxt.WriteString("link:")
	}
	Elements(cxt, "", il.URL.Path.(asciidoc.Set)...)

	return renderAttributes(cxt, il, il.Attributes(), true)
}

func renderImageBlock(cxt *Context, ib *asciidoc.BlockImage) (err error) {
	cxt.EnsureNewLine()
	err = renderSelectAttributes(cxt, ib, ib.Attributes(), AttributeFilterID|AttributeFilterTitle, AttributeFilterNone, false)
	if err != nil {
		return
	}
	cxt.EnsureNewLine()
	cxt.WriteString("image::")
	Elements(cxt, "", ib.Path...)
	err = renderSelectAttributes(cxt, ib, ib.Attributes(), AttributeFilterAll, AttributeFilterID|AttributeFilterTitle|AttributeFilterCols, true)
	if err != nil {
		return
	}
	cxt.EnsureNewLine()
	return
}

func renderInlineImage(cxt *Context, ib *asciidoc.InlineImage) (err error) {
	cxt.EnsureNewLine()
	err = renderSelectAttributes(cxt, ib, ib.Attributes(), AttributeFilterID|AttributeFilterTitle, AttributeFilterNone, false)
	if err != nil {
		return
	}
	cxt.EnsureNewLine()
	cxt.WriteString("image:")
	Elements(cxt, "", ib.Path...)
	err = renderSelectAttributes(cxt, ib, ib.Attributes(), AttributeFilterAll, AttributeFilterID|AttributeFilterTitle|AttributeFilterCols, true)
	if err != nil {
		return
	}
	return
}

func getPath(l asciidoc.URL) (string, error) {
	switch p := l.Path.(type) {
	case string:
		return p, nil
	default:
		return "", fmt.Errorf("unknown image location path type: %T", p)
	}
}
