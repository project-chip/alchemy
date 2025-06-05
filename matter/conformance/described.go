package conformance

type Described struct {
}

func (d *Described) Type() Type {
	return TypeDescribed
}

func (d *Described) ASCIIDocString() string {
	return "desc"
}

func (d *Described) Description() string {
	return "described"
}

func (d *Described) Eval(context Context) (ConformanceState, error) {
	return ConformanceState{State: StateOptional, Confidence: ConfidencePossible}, nil
}

func (d *Described) Equal(c Conformance) bool {
	_, ok := c.(*Described)
	return ok
}

func (d *Described) Clone() Conformance {
	return &Described{}
}
