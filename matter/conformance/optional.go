package conformance

import (
	"strings"
)

type Optional struct {
	Expression Expression `json:"expression,omitempty"`
	Choice     *Choice    `json:"choice,omitempty"`
}

func (c *Optional) Type() Type {
	return TypeOptional
}

func (cc *Optional) String() string {
	var s strings.Builder
	s.WriteString("optional")
	if cc.Expression != nil {
		s.WriteString(" if ")
		s.WriteString(cc.Expression.String())
	}
	if cc.Choice != nil {
		s.WriteString(" (")
		s.WriteString(cc.Choice.String())
		s.WriteString(")")
	}
	return s.String()
}

func (oc *Optional) Eval(context Context) (State, error) {
	if oc.Expression == nil {
		return StateOptional, nil
	}
	t, err := oc.Expression.Eval(context)
	if err != nil {
		return StateUnknown, err
	}
	if t {
		return StateOptional, nil
	}
	return StateUnknown, nil
}

func (oc *Optional) Equal(c Conformance) bool {
	ooc, ok := c.(*Optional)
	if !ok {
		return false
	}
	if !oc.Choice.Equal(ooc.Choice) {
		return false
	}
	if !oc.Expression.Equal(ooc.Expression) {
		return false
	}
	return true
}

func (c *Optional) Clone() Conformance {
	nm := &Optional{}
	if c.Expression != nil {
		nm.Expression = c.Expression.Clone()
	}
	if c.Choice != nil {
		nm.Choice = c.Choice.Clone()
	}
	return nm
}
