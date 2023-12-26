package conformance

type Disallowed struct {
	raw string
}

func (c *Disallowed) Type() Type {
	return TypeDisallowed
}

func (dc *Disallowed) RawText() string {
	return dc.raw
}

func (dc *Disallowed) String() string {
	return "disallowed"
}

func (id *Disallowed) Eval(context Context) (State, error) {
	return StateDisallowed, nil
}

func (oc *Disallowed) Equal(c Conformance) bool {
	_, ok := c.(*Disallowed)
	return ok
}

func (c *Disallowed) Clone() Conformance {
	return &Disallowed{raw: c.raw}
}
