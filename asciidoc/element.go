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
	Append(e Element) error
	SetElements(e Set) error
}

type HasChild interface {
	Child() Element
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

func MergeStrings(els []any) (out Set) {
	var s strings.Builder
	for _, e := range els {
		switch e := e.(type) {
		case string:
			s.WriteString(e)
		case Element:
			if s.Len() > 0 {
				out = append(out, NewString(s.String()))
				s.Reset()
			}
			out = append(out, e)
		default:
			fmt.Printf("unexpected type in string merge: %T\n", e)
		}
	}
	if s.Len() > 0 {
		out = append(out, NewString(s.String()))
	}
	return
}
