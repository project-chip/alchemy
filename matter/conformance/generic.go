package conformance

type Generic struct {
	raw string
}

func (c *Generic) Type() Type {
	return TypeGeneric
}

func (dc *Generic) RawText() string {
	return dc.raw
}

func (g *Generic) ASCIIDocString() string {
	return g.raw
}

func (dc *Generic) Description() string {
	return dc.raw
}

func (id *Generic) Eval(context Context) (State, error) {
	return StateUnknown, nil
}

func (g *Generic) Equal(c Conformance) bool {
	og, ok := c.(*Generic)
	if !ok {
		return false
	}
	if g.raw != og.raw {
		return false
	}
	return true
}

func (c *Generic) Clone() Conformance {
	return &Generic{raw: c.raw}
}
