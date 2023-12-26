package matter

import "github.com/hasty/alchemy/matter/conformance"

type Event struct {
	ID                *Number           `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	Description       string            `json:"description,omitempty"`
	Priority          string            `json:"priority,omitempty"`
	FabricSensitivity FabricSensitivity `json:"fabricSensitive,omitempty"`
	Conformance       conformance.Set   `json:"conformance,omitempty"`
	Access            Access            `json:"access,omitempty"`

	Fields FieldSet `json:"fields,omitempty"`
}

func (e *Event) GetConformance() conformance.Set {
	return e.Conformance
}

func (e *Event) Entity() Entity {
	return EntityEvent
}

func (e *Event) Clone() *Event {
	ne := &Event{ID: e.ID.Clone(), Name: e.Name, Description: e.Description, Priority: e.Priority, FabricSensitivity: e.FabricSensitivity, Access: e.Access}
	if len(e.Conformance) > 0 {
		ne.Conformance = e.Conformance.CloneSet()
	}
	ne.Fields = make(FieldSet, 0, len(e.Fields))
	for _, f := range e.Fields {
		ne.Fields = append(ne.Fields, f.Clone())
	}
	return ne
}

func (e *Event) Inherit(parent *Event) {
	if len(e.Description) == 0 {
		e.Description = parent.Description
	}
	if len(e.Priority) == 0 {
		e.Priority = parent.Priority
	}
	if e.FabricSensitivity == FabricSensitivityUnknown {
		e.FabricSensitivity = parent.FabricSensitivity
	}
	if len(e.Conformance) == 0 {
		e.Conformance = parent.Conformance.CloneSet()
	}
	e.Access.Inherit(parent.Access)
	e.Fields = e.Fields.Inherit(parent.Fields)
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
