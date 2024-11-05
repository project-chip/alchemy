package spec

import (
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
)

type Element struct {
	Parent any
	Base   asciidoc.Element
}

func NewElement(parent any, base asciidoc.Element) *Element {
	return &Element{Parent: parent, Base: base}
}

func (e *Element) GetElements() asciidoc.Set {
	if we, ok := e.Base.(asciidoc.HasElements); ok {
		return we.Elements()
	}
	return asciidoc.Set{}
}

func (e *Element) SetElements(els asciidoc.Set) error {
	if we, ok := e.Base.(asciidoc.HasElements); ok {
		we.SetElements(els)
		return nil
	}
	return fmt.Errorf("base element does not have elements: %T", e.Base)
}

func (e *Element) GetBase() asciidoc.Element {
	return e.Base
}

func (e *Element) Type() asciidoc.ElementType {
	return e.Base.Type()
}

func (e *Element) Equals(o asciidoc.Element) bool {
	return e.Base.Equals(o)
}
