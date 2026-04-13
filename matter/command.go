package matter

import (
	"iter"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

type CommandDirection uint8

type Command struct {
	entity
	ID          *Number         `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Direction   Interface       `json:"direction,omitempty"`
	Response    *types.DataType `json:"response,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
	Quality     Quality         `json:"quality,omitempty"`
	Access      Access          `json:"access,omitempty"`

	Fields FieldSet `json:"fields,omitempty"`
}

func NewCommand(source asciidoc.Element, parent types.Entity) *Command {
	return &Command{
		entity: entity{source: source, parent: parent},
	}
}

func (c *Command) EntityType() types.EntityType {
	return types.EntityTypeCommand
}

func (c *Command) Equals(e types.Entity) bool {
	oc, ok := e.(*Command)
	if !ok {
		return false
	}
	if c.Direction != oc.Direction {
		return false
	}
	if c.ID.Valid() && oc.ID.Valid() {
		return c.ID.Equals(oc.ID)
	}
	return c.Name == oc.Name
}

func (c *Command) GetConformance() conformance.Set {
	return c.Conformance
}

func (c *Command) Clone() *Command {
	nc := &Command{entity: entity{source: c.source, parent: c.parent}, ID: c.ID.Clone(), Name: c.Name, Description: c.Description, Direction: c.Direction, Response: c.Response, Quality: c.Quality, Access: c.Access}
	if len(c.Conformance) > 0 {
		nc.Conformance = c.Conformance.CloneSet()
	}
	nc.Fields = make(FieldSet, 0, len(c.Fields))
	for _, f := range c.Fields {
		nc.Fields = append(nc.Fields, f.Clone())
	}
	return nc
}

func (c *Command) Inherit(parent *Command) {
	if len(c.Description) == 0 {
		c.Description = parent.Description
	}
	if c.Direction == InterfaceUnknown {
		c.Direction = parent.Direction
	}
	c.Response = parent.Response
	if len(c.Conformance) == 0 {
		c.Conformance = parent.Conformance.CloneSet()
	}
	c.Access.Inherit(parent.Access)
	c.Fields = c.Fields.Inherit(c, parent.Fields)
}

type CommandSet []*Command

func (cs CommandSet) Identifier(name string) (types.Entity, bool) {
	for _, cmd := range cs {
		if cmd.Name == name {
			return cmd, true
		}
	}
	return nil, false
}

func (cs CommandSet) ToEntities() []types.Entity {
	es := make([]types.Entity, 0, len(cs))
	for _, c := range cs {
		es = append(es, c)
	}
	return es
}

func (cs CommandSet) Iterate() iter.Seq[types.Entity] {
	return func(yield func(types.Entity) bool) {
		for _, c := range cs {
			if !yield(c) {
				return
			}
		}
	}
}
