package conformance

import (
	"strings"
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

func (oc *OptionalConformance) Eval(context ConformanceContext) (ConformanceState, error) {
	if oc.Expression == nil {
		return ConformanceStateOptional, nil
	}
	t, err := oc.Expression.Eval(context)
	if err != nil {
		return ConformanceStateUnknown, err
	}
	if t {
		return ConformanceStateOptional, nil
	}
	return ConformanceStateUnknown, nil
}
