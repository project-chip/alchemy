package matter

type Field struct {
	ID   *Number   `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
	Type *DataType `json:"type,omitempty"`

	Constraint  Constraint  `json:"constraint,omitempty"`
	Quality     Quality     `json:"quality,omitempty"`
	Access      Access      `json:"access,omitempty"`
	Default     string      `json:"default,omitempty"`
	Conformance Conformance `json:"conformance,omitempty"`
}

func (f *Field) GetConformance() Conformance {
	return f.Conformance
}

type FieldSet []*Field

func (fs FieldSet) GetField(name string) *Field {
	for _, f := range fs {
		if f.Name == name {
			return f
		}
	}
	return nil
}

func (fs FieldSet) ConformanceReference(name string) HasConformance {
	for _, f := range fs {
		if f.Name == name {
			return f
		}
	}
	return nil
}
