package conformance

import "fmt"

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
	return fmt.Sprintf("generic: %s", g.raw)
}

func (g *Generic) Eval(context Context) (ConformanceState, error) {
	return ConformanceState{State: StateUnknown, Confidence: ConfidenceDefinite}, nil
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
