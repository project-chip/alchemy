package matter

import "github.com/hasty/alchemy/matter/conformance"

type Field struct {
	ID   *Number   `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
	Type *DataType `json:"type,omitempty"`

	Constraint  Constraint      `json:"constraint,omitempty"`
	Quality     Quality         `json:"quality,omitempty"`
	Access      Access          `json:"access,omitempty"`
	Default     string          `json:"default,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`

	entity Entity
}

func NewField() *Field {
	return &Field{entity: EntityField}
}

func NewAttribute() *Field {
	return &Field{entity: EntityAttribute}
}

func (f *Field) GetConformance() conformance.Set {
	return f.Conformance
}

func (f *Field) Entity() Entity {
	return f.entity
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

func (fs FieldSet) Reference(name string) conformance.HasConformance {
	for _, f := range fs {
		if f.Name == name {
			return f
		}
	}
	return nil
}
