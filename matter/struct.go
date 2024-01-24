package matter

import "github.com/hasty/alchemy/matter/types"

type Struct struct {
	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description,omitempty"`
	Fields        FieldSet      `json:"fields,omitempty"`
	FabricScoping FabricScoping `json:"fabricScoped,omitempty"`
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
	s.Fields = s.Fields.Inherit(parent.Fields)
}

type StructSet []*Struct

func (ss StructSet) Reference(name string) (types.Entity, bool) {
	for _, e := range ss {
		if e.Name == name {
			return e, true
		}
	}
	return nil, false
}
