package conformance

type Provisional struct {
}

func (c *Provisional) Type() Type {
	return TypeProvisional
}

func (p *Provisional) ASCIIDocString() string {
	return "P"
}

func (dc *Provisional) Description() string {
	return "provisional"
}

func (id *Provisional) Eval(context Context) (State, error) {
	return StateProvisional, nil
}

func (oc *Provisional) Equal(c Conformance) bool {
	_, ok := c.(*Provisional)
	return ok
}

func (c *Provisional) Clone() Conformance {
	return &Provisional{}
}
