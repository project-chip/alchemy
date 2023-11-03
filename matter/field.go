package matter

type Field struct {
	ID   string    `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
	Type *DataType `json:"type,omitempty"`

	Constraint  Constraint `json:"constraint,omitempty"`
	Quality     Quality    `json:"quality,omitempty"`
	Access      *Access    `json:"access,omitempty"`
	Default     string     `json:"default,omitempty"`
	Conformance string     `json:"conformance,omitempty"`
}
