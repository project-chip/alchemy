package matter

import (
	"iter"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

type Event struct {
	entity
	ID          *Number         `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Priority    string          `json:"priority,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
	Access      Access          `json:"access,omitempty"`

	Fields FieldSet `json:"fields,omitempty"`
}

func NewEvent(source asciidoc.Element, parent types.Entity) *Event {
	return &Event{
		entity: entity{source: source, parent: parent},
	}
}

func (e *Event) GetConformance() conformance.Set {
	return e.Conformance
}

func (e *Event) EntityType() types.EntityType {
	return types.EntityTypeEvent
}

func (ev *Event) Equals(e types.Entity) bool {
	oev, ok := e.(*Event)
	if !ok {
		return false
	}
	if ev.ID.Valid() && oev.ID.Valid() {
		return ev.ID.Equals(oev.ID)
	}
	return ev.Name == oev.Name
}

func (e *Event) Clone() *Event {
	ne := &Event{entity: entity{source: e.source}, ID: e.ID.Clone(), Name: e.Name, Description: e.Description, Priority: e.Priority, Access: e.Access}
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
	e.Fields = e.Fields.Inherit(e, parent.Fields)
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

func (es EventSet) ToEntities() []types.Entity {
	evs := make([]types.Entity, 0, len(es))
	for _, e := range es {
		evs = append(evs, e)
	}
	return evs
}

func (es EventSet) Iterate() iter.Seq[types.Entity] {
	return func(yield func(types.Entity) bool) {
		for _, ev := range es {
			if !yield(ev) {
				return
			}
		}
	}
}
