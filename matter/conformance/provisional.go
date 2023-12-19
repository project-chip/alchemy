package conformance

type Provisional struct {
}

func (c *Provisional) Type() Type {
	return TypeProvisional
}

func (dc *Provisional) String() string {
	return "provisional"
}

func (id *Provisional) Eval(context Context) (State, error) {
	return StateProvisional, nil
}

func (oc *Provisional) Equal(c Conformance) bool {
	_, ok := c.(*Provisional)
	return ok
}
