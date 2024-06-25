package matter

import (
	"log/slog"

	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

type Field struct {
	ID   *Number         `json:"id,omitempty"`
	Name string          `json:"name,omitempty"`
	Type *types.DataType `json:"type,omitempty"`

	Constraint  constraint.Constraint `json:"constraint,omitempty"`
	Quality     Quality               `json:"quality,omitempty"`
	Access      Access                `json:"access,omitempty"`
	Default     string                `json:"default,omitempty"`
	Conformance conformance.Set       `json:"conformance,omitempty"`

	// Hopefully this will go away as we continue disco-balling the spec
	AnonymousType any `json:"anonymousType,omitempty"`

	entity types.EntityType

	Source Source `json:"source,omitempty"`
}

func NewField(source Source) *Field {
	return &Field{entity: types.EntityTypeField, Source: source}
}

func NewAttribute() *Field {
	return &Field{entity: types.EntityTypeAttribute}
}

func (f *Field) GetConformance() conformance.Set {
	return f.Conformance
}

func (f *Field) EntityType() types.EntityType {
	return f.entity
}

func (f *Field) Inherit(parent *Field) {
	slog.Debug("inheriting field", "parent", parent.Name, "parentType", parent.Type, "childType", f.Type)
	if (f.Type == nil || f.Type.BaseType == types.BaseDataTypeUnknown) && parent.Type != nil {
		f.Type = parent.Type.Clone()
	}
	if !constraint.IsBlank(parent.Constraint) {
		if constraint.IsBlank(f.Constraint) {
			f.Constraint = parent.Constraint.Clone()
		}
	}
	if !conformance.IsBlank(parent.Conformance) {
		if conformance.IsBlank(f.Conformance) {
			f.Conformance = parent.Conformance.CloneSet()
		}
	}
	if f.Quality == QualityNone {
		f.Quality = parent.Quality
	}
	if len(f.Default) == 0 {
		f.Default = parent.Default
	}
	if f.entity == types.EntityTypeUnknown && parent.entity != types.EntityTypeUnknown {
		f.entity = parent.entity
	}
	f.Access.Inherit(parent.Access)
}

func (f *Field) Clone() *Field {
	nf := &Field{ID: f.ID.Clone(), Name: f.Name, Quality: f.Quality, Access: f.Access, Default: f.Default, entity: f.entity, Source: f.Source}
	if f.Type != nil {
		nf.Type = f.Type.Clone()
	}
	if f.Constraint != nil {
		nf.Constraint = f.Constraint.Clone()
	}
	if len(f.Conformance) > 0 {
		nf.Conformance = f.Conformance.CloneSet()
	}
	return nf
}

type FieldSet []*Field

func (fs FieldSet) GetField(name string) *Field {
	for _, f := range fs {
		if f.Name == name {
			return f
		}
	}
	return nil
}

func (fs FieldSet) Identifier(name string) (types.Entity, bool) {
	f := fs.GetField(name)
	if f != nil {
		return f, true
	}
	return nil, false
}

func (fs FieldSet) Inherit(parent FieldSet) (nfs FieldSet) {
	nfs = make(FieldSet, len(fs))
	copy(nfs, fs)
	for _, pf := range parent {
		var matching *Field
		for _, f := range nfs {
			if f.ID.Equals(pf.ID) {
				matching = f
				break
			}
		}
		if matching == nil {
			nfs = append(nfs, pf.Clone())
			continue
		}
		matching.Inherit(pf)
	}
	return
}

func (fs FieldSet) ToEntities() []types.Entity {
	es := make([]types.Entity, 0, len(fs))
	for _, f := range fs {
		es = append(es, f)
	}
	return es
}
