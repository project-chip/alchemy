package matter

import (
	"slices"

	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/types"
)

type Enum struct {
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Type        *types.DataType `json:"type,omitempty"`
	Values      EnumSet         `json:"values,omitempty"`
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
	ne := &Enum{Name: e.Name, Description: e.Description}
	if e.Type != nil {
		ne.Type = e.Type.Clone()
	}
	ne.Values = make(EnumSet, 0, len(e.Values))
	for _, ev := range e.Values {
		ne.Values = append(ne.Values, ev.Clone())
	}
	return ne
}

func (en *Enum) Inherit(parent *Enum) error {
	mergedValues := make(EnumSet, 0, len(parent.Values))
	for _, ev := range parent.Values {
		mergedValues = append(mergedValues, ev.Clone())
	}
	for _, ev := range en.Values {
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
	if en.Type == nil {
		en.Type = parent.Type
	}
	if len(en.Description) == 0 {
		en.Description = parent.Description
	}
	slices.SortFunc(mergedValues, func(a, b *EnumValue) int {
		return a.Value.Compare(b.Value)
	})
	en.Values = mergedValues
	return nil
}

type EnumValue struct {
	Value       *Number         `json:"value,omitempty"`
	Name        string          `json:"name,omitempty"`
	Summary     string          `json:"summary,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
}

func (ev *EnumValue) EntityType() types.EntityType {
	return types.EntityTypeEnumValue
}

func (ev *EnumValue) Clone() *EnumValue {
	nev := &EnumValue{Name: ev.Name, Value: ev.Value.Clone(), Summary: ev.Summary}
	if len(ev.Conformance) > 0 {
		nev.Conformance = ev.Conformance.CloneSet()
	}
	return nev
}

func (ev *EnumValue) GetConformance() conformance.Set {
	return ev.Conformance
}

type EnumSet []*EnumValue

func (es EnumSet) Reference(name string) conformance.HasConformance {
	for _, e := range es {
		if e.Name == name {
			return e
		}
	}
	return nil
}
