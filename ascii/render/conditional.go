package render

import "github.com/hasty/adoc/elements"

func renderConditional(cxt *Context, prefix string, attributes elements.AttributeNames, union elements.ConditionalUnion) {
	cxt.WriteString(prefix)
	var separator rune
	switch union {
	case elements.ConditionalUnionAll:
		separator = '+'
	case elements.ConditionalUnionAny:
		separator = ','
	}
	for i, a := range attributes {
		if i > 0 {
			cxt.WriteRune(separator)
		}
		cxt.WriteString(string(a))
	}
	cxt.WriteString("[]\n")
}

func renderIfEval(cxt *Context, el *elements.IfEval) {
	cxt.WriteString("ifeval::[")
	switch el.Left.Quote {
	case elements.AttributeQuoteTypeDouble:
		cxt.WriteRune('"')
	case elements.AttributeQuoteTypeSingle:
		cxt.WriteRune('\'')
	}
	Elements(cxt, "", el.Left.Value...)
	switch el.Left.Quote {
	case elements.AttributeQuoteTypeDouble:
		cxt.WriteRune('"')
	case elements.AttributeQuoteTypeSingle:
		cxt.WriteRune('\'')
	}
	cxt.WriteRune(' ')
	cxt.WriteString(el.Operator.String())
	cxt.WriteRune(' ')
	switch el.Right.Quote {
	case elements.AttributeQuoteTypeDouble:
		cxt.WriteRune('"')
	case elements.AttributeQuoteTypeSingle:
		cxt.WriteRune('\'')
	}
	Elements(cxt, "", el.Right.Value...)
	switch el.Right.Quote {
	case elements.AttributeQuoteTypeDouble:
		cxt.WriteRune('"')
	case elements.AttributeQuoteTypeSingle:
		cxt.WriteRune('\'')
	}

	cxt.WriteString("]\n")
}
