package ascii

import (
	"fmt"
)

type Element struct {
	Parent any
	Base   any
}

func NewElement(parent any, base any) *Element {
	return &Element{Parent: parent, Base: base}
}

func (e *Element) GetElements() []any {
	if we, ok := e.Base.(elements.WithElements); ok {
		return we.GetElements()
	}
	return []any{}
}

func (e *Element) SetElements(elements []any) error {
	if we, ok := e.Base.(elements.WithElements); ok {
		return we.SetElements(elements)
	}
	return fmt.Errorf("base element does not have elements: %T", e.Base)
}

func (e *Element) GetBase() any {
	return e.Base
}
