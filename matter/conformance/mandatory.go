package conformance

import (
	"strings"
)

type MandatoryConformance struct {
	Expression ConformanceExpression
}

func (cc *MandatoryConformance) String() string {
	var s strings.Builder
	s.WriteString("mandatory")
	if cc.Expression != nil {
		s.WriteString(" if ")
		s.WriteString(cc.Expression.String())
	}
	return s.String()
}

func (oc *MandatoryConformance) Eval(context ConformanceContext) (ConformanceState, error) {
	if oc.Expression == nil {
		return ConformanceStateMandatory, nil
	}
	t, err := oc.Expression.Eval(context)
	if err != nil {
		return ConformanceStateUnknown, err
	}
	if t {
		return ConformanceStateMandatory, nil
	}
	return ConformanceStateUnknown, nil
}
