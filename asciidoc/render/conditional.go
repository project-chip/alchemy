package render

import "github.com/project-chip/alchemy/asciidoc"

func renderConditional(cxt Target, prefix string, attributes asciidoc.AttributeNames, union asciidoc.ConditionalUnion) {
	cxt.FlushWrap()
	cxt.DisableWrap()
	cxt.WriteString(prefix)
	var separator rune
	switch union {
	case asciidoc.ConditionalUnionAll:
		separator = '+'
	case asciidoc.ConditionalUnionAny:
		separator = ','
	}
	for i, a := range attributes {
		if i > 0 {
			cxt.WriteRune(separator)
		}
		cxt.WriteString(string(a))
	}
	cxt.WriteString("[]\n")
	cxt.EnableWrap()
}

func renderIfEval(cxt Target, el *asciidoc.IfEval) {
	cxt.FlushWrap()
	cxt.DisableWrap()
	cxt.WriteString("ifeval::[")
	switch el.Left.Quote {
	case asciidoc.AttributeQuoteTypeDouble:
		cxt.WriteRune('"')
	case asciidoc.AttributeQuoteTypeSingle:
		cxt.WriteRune('\'')
	}
	Elements(cxt, "", el.Left.Value...)
	switch el.Left.Quote {
	case asciidoc.AttributeQuoteTypeDouble:
		cxt.WriteRune('"')
	case asciidoc.AttributeQuoteTypeSingle:
		cxt.WriteRune('\'')
	}
	cxt.WriteRune(' ')
	cxt.WriteString(el.Operator.String())
	cxt.WriteRune(' ')
	switch el.Right.Quote {
	case asciidoc.AttributeQuoteTypeDouble:
		cxt.WriteRune('"')
	case asciidoc.AttributeQuoteTypeSingle:
		cxt.WriteRune('\'')
	}
	Elements(cxt, "", el.Right.Value...)
	switch el.Right.Quote {
	case asciidoc.AttributeQuoteTypeDouble:
		cxt.WriteRune('"')
	case asciidoc.AttributeQuoteTypeSingle:
		cxt.WriteRune('\'')
	}

	cxt.WriteString("]\n")
	cxt.EnableWrap()
}
