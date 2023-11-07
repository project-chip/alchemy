package conformance

import (
	"strings"

	"github.com/hasty/alchemy/matter"
)

type OptionalConformance struct {
	Expression ConformanceExpression
	Choice     *Choice
}

func (cc *OptionalConformance) String() string {
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

func (oc *OptionalConformance) Eval(context matter.ConformanceContext) (matter.ConformanceState, error) {
	if oc.Expression == nil {
		return matter.ConformanceStateOptional, nil
	}
	t, err := oc.Expression.Eval(context)
	if err != nil {
		return matter.ConformanceStateUnknown, err
	}
	if t {
		return matter.ConformanceStateOptional, nil
	}
	return matter.ConformanceStateUnknown, nil
}
