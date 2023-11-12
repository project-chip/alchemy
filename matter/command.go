package matter

type CommandDirection uint8

type Command struct {
	ID             *ID       `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Description    string    `json:"description,omitempty"`
	Direction      Interface `json:"direction,omitempty"`
	Response       string    `json:"response,omitempty"`
	Conformance    string    `json:"conformance,omitempty"`
	Access         Access    `json:"access,omitempty"`
	IsFabricScoped bool      `json:"fabricScoped,omitempty"`

	Fields []*Field `json:"fields,omitempty"`
}

func (c *Command) ModelType() Entity {
	return EntityCommand
}
