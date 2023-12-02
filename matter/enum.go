package matter

type Enum struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Type        *DataType `json:"type,omitempty"`
	Values      EnumSet   `json:"values,omitempty"`
}

type EnumValue struct {
	Value       string      `json:"value,omitempty"`
	Name        string      `json:"name,omitempty"`
	Summary     string      `json:"summary,omitempty"`
	Conformance Conformance `json:"conformance,omitempty"`
}

func (ev *EnumValue) GetConformance() Conformance {
	return ev.Conformance
}

type EnumSet []*EnumValue

func (es EnumSet) ConformanceReference(name string) HasConformance {
	for _, e := range es {
		if e.Name == name {
			return e
		}
	}
	return nil
}
