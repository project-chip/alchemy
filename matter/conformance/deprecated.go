package conformance

type Deprecated struct {
	raw string
}

func (d *Deprecated) Type() Type {
	return TypeDeprecated
}

func (d *Deprecated) RawText() string {
	return d.raw
}

func (d *Deprecated) ASCIIDocString() string {
	return "D"
}

func (d *Deprecated) Description() string {
	return "deprecated"
}

func (d *Deprecated) Eval(context Context) (ConformanceState, error) {
	return ConformanceState{State: StateDeprecated, Confidence: ConfidenceDefinite}, nil
}

func (d *Deprecated) Equal(c Conformance) bool {
	_, ok := c.(*Deprecated)
	return ok
}

func (d *Deprecated) Clone() Conformance {
	return &Deprecated{raw: d.raw}
}
