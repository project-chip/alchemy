package matter

type Event struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	Priority        string `json:"priority,omitempty"`
	FabricSensitive bool   `json:"fabricSensitive,omitempty"`
	Conformance     string `json:"conformance,omitempty"`
	Access          Access `json:"access,omitempty"`

	Fields []*Field `json:"fields,omitempty"`
}
