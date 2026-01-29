package conformance

type Obsolete struct {
	raw string
}

func (d *Obsolete) Type() Type {
	return TypeDeprecated
}

func (d *Obsolete) RawText() string {
	return d.raw
}

func (d *Obsolete) ASCIIDocString() string {
	return "Z"
}

func (d *Obsolete) Description() string {
	return "obsolete"
}

func (d *Obsolete) Eval(context Context) (ConformanceState, error) {
	return ConformanceState{State: StateDeprecated, Confidence: ConfidenceDefinite}, nil
}

func (d *Obsolete) Equal(c Conformance) bool {
	_, ok := c.(*Obsolete)
	return ok
}

func (d *Obsolete) Clone() Conformance {
	return &Obsolete{raw: d.raw}
}
