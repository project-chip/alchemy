package matter

type CommandDirection uint8

const (
	CommandDirectionUnknown CommandDirection = iota
	CommandDirectionClientToServer
	CommandDirectionServerToClient
)

var CommandDirectionNames = map[CommandDirection]string{
	CommandDirectionUnknown:        "Unknown",
	CommandDirectionClientToServer: "client => server",
	CommandDirectionServerToClient: "client <= server",
}

type Command struct {
	ID             *ID              `json:"id,omitempty"`
	Name           string           `json:"name,omitempty"`
	Description    string           `json:"description,omitempty"`
	Direction      CommandDirection `json:"direction,omitempty"`
	Response       string           `json:"response,omitempty"`
	Conformance    string           `json:"conformance,omitempty"`
	Access         Access           `json:"access,omitempty"`
	IsFabricScoped bool             `json:"fabricScoped,omitempty"`

	Fields []*Field `json:"fields,omitempty"`
}

func (c *Command) ModelType() ModelType {
	return ModelTypeCommand
}
