package matter

import (
	"slices"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

type Enum struct {
	entity

	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Type        *types.DataType `json:"type,omitempty"`
	Values      EnumValueSet    `json:"values,omitempty"`
}

func NewEnum(source asciidoc.Element, parent types.Entity) *Enum {
	return &Enum{
		entity: entity{source: source, parent: parent},
	}
}

func (*Enum) EntityType() types.EntityType {
	return types.EntityTypeEnum
}

func (e *Enum) BaseDataType() types.BaseDataType {
	return e.Type.BaseType
}

func (e *Enum) NullValue() uint64 {
	return e.Type.NullValue()
}

func (e *Enum) Clone() *Enum {
	ne := &Enum{entity: entity{source: e.source, parent: e.parent}, Name: e.Name, Description: e.Description}
	if e.Type != nil {
		ne.Type = e.Type.Clone()
	}
	ne.Values = make(EnumValueSet, 0, len(e.Values))
	for _, ev := range e.Values {
		ne.Values = append(ne.Values, ev.Clone())
	}
	return ne
}

func (e *Enum) Inherit(parent *Enum) error {
	mergedValues := make(EnumValueSet, 0, len(parent.Values))
	for _, ev := range parent.Values {
		mergedValues = append(mergedValues, ev.Clone())
	}
	for _, ev := range e.Values {
		var matching *EnumValue
		for _, mev := range mergedValues {
			if ev.Name == mev.Name {
				matching = mev
				break
			}
		}
		if matching == nil {
			mergedValues = append(mergedValues, ev.Clone())
			continue
		}
		if len(ev.Summary) > 0 {
			matching.Summary = ev.Summary
		}
		if len(ev.Conformance) > 0 {
			matching.Conformance = ev.Conformance.CloneSet()
		}
	}
	if e.Type == nil {
		e.Type = parent.Type
	}
	if len(e.Description) == 0 {
		e.Description = parent.Description
	}
	slices.SortStableFunc(mergedValues, func(a, b *EnumValue) int {
		return a.Value.Compare(b.Value)
	})
	e.Values = mergedValues
	for _, ev := range e.Values {
		ev.parent = e
	}
	return nil
}

type EnumSet []*Enum

func (es EnumSet) Identifier(name string) (types.Entity, bool) {
	for _, e := range es {
		if e.Name == name {
			return e, true
		}
	}
	return nil, false
}

type EnumValue struct {
	entity
	Value       *Number         `json:"value,omitempty"`
	Name        string          `json:"name,omitempty"`
	Summary     string          `json:"summary,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
}

func NewEnumValue(source asciidoc.Element, parent types.Entity) *EnumValue {
	return &EnumValue{
		entity: entity{source: source, parent: parent},
	}
}

func (ev *EnumValue) EntityType() types.EntityType {
	return types.EntityTypeEnumValue
}

func (ev *EnumValue) Clone() *EnumValue {
	nev := &EnumValue{entity: entity{source: ev.source}, Name: ev.Name, Value: ev.Value.Clone(), Summary: ev.Summary}
	if len(ev.Conformance) > 0 {
		nev.Conformance = ev.Conformance.CloneSet()
	}
	return nev
}

func (ev *EnumValue) GetConformance() conformance.Set {
	return ev.Conformance
}

type EnumValueSet []*EnumValue

func (es EnumValueSet) Identifier(name string) (types.Entity, bool) {
	for _, e := range es {
		if e.Name == name {
			return e, true
		}
	}
	return nil, false
}

func NewAnonymousEnum(source asciidoc.Element, parent types.Entity) *AnonymousEnum {
	return &AnonymousEnum{
		entity: entity{source: source, parent: parent},
	}
}

type AnonymousEnum struct {
	entity
	Type   *types.DataType `json:"type,omitempty"`
	Values EnumValueSet    `json:"values,omitempty"`
}

func (AnonymousEnum) EntityType() types.EntityType {
	return types.EntityTypeEnum
}
