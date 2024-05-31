package asciidoc

type Counter struct {
	position
	raw

	Name         string
	InitialValue string
	Display      bool
}

func NewCounter(name string, initialValue any, display bool) *Counter {
	iv, _ := initialValue.(string)
	return &Counter{Name: name, InitialValue: iv, Display: display}
}

func (Counter) Type() ElementType {
	return ElementTypeInline
}

func (a *Counter) Equals(o Element) bool {
	oa, ok := o.(*Counter)
	if !ok {
		return false
	}
	if a.Name != oa.Name {
		return false
	}
	if a.InitialValue != oa.InitialValue {
		return false
	}
	return a.Display == oa.Display
}
