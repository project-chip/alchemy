package matter

type FieldSet []*Field

type Field struct {
	ID   *ID       `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
	Type *DataType `json:"type,omitempty"`

	Constraint  Constraint `json:"constraint,omitempty"`
	Quality     Quality    `json:"quality,omitempty"`
	Access      Access     `json:"access,omitempty"`
	Default     string     `json:"default,omitempty"`
	Conformance string     `json:"conformance,omitempty"`
}

func (fs FieldSet) GetField(name string) *Field {
	for _, f := range fs {
		if f.Name == name {
			return f
		}
	}
	return nil
}
