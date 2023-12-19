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

func (dc *Deprecated) String() string {
	return "deprecated"
}

func (dc *Deprecated) Eval(context Context) (State, error) {
	return StateDeprecated, nil
}

func (oc *Deprecated) Equal(c Conformance) bool {
	_, ok := c.(*Deprecated)
	return ok
}
