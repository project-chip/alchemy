package asciidoc

type Counter struct {
	position
	raw

	Name         string
	InitialValue string
	Display      CounterVisibility
}

func NewCounter(name string, initialValue any, display CounterVisibility) *Counter {
	iv, _ := initialValue.(string)
	return &Counter{Name: name, InitialValue: iv, Display: display}
}

func (Counter) Type() ElementType {
	return ElementTypeInline
}

func (c *Counter) Equals(o Element) bool {
	oa, ok := o.(*Counter)
	if !ok {
		return false
	}
	if c.Name != oa.Name {
		return false
	}
	if c.InitialValue != oa.InitialValue {
		return false
	}
	return c.Display == oa.Display
}

func (c *Counter) Clone() Element {
	return &Counter{position: c.position, raw: c.raw, Name: c.Name, InitialValue: c.InitialValue}
}

type CounterType uint8

const (
	CounterTypeInteger CounterType = iota
	CounterTypeUpperCase
	CounterTypeLowerCase
)

type CounterVisibility uint8

const (
	CounterVisibilityHidden CounterVisibility = iota
	CounterVisibilityVisible
)

func (cb CounterVisibility) Visible() bool {
	return cb == CounterVisibilityVisible
}

type CounterState struct {
	CounterType CounterType
	Value       int
}
