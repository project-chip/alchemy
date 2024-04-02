package conformance

type Generic struct {
	raw string
}

func (g *Generic) Type() Type {
	return TypeGeneric
}

func (g *Generic) RawText() string {
	return g.raw
}

func (g *Generic) ASCIIDocString() string {
	return g.raw
}

func (g *Generic) Description() string {
	return g.raw
}

func (g *Generic) Eval(context Context) (State, error) {
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

func (g *Generic) Clone() Conformance {
	return &Generic{raw: g.raw}
}
