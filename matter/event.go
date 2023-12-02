package matter

type Event struct {
	ID              *Number     `json:"id,omitempty"`
	Name            string      `json:"name,omitempty"`
	Description     string      `json:"description,omitempty"`
	Priority        string      `json:"priority,omitempty"`
	FabricSensitive bool        `json:"fabricSensitive,omitempty"`
	Conformance     Conformance `json:"conformance,omitempty"`
	Access          Access      `json:"access,omitempty"`

	Fields FieldSet `json:"fields,omitempty"`
}

func (e *Event) GetConformance() Conformance {
	return e.Conformance
}

type EventSet []*Event

func (es EventSet) ConformanceReference(name string) HasConformance {
	for _, e := range es {
		if e.Name == name {
			return e
		}
	}
	return nil
}
