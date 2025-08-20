package asciidoc

import (
	"fmt"
	"iter"
	"strings"
)

type ElementIterator iter.Seq[Element]

func (ei ElementIterator) List() (elements Elements) {
	ei(func(e Element) bool {
		elements = append(elements, e)
		return true
	})
	return
}

func (ei ElementIterator) Count() (count int) {
	ei(func(e Element) bool {
		count++
		return true
	})
	return
}

type Reader interface {
	Iterate(parent Parent, elements Elements) ElementIterator
	StringValue(parent Parent, elements Elements) (string, error)
	Parent(child ChildElement) Element
	Children(parent ParentElement) Elements
}

type rawReader struct {
}

func (rr *rawReader) Iterate(parent Parent, elements Elements) ElementIterator {
	return func(yield func(Element) bool) {
		for _, e := range elements {
			if !yield(e) {
				return
			}
		}
	}
}

func (rr *rawReader) StringValue(parent Parent, elements Elements) (string, error) {
	var sb strings.Builder
	for el := range rr.Iterate(parent, elements) {
		switch el := el.(type) {
		case *String:
			sb.WriteString(el.Value)
		case *NewLine:
			sb.WriteRune('\n')
		case *EmptyLine:
			sb.WriteRune('\n')
		case *CharacterReplacementReference:
			sb.WriteString(el.ReplacementValue())
		default:
			return "", fmt.Errorf("unexpected type rendering preparsed doc: %T", el)
		}
	}
	return sb.String(), nil
}

func (rr *rawReader) Parent(child ChildElement) Element {
	return child.Parent()
}

func (rr *rawReader) Children(parent ParentElement) Elements {
	return parent.Children()
}

var RawReader Reader = &rawReader{}

type Traverser interface {
	Traverse(parent ParentElement) iter.Seq2[ParentElement, Parent]
}
