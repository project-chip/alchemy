package conformance

import (
	"strings"
)

type Mandatory struct {
	Expression Expression `json:"expression,omitempty"`
}

func (m *Mandatory) Type() Type {
	return TypeMandatory
}

func (m *Mandatory) ASCIIDocString() string {
	if m.Expression != nil {
		return m.Expression.ASCIIDocString()
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

func (m *Mandatory) Eval(context Context) (ConformanceState, error) {
	if m.Expression == nil {
		return ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}, nil
	}
	t, err := m.Expression.Eval(context)
	if err != nil {
		return ConformanceState{State: StateUnknown, Confidence: ConfidenceDefinite}, nil

	}
	if t.IsTrue() {
		return ConformanceState{State: StateMandatory, Confidence: t.Confidence()}, nil
	}
	return ConformanceState{State: StateUnknown, Confidence: ConfidenceDefinite}, nil
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

func (m *Mandatory) Clone() Conformance {
	nm := &Mandatory{}
	if m.Expression != nil {
		nm.Expression = m.Expression.Clone()
	}
	return nm
}
