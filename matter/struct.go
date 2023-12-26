package matter

type Struct struct {
	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description,omitempty"`
	Fields        FieldSet      `json:"fields,omitempty"`
	FabricScoping FabricScoping `json:"fabricScoped,omitempty"`
}

func (*Struct) Entity() Entity {
	return EntityStruct
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
