package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderInlineLink(cxt *output.Context, il *types.InlineLink) (err error) {
	if il.Location == nil {
		return nil
	}
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
	return renderAttributes(cxt, il, il.Attributes)
}

func renderImageBlock(cxt *output.Context, ib *types.ImageBlock) (err error) {
	cxt.WriteNewline()
	err = renderSelectAttributes(cxt, ib, ib.Attributes, AttributeTypeID|AttributeTypeTitle, AttributeTypeNone)
	if err != nil {
		return
	}
	cxt.WriteString("image::")
	cxt.WriteString(ib.Location.Scheme)
	var path string
	path, err = getPath(ib.Location)
	if err != nil {
		return
	}
	cxt.WriteString(path)
	err = renderSelectAttributes(cxt, ib, ib.Attributes, AttributeTypeAll, AttributeTypeID|AttributeTypeTitle|AttributeTypeCols)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	return
}

func renderInlineImage(cxt *output.Context, ib *types.InlineImage) (err error) {
	cxt.WriteNewline()
	err = renderSelectAttributes(cxt, ib, ib.Attributes, AttributeTypeID|AttributeTypeTitle, AttributeTypeNone)
	if err != nil {
		return
	}
	cxt.WriteString("image:")
	cxt.WriteString(ib.Location.Scheme)
	var path string
	path, err = getPath(ib.Location)
	if err != nil {
		return
	}
	cxt.WriteString(path)
	err = renderSelectAttributes(cxt, ib, ib.Attributes, AttributeTypeAll, AttributeTypeID|AttributeTypeTitle|AttributeTypeCols)
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
