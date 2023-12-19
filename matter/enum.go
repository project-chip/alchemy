package matter

import "github.com/hasty/alchemy/matter/conformance"

type Enum struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Type        *DataType `json:"type,omitempty"`
	Values      EnumSet   `json:"values,omitempty"`
}

func (*Enum) Entity() Entity {
	return EntityEnum
}

type EnumValue struct {
	Value       string          `json:"value,omitempty"`
	Name        string          `json:"name,omitempty"`
	Summary     string          `json:"summary,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
}

func (ev *EnumValue) Entity() Entity {
	return EntityEnumValue
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
