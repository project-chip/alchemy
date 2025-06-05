package conformance

import (
	"fmt"
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

func (o *Mandatory) Eval(context Context) (state ConformanceState, err error) {
	if o.Expression == nil {
		state.State = StateMandatory
		state.Confidence = ConfidenceDefinite
		return

	}
	var t ExpressionResult
	t, err = o.Expression.Eval(context)
	if err != nil {
		return
	}
	switch t.Confidence() {
	case ConfidenceDefinite:
		state.Confidence = ConfidenceDefinite
		if t.IsTrue() {
			state.State = StateMandatory
		}
	case ConfidenceImpossible:
		state.State = StateMandatory
		state.Confidence = ConfidenceImpossible
	case ConfidencePossible:
		state.State = StateMandatory
		state.Confidence = ConfidencePossible
	case ConfidenceUnknown:
	default:
		err = fmt.Errorf("unexpected confidence: %v", t.Confidence())
	}
	return
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
