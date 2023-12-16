package matter

type CommandDirection uint8

type Command struct {
	ID             *Number     `json:"id,omitempty"`
	Name           string      `json:"name,omitempty"`
	Description    string      `json:"description,omitempty"`
	Direction      Interface   `json:"direction,omitempty"`
	Response       string      `json:"response,omitempty"`
	Conformance    Conformance `json:"conformance,omitempty"`
	Access         Access      `json:"access,omitempty"`
	IsFabricScoped bool        `json:"fabricScoped,omitempty"`

	Fields FieldSet `json:"fields,omitempty"`
}

func (c *Command) Entity() Entity {
	return EntityCommand
}

func (c *Command) GetConformance() Conformance {
	return c.Conformance
}

type CommandSet []*Command

func (cs CommandSet) ConformanceReference(name string) HasConformance {
	for _, cmd := range cs {
		if cmd.Name == name {
			return cmd
		}
	}
	return nil
}
