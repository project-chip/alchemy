package parse

import (
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
)

type attributeValue struct {
	quote asciidoc.AttributeQuoteType
	value asciidoc.Elements
}

func extractAttributes(els []any, out []asciidoc.Attribute) ([]asciidoc.Attribute, error) {
	var err error
	for _, e := range els {
		switch e := e.(type) {
		case asciidoc.Attribute:
			out = append(out, e)
		case []asciidoc.Attribute:
			out = append(out, e...)
		case []any:
			out, err = extractAttributes(e, out)
		default:
			err = fmt.Errorf("unexpected value looking for attributes: %T", e)
		}
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

type conditionalAttributes struct {
	names []asciidoc.AttributeName
	union asciidoc.ConditionalUnion
}
