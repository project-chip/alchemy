package conformance

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type LogicalExpression struct {
	Operand string
	Left    ConformanceExpression
	Right   ConformanceExpression
	Not     bool
}

func (le *LogicalExpression) String() string {
	switch le.Operand {
	case "|":
		if le.Not {
			return fmt.Sprintf("!%s and !%s", le.Left.String(), le.Right.String())
		}
		return fmt.Sprintf("(%s or %s)", le.Left.String(), le.Right.String())
	case "&":
		if le.Not {
			return fmt.Sprintf("(%s or %s)", le.Left.String(), le.Right.String())
		}
		return fmt.Sprintf("(%s and %s)", le.Left.String(), le.Right.String())
	case "^":
		if le.Not {
			return fmt.Sprintf("!(%s xor %s)", le.Left.String(), le.Right.String())
		}
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
	var result bool
	switch le.Operand {
	case "|":
		result = l || r
	case "&":
		result = l && r
	case "^":
		result = (l || r) && !(l && r)
	default:
		return false, fmt.Errorf("unknown operand: %s", le.Operand)
	}
	if le.Not {
		return !result, nil
	}
	return result, nil
}
