package conformance

type Described struct {
}

func (c *Described) Type() Type {
	return TypeDescribed
}

func (d *Described) AsciiDocString() string {
	return "desc"
}

func (d *Described) Description() string {
	return "described"
}

func (d *Described) Eval(context Context) (State, error) {
	return StateUnknown, nil
}

func (d *Described) Equal(c Conformance) bool {
	_, ok := c.(*Described)
	return ok
}

func (c *Described) Clone() Conformance {
	return &Described{}
}
