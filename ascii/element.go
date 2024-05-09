package ascii

import (
	"fmt"

	"github.com/hasty/adoc/elements"
)

type Element struct {
	Parent any
	Base   elements.Element
}

func NewElement(parent any, base elements.Element) *Element {
	return &Element{Parent: parent, Base: base}
}

func (e *Element) GetElements() elements.Set {
	if we, ok := e.Base.(elements.HasElements); ok {
		return we.Elements()
	}
	return elements.Set{}
}

func (e *Element) SetElements(els elements.Set) error {
	if we, ok := e.Base.(elements.HasElements); ok {
		return we.SetElements(els)
	}
	return fmt.Errorf("base element does not have elements: %T", e.Base)
}

func (e *Element) GetBase() elements.Element {
	return e.Base
}

func (e *Element) Type() elements.ElementType {
	return e.Base.Type()
}

func (e *Element) Equals(o elements.Element) bool {
	return e.Base.Equals(o)
}
