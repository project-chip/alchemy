package matter

import (
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
	return c.Name == oc.Name
}

func (s *Constant) Clone() *Constant {
	ns := &Constant{Name: s.Name, Value: s.Value}
	return ns
}

func (s *Constant) Inherit(parent *Constant) {

}

type ConstantSet []*Constant

func (ss ConstantSet) Identifier(name string) (types.Entity, bool) {
	for _, e := range ss {
		if e.Name == name {
			return e, true
		}
	}
	return nil, false
}
