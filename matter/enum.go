package matter

type Enum struct {
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Type        *DataType    `json:"type,omitempty"`
	Values      []*EnumValue `json:"values,omitempty"`
}

type EnumValue struct {
	Value       string      `json:"value,omitempty"`
	Name        string      `json:"name,omitempty"`
	Summary     string      `json:"summary,omitempty"`
	Conformance Conformance `json:"conformance,omitempty"`
}
