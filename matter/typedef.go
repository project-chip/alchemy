package matter

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/types"
)

type TypeDef struct {
	entity
	ParentEntity types.Entity `json:"-"`

	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Type        *types.DataType `json:"type,omitempty"`
}

func NewTypeDef(source asciidoc.Element) *TypeDef {
	return &TypeDef{
		entity: entity{source: source},
	}
}

func (*TypeDef) EntityType() types.EntityType {
	return types.EntityTypeDef
}

func (s *TypeDef) Clone() *TypeDef {
	ns := &TypeDef{Name: s.Name, Description: s.Description, Type: s.Type}
	return ns
}

func (s *TypeDef) Inherit(parent *TypeDef) {
	if len(s.Description) == 0 {
		s.Description = parent.Description
	}
	s.Type = s.Type
}

type TypeDefSet []*TypeDef

func (ss TypeDefSet) Identifier(name string) (types.Entity, bool) {
	for _, e := range ss {
		if e.Name == name {
			return e, true
		}
	}
	return nil, false
}
