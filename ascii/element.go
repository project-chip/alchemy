package ascii

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

type Element struct {
	Parent interface{}
	Base   interface{}
}

func NewElement(parent interface{}, base interface{}) *Element {
	return &Element{Parent: parent, Base: base}
}

func (e *Element) GetElements() []interface{} {
	if we, ok := e.Base.(types.WithElements); ok {
		return we.GetElements()
	}
	return []interface{}{}
}

func (e *Element) SetElements(elements []interface{}) error {
	if we, ok := e.Base.(types.WithElements); ok {
		return we.SetElements(elements)
	}
	return fmt.Errorf("base element does not have elements: %T", e.Base)
}
