package conformance

type Provisional struct {
}

func (p *Provisional) Type() Type {
	return TypeProvisional
}

func (p *Provisional) ASCIIDocString() string {
	return "P"
}

func (p *Provisional) Description() string {
	return "provisional"
}

func (p *Provisional) Eval(context Context) (ConformanceState, error) {
	return ConformanceState{State: StateProvisional, Confidence: ConfidenceDefinite}, nil
}

func (p *Provisional) Equal(c Conformance) bool {
	_, ok := c.(*Provisional)
	return ok
}

func (p *Provisional) Clone() Conformance {
	return &Provisional{}
}
