package conformance

type Disallowed struct {
	raw string
}

func (d *Disallowed) Type() Type {
	return TypeDisallowed
}

func (d *Disallowed) RawText() string {
	return d.raw
}

func (d *Disallowed) ASCIIDocString() string {
	return "X"
}

func (d *Disallowed) Description() string {
	return "disallowed"
}

func (d *Disallowed) Eval(context Context) (State, error) {
	return StateDisallowed, nil
}

func (d *Disallowed) Equal(c Conformance) bool {
	_, ok := c.(*Disallowed)
	return ok
}

func (d *Disallowed) Clone() Conformance {
	return &Disallowed{raw: d.raw}
}
