package conformance

import (
	"strings"
)

type Mandatory struct {
	Expression Expression `json:"expression,omitempty"`
}

func (c *Mandatory) Type() Type {
	return TypeMandatory
}

func (m *Mandatory) AsciiDocString() string {
	if m.Expression != nil {
		return m.Expression.AsciiDocString()
	}
	return "M"
}

func (m *Mandatory) Description() string {
	var s strings.Builder
	s.WriteString("mandatory")
	if m.Expression != nil {
		s.WriteString(" if ")
		s.WriteString(m.Expression.Description())
	}
	return s.String()
}

func (m *Mandatory) Eval(context Context) (State, error) {
	if m.Expression == nil {
		return StateMandatory, nil
	}
	t, err := m.Expression.Eval(context)
	if err != nil {
		return StateUnknown, err
	}
	if t {
		return StateMandatory, nil
	}
	return StateUnknown, nil
}

func (m *Mandatory) Equal(c Conformance) bool {
	om, ok := c.(*Mandatory)
	if !ok {
		return false
	}
	if m.Expression != nil {
		if om.Expression == nil {
			return false
		}
		if !m.Expression.Equal(om.Expression) {
			return false
		}
	} else if om.Expression != nil {
		return false
	}
	return true
}

func (c *Mandatory) Clone() Conformance {
	nm := &Mandatory{}
	if c.Expression != nil {
		nm.Expression = c.Expression.Clone()
	}
	return nm
}
