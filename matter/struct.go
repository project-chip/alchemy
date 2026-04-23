package matter

import (
	"iter"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/types"
)

type Struct struct {
	entity

	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description,omitempty"`
	Fields        FieldSet      `json:"fields,omitempty"`
	FabricScoping FabricScoping `json:"fabricScoped,omitempty"`
}

func NewStruct(source asciidoc.Element, parent types.Entity) *Struct {
	return &Struct{
		entity: entity{source: source, parent: parent},
	}
}

func (*Struct) EntityType() types.EntityType {
	return types.EntityTypeStruct
}

func (s *Struct) Clone() *Struct {
	ns := &Struct{Name: s.Name, Description: s.Description, FabricScoping: s.FabricScoping}
	ns.Fields = make(FieldSet, 0, len(s.Fields))
	for _, f := range s.Fields {
		ns.Fields = append(ns.Fields, f.Clone())
	}
	return ns
}

func (s *Struct) Inherit(parent *Struct) {
	if len(s.Description) == 0 {
		s.Description = parent.Description
	}
	if s.FabricScoping == FabricScopingUnknown {
		s.FabricScoping = parent.FabricScoping
	}
	s.Fields = s.Fields.Inherit(s, parent.Fields)
}

func (s *Struct) Equals(e types.Entity) bool {
	os, ok := e.(*Struct)
	if !ok {
		return false
	}
	return s.Name == os.Name
}

type StructSet []*Struct

func (ss StructSet) Identifier(name string) (types.Entity, bool) {
	for _, e := range ss {
		if e.Name == name {
			return e, true
		}
	}
	return nil, false
}

func (ss StructSet) Iterate() iter.Seq[types.Entity] {
	return func(yield func(types.Entity) bool) {
		for _, s := range ss {
			if !yield(s) {
				return
			}
		}
	}
}
