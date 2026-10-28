package matter

import (
	"iter"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/types"
)

type Constant struct {
	entity

	Name  string `json:"name,omitempty"`
	Value any    `json:"value,omitempty"`
}

func NewConstant(source asciidoc.Element) *Constant {
	return &Constant{
		entity: entity{source: source},
	}
}

func (*Constant) EntityType() types.EntityType {
	return types.EntityTypeConstant
}

func (c *Constant) Equals(e types.Entity) bool {
	oc, ok := e.(*Constant)
	if !ok {
		return false
	}
	if !c.sameParent(e) {
		return false
	}
	return c.Name == oc.Name
}

func (s *Constant) CloneTo(newParent types.Entity) *Constant {
	return &Constant{entity: entity{source: s.source, parent: newParent}, Name: s.Name, Value: s.Value}
}

func (s *Constant) Inherit(parent *Constant) {

}

type ConstantSet []*Constant

func (cs ConstantSet) Identifier(name string) (types.Entity, bool) {
	for _, e := range cs {
		if e.Name == name {
			return e, true
		}
	}
	return nil, false
}

func (cs ConstantSet) Iterate() iter.Seq[types.Entity] {
	return func(yield func(types.Entity) bool) {
		for _, c := range cs {
			if !yield(c) {
				return
			}
		}
	}
}
