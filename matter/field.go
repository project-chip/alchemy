package matter

import (
	"iter"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

type Field struct {
	entity
	ID   *Number         `json:"id,omitempty"`
	Name string          `json:"name,omitempty"`
	Type *types.DataType `json:"type,omitempty"`

	Constraint  constraint.Constraint `json:"constraint,omitempty"`
	Quality     Quality               `json:"quality,omitempty"`
	Access      Access                `json:"access,omitempty"`
	Fallback    constraint.Limit      `json:"fallback,omitempty"`
	Conformance conformance.Set       `json:"conformance,omitempty"`

	// Hopefully this will go away as we continue disco-balling the spec
	AnonymousType types.Entity `json:"anonymousType,omitempty"`

	entityType types.EntityType
}

func NewField(source asciidoc.Element, parent types.Entity, entityType types.EntityType) *Field {
	return &Field{entity: entity{source: source, parent: parent}, entityType: entityType}
}

func NewAttribute(source asciidoc.Element, parent types.Entity) *Field {
	return &Field{entity: entity{source: source, parent: parent}, entityType: types.EntityTypeAttribute}
}

func (f *Field) GetConformance() conformance.Set {
	return f.Conformance
}

func (f *Field) EntityType() types.EntityType {
	return f.entityType
}

func (f *Field) Equals(e types.Entity) bool {
	of, ok := e.(*Field)
	if !ok {
		return false
	}
	if f.ID.Valid() && of.ID.Valid() {
		return f.ID.Equals(of.ID)
	}
	return f.Name == of.Name
}

func (f *Field) Inherit(parent *Field) {
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
	if constraint.IsBlankLimit(f.Fallback) {
		f.Fallback = parent.Fallback
	}
	if f.entityType == types.EntityTypeUnknown && parent.entityType != types.EntityTypeUnknown {
		f.entityType = parent.entityType
	}
	f.Access.Inherit(parent.Access)
}

func (f *Field) Clone() *Field {
	nf := &Field{
		entity: entity{
			source: f.source,
			parent: f.parent,
		},
		ID:         f.ID.Clone(),
		Name:       f.Name,
		Quality:    f.Quality,
		Access:     f.Access,
		Fallback:   f.Fallback,
		entityType: f.entityType,
	}
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

func (f *Field) IterateDataTypes() EntityIterator {
	return iterateOverFieldDataTypes(f)
}

type FieldSet []*Field

func (fs FieldSet) Get(name string) *Field {
	for _, f := range fs {
		if f.Name == name {
			return f
		}
	}
	return nil
}

func (fs FieldSet) Identifier(name string) (types.Entity, bool) {
	f := fs.Get(name)
	if f != nil {
		return f, true
	}
	return nil, false
}

func (fs FieldSet) Inherit(parent types.Entity, fields FieldSet) (nfs FieldSet) {
	nfs = make(FieldSet, len(fs))
	copy(nfs, fs)
	for _, pf := range fields {
		var matching *Field
		for _, f := range nfs {
			if f.ID.Equals(pf.ID) {
				matching = f
				break
			}
		}
		if matching == nil {
			cf := pf.Clone()
			cf.parent = parent
			nfs = append(nfs, cf)
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

func (fs FieldSet) Iterate() iter.Seq[types.Entity] {
	return func(yield func(types.Entity) bool) {
		for _, f := range fs {
			if !yield(f) {
				return
			}
		}
	}
}
