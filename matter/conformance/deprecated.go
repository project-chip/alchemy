package conformance

type Deprecated struct {
	raw string
}

func (c *Deprecated) Type() Type {
	return TypeDeprecated
}

func (dc *Deprecated) RawText() string {
	return dc.raw
}

func (dc *Deprecated) AsciiDocString() string {
	return "D"
}

func (dc *Deprecated) Description() string {
	return "deprecated"
}

func (dc *Deprecated) Eval(context Context) (State, error) {
	return StateDeprecated, nil
}

func (oc *Deprecated) Equal(c Conformance) bool {
	_, ok := c.(*Deprecated)
	return ok
}

func (c *Deprecated) Clone() Conformance {
	return &Deprecated{raw: c.raw}
}
