package matter

type Command struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Direction      string `json:"direction,omitempty"`
	Response       string `json:"response,omitempty"`
	Conformance    string `json:"conformance,omitempty"`
	Access         Access `json:"access,omitempty"`
	IsFabricScoped bool   `json:"fabricScoped,omitempty"`

	Fields []*Field `json:"fields,omitempty"`
}
