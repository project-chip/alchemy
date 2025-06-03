package conformance

import (
	"strings"

	"github.com/project-chip/alchemy/internal/text"
)

type Optional struct {
	Expression Expression `json:"expression,omitempty"`
	Choice     *Choice    `json:"choice,omitempty"`
}

func (o *Optional) Type() Type {
	return TypeOptional
}

func (o *Optional) ASCIIDocString() string {
	var s strings.Builder
	if o.Expression != nil {
		s.WriteString("[")
		s.WriteString(text.TrimUnnecessaryParens(o.Expression.ASCIIDocString()))
		s.WriteString("]")
	} else {
		s.WriteString("O")
	}
	if o.Choice != nil {
		s.WriteString(".")
		s.WriteString(o.Choice.ASCIIDocString())
	}
	return s.String()
}

func (o *Optional) Description() string {
	var s strings.Builder
	s.WriteString("optional")
	if o.Expression != nil {
		s.WriteString(" if ")
		s.WriteString(o.Expression.Description())
	}
	if o.Choice != nil {
		s.WriteString(" (")
		s.WriteString(o.Choice.Description())
		s.WriteString(")")
	}
	return s.String()
}

func (o *Optional) Eval(context Context) (State, error) {
	if o.Expression == nil {
		return StateOptional, nil
	}
	t, err := o.Expression.Eval(context)
	if err != nil {
		return StateUnknown, err
	}
	if t.IsTrue() {
		return StateOptional, nil
	}
	return StateUnknown, nil
}

func (o *Optional) Equal(c Conformance) bool {
	oc, ok := c.(*Optional)
	if !ok {
		return false
	}
	if o.Choice != nil {
		if oc.Choice == nil {
			return false
		}
		if !o.Choice.Equal(oc.Choice) {
			return false
		}
	} else if oc.Choice != nil {
		return false
	}
	if o.Expression != nil {
		if oc.Expression == nil {
			return false
		}
		if !o.Expression.Equal(oc.Expression) {
			return false
		}
	} else if oc.Expression != nil {
		return false
	}
	return true
}

func (o *Optional) Clone() Conformance {
	nm := &Optional{}
	if o.Expression != nil {
		nm.Expression = o.Expression.Clone()
	}
	if o.Choice != nil {
		nm.Choice = o.Choice.Clone()
	}
	return nm
}
