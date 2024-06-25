package parse

import (
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
)

type attributeValue struct {
	quote asciidoc.AttributeQuoteType
	value asciidoc.Set
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

type AttributeContext struct {
	attributes map[string]any
	counters   map[string]*counter
}

type counterType uint8

const (
	counterTypeInteger counterType = iota
	counterTypeUpperCase
	counterTypeLowerCase
)

type counter struct {
	counterType counterType
	value       int
}

func (ac *AttributeContext) IsSet(name string) bool {
	_, ok := ac.attributes[name]
	return ok
}

func (ac *AttributeContext) Get(name string) any {
	return ac.attributes[name]
}

func (ac *AttributeContext) Set(name string, value any) {
	if ac.attributes == nil {
		ac.attributes = make(map[string]any)
	}
	ac.attributes[name] = value
}

func (ac *AttributeContext) Unset(name string) {
	delete(ac.attributes, name)
}
