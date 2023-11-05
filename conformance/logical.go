package conformance

import (
	"fmt"

	"github.com/hasty/matterfmt/matter"
)

type LogicalExpression struct {
	Operand string
	Left    ConformanceExpression
	Right   ConformanceExpression
}

func (le *LogicalExpression) String() string {
	switch le.Operand {
	case "|":
		return fmt.Sprintf("(%s or %s)", le.Left.String(), le.Right.String())
	case "&":
		return fmt.Sprintf("(%s and %s)", le.Left.String(), le.Right.String())
	case "^":
		return fmt.Sprintf("(%s xor %s)", le.Left.String(), le.Right.String())
	default:
		return "unknown operator"
	}
}

func (le *LogicalExpression) Eval(context matter.ConformanceContext) (bool, error) {
	l, err := le.Left.Eval(context)
	if err != nil {
		return false, err
	}
	r, err := le.Right.Eval(context)
	if err != nil {
		return false, err
	}
	switch le.Operand {
	case "|":
		return l || r, nil
	case "&":
		return l && r, nil
	case "^":
		return (l || r) && !(l && r), nil
	default:
		return false, fmt.Errorf("unknown operand: %s", le.Operand)
	}
}
