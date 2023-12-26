package matter

import "github.com/hasty/alchemy/matter/conformance"

type CommandDirection uint8

type Command struct {
	ID            *Number         `json:"id,omitempty"`
	Name          string          `json:"name,omitempty"`
	Description   string          `json:"description,omitempty"`
	Direction     Interface       `json:"direction,omitempty"`
	Response      string          `json:"response,omitempty"`
	Conformance   conformance.Set `json:"conformance,omitempty"`
	Access        Access          `json:"access,omitempty"`
	FabricScoping FabricScoping   `json:"fabricScoped,omitempty"`

	Fields FieldSet `json:"fields,omitempty"`
}

func (c *Command) Entity() Entity {
	return EntityCommand
}

func (c *Command) GetConformance() conformance.Set {
	return c.Conformance
}

func (c *Command) Clone() *Command {
	nc := &Command{ID: c.ID.Clone(), Name: c.Name, Description: c.Description, Direction: c.Direction, Response: c.Response, Access: c.Access, FabricScoping: c.FabricScoping}
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
	if c.FabricScoping == FabricScopingUnknown {
		c.FabricScoping = parent.FabricScoping
	}
	if len(c.Response) == 0 {
		c.Response = parent.Response
	}
	if len(c.Conformance) == 0 {
		c.Conformance = parent.Conformance.CloneSet()
	}
	c.Access.Inherit(parent.Access)
	c.Fields = c.Fields.Inherit(parent.Fields)
}

type CommandSet []*Command

func (cs CommandSet) Reference(name string) conformance.HasConformance {
	for _, cmd := range cs {
		if cmd.Name == name {
			return cmd
		}
	}
	return nil
}
