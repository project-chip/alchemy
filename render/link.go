package render

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderInlineLink(cxt *output.Context, il *types.InlineLink) {
	if il.Location == nil {
		return
	}
	if len(il.Location.Scheme) > 0 {
		cxt.WriteString(il.Location.Scheme)
	} else {
		cxt.WriteString("link:")
	}
	cxt.WriteString(getPath(il.Location))
	renderAttributes(cxt, il, il.Attributes)
}

func renderImageBlock(cxt *output.Context, ib *types.ImageBlock) {
	cxt.WriteNewline()
	renderSelectAttributes(cxt, ib, ib.Attributes, AttributeTypeID|AttributeTypeTitle, AttributeTypeNone)
	cxt.WriteString("image::")
	cxt.WriteString(ib.Location.Scheme)
	cxt.WriteString(getPath(ib.Location))
	renderSelectAttributes(cxt, ib, ib.Attributes, AttributeTypeAll, AttributeTypeID|AttributeTypeTitle|AttributeTypeCols)
	cxt.WriteNewline()
}

func renderInlineImage(cxt *output.Context, ib *types.InlineImage) {
	cxt.WriteNewline()
	renderSelectAttributes(cxt, ib, ib.Attributes, AttributeTypeID|AttributeTypeTitle, AttributeTypeNone)
	cxt.WriteString("image:")
	cxt.WriteString(ib.Location.Scheme)
	cxt.WriteString(getPath(ib.Location))
	renderSelectAttributes(cxt, ib, ib.Attributes, AttributeTypeAll, AttributeTypeID|AttributeTypeTitle|AttributeTypeCols)
	cxt.WriteNewline()
}

func getPath(l *types.Location) string {
	var out strings.Builder
	switch p := l.Path.(type) {
	case string:
		out.WriteString(p)
	default:
		panic(fmt.Errorf("unknown image location path type: %T", p))
	}
	return out.String()
}
