package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func renderInlineLink(cxt *Context, il *types.InlineLink) (err error) {
	if il.Location != nil {
		if len(il.Location.Scheme) > 0 {
			cxt.WriteString(il.Location.Scheme)
		} else {
			cxt.WriteString("link:")
		}
		var path string
		path, err = getPath(il.Location)
		if err != nil {
			return
		}
		cxt.WriteString(path)
	}
	return renderAttributes(cxt, il, il.Attributes, true)
}

func renderImageBlock(cxt *Context, ib *types.ImageBlock) (err error) {
	cxt.WriteNewline()
	err = renderSelectAttributes(cxt, ib, ib.Attributes, AttributeFilterID|AttributeFilterTitle, AttributeFilterNone, false)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	cxt.WriteString("image::")
	cxt.WriteString(ib.Location.Scheme)
	var path string
	path, err = getPath(ib.Location)
	if err != nil {
		return
	}
	cxt.WriteString(path)
	err = renderSelectAttributes(cxt, ib, ib.Attributes, AttributeFilterAll, AttributeFilterID|AttributeFilterTitle|AttributeFilterCols, true)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	return
}

func renderInlineImage(cxt *Context, ib *types.InlineImage) (err error) {
	cxt.WriteNewline()
	err = renderSelectAttributes(cxt, ib, ib.Attributes, AttributeFilterID|AttributeFilterTitle, AttributeFilterNone, true)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	cxt.WriteString("image:")
	cxt.WriteString(ib.Location.Scheme)
	var path string
	path, err = getPath(ib.Location)
	if err != nil {
		return
	}
	cxt.WriteString(path)
	err = renderSelectAttributes(cxt, ib, ib.Attributes, AttributeFilterAll, AttributeFilterID|AttributeFilterTitle|AttributeFilterCols, true)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	return
}

func getPath(l *types.Location) (string, error) {
	switch p := l.Path.(type) {
	case string:
		return p, nil
	default:
		return "", fmt.Errorf("unknown image location path type: %T", p)
	}
}
