package matter

type Struct struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Fields      FieldSet `json:"fields,omitempty"`
}
