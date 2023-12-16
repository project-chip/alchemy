package matter

type Struct struct {
	Name         string   `json:"name,omitempty"`
	Description  string   `json:"description,omitempty"`
	Fields       FieldSet `json:"fields,omitempty"`
	FabricScoped bool     `json:"fabricScoped,omitempty"`
}

func (*Struct) Entity() Entity {
	return EntityStruct
}
