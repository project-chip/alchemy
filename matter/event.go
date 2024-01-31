package matter

import (
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/types"
)

type Event struct {
	ID          *Number         `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Priority    string          `json:"priority,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
	Access      Access          `json:"access,omitempty"`

	Fields FieldSet `json:"fields,omitempty"`
}

func (e *Event) GetConformance() conformance.Set {
	return e.Conformance
}

func (e *Event) EntityType() types.EntityType {
	return types.EntityTypeEvent
}

func (e *Event) Clone() *Event {
	ne := &Event{ID: e.ID.Clone(), Name: e.Name, Description: e.Description, Priority: e.Priority, Access: e.Access}
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
	if e.Access.FabricSensitivity == FabricSensitivityUnknown {
		e.Access.FabricSensitivity = parent.Access.FabricSensitivity
	}
	if len(e.Conformance) == 0 {
		e.Conformance = parent.Conformance.CloneSet()
	}
	e.Access.Inherit(parent.Access)
	e.Fields = e.Fields.Inherit(parent.Fields)
}

type EventSet []*Event

func (es EventSet) Identifier(name string) (types.Entity, bool) {
	for _, e := range es {
		if e.Name == name {
			return e, true
		}
	}
	return nil, false
}
