package asciidoc

import (
	"fmt"
	"strings"
)

type ElementType uint8

const (
	ElementTypeUnknown ElementType = iota
	ElementTypeDocument
	ElementTypeBlock
	ElementTypeInline
	ElementTypeInlineLiteral
	ElementTypeAttribute
	ElementTypeAttributes
)

type Element interface {
	Type() ElementType
	Equals(o Element) bool
}

type HasElements interface {
	Element
	Elements() Set
	Append(e ...Element)
	SetElements(e Set)
}

type HasChild interface {
	Child() Element
}

type HasParent interface {
	Parent() Element
	SetParent(e Element)
}

func ValueToString(e any) string {
	var sb strings.Builder
	elementToString(&sb, e)
	return sb.String()
}

func elementToString(sb *strings.Builder, e any) {
	switch e := e.(type) {
	case []any:
		for _, ee := range e {
			elementToString(sb, ee)
		}
	case Set:
		for _, ee := range e {
			elementToString(sb, ee)
		}
	case []byte:
		sb.WriteString(string(e))
	case string:
		sb.WriteString(e)
	case *String:
		sb.WriteString(e.Value)
	default:
		panic(fmt.Errorf("unexpected element type: %T", e))
	}
}
