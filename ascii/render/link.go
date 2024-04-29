package render

import (
	"fmt"

	"github.com/hasty/adoc/elements"
)

func renderLink(cxt *Context, il *elements.Link) (err error) {
	if len(il.URL.Scheme) > 0 {
		cxt.WriteString(il.URL.Scheme)
	} else {
		cxt.WriteString("link:")
	}
	var path string
	path, err = getPath(il.URL)
	if err != nil {
		return
	}
	cxt.WriteString(path)

	return renderAttributes(cxt, il, il.Attributes(), true)
}

func renderImageBlock(cxt *Context, ib *elements.BlockImage) (err error) {
	cxt.WriteNewline()
	err = renderSelectAttributes(cxt, ib, ib.Attributes(), AttributeFilterID|AttributeFilterTitle, AttributeFilterNone, false)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	cxt.WriteString("image::")
	Elements(cxt, "", ib.Path...)
	err = renderSelectAttributes(cxt, ib, ib.Attributes(), AttributeFilterAll, AttributeFilterID|AttributeFilterTitle|AttributeFilterCols, true)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	return
}

func renderInlineImage(cxt *Context, ib *elements.InlineImage) (err error) {
	cxt.WriteNewline()
	err = renderSelectAttributes(cxt, ib, ib.Attributes(), AttributeFilterID|AttributeFilterTitle, AttributeFilterNone, true)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	cxt.WriteString("image:")
	Elements(cxt, "", ib.Path...)
	err = renderSelectAttributes(cxt, ib, ib.Attributes(), AttributeFilterAll, AttributeFilterID|AttributeFilterTitle|AttributeFilterCols, true)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	return
}

func getPath(l elements.URL) (string, error) {
	switch p := l.Path.(type) {
	case string:
		return p, nil
	default:
		return "", fmt.Errorf("unknown image location path type: %T", p)
	}
}
