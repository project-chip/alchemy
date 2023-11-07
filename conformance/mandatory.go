package conformance

import (
	"strings"

	"github.com/hasty/alchemy/matter"
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

func (oc *MandatoryConformance) Eval(context matter.ConformanceContext) (matter.ConformanceState, error) {
	if oc.Expression == nil {
		return matter.ConformanceStateMandatory, nil
	}
	t, err := oc.Expression.Eval(context)
	if err != nil {
		return matter.ConformanceStateUnknown, err
	}
	if t {
		return matter.ConformanceStateMandatory, nil
	}
	return matter.ConformanceStateUnknown, nil
}
