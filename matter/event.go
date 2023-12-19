package matter

import "github.com/hasty/alchemy/matter/conformance"

type Event struct {
	ID              *Number         `json:"id,omitempty"`
	Name            string          `json:"name,omitempty"`
	Description     string          `json:"description,omitempty"`
	Priority        string          `json:"priority,omitempty"`
	FabricSensitive bool            `json:"fabricSensitive,omitempty"`
	Conformance     conformance.Set `json:"conformance,omitempty"`
	Access          Access          `json:"access,omitempty"`

	Fields FieldSet `json:"fields,omitempty"`
}

func (e *Event) GetConformance() conformance.Set {
	return e.Conformance
}

func (e *Event) Entity() Entity {
	return EntityEvent
}

type EventSet []*Event

func (es EventSet) Reference(name string) conformance.HasConformance {
	for _, e := range es {
		if e.Name == name {
			return e
		}
	}
	return nil
}
