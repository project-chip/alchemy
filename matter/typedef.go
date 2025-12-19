package matter

import (
	"iter"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/types"
)

type TypeDef struct {
	entity

	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Type        *types.DataType `json:"type,omitempty"`
}

func NewTypeDef(source asciidoc.Element, parent types.Entity) *TypeDef {
	return &TypeDef{
		entity: entity{source: source, parent: parent},
	}
}

func (*TypeDef) EntityType() types.EntityType {
	return types.EntityTypeDef
}

func (td *TypeDef) Equals(e types.Entity) bool {
	otd, ok := e.(*TypeDef)
	if !ok {
		return false
	}
	return td.Name == otd.Name
}

func (td *TypeDef) Clone() *TypeDef {
	ns := &TypeDef{Name: td.Name, Description: td.Description, Type: td.Type}
	return ns
}

func (td *TypeDef) Inherit(parent *TypeDef) {
	if len(td.Description) == 0 {
		td.Description = parent.Description
	}
	td.Type = parent.Type
}

type TypeDefSet []*TypeDef

func (tds TypeDefSet) Identifier(name string) (types.Entity, bool) {
	for _, e := range tds {
		if e.Name == name {
			return e, true
		}
	}
	return nil, false
}

func (tds TypeDefSet) Iterate() iter.Seq[types.Entity] {
	return func(yield func(types.Entity) bool) {
		for _, td := range tds {
			if !yield(td) {
				return
			}
		}
	}
}
