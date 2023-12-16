package conformance

import (
	"fmt"
	"strings"
)

type LogicalExpression struct {
	Operand string
	Left    ConformanceExpression
	Right   []ConformanceExpression
	Not     bool
}

func NewLogicalExpression(operand string, left ConformanceExpression, right []any) (*LogicalExpression, error) {
	le := &LogicalExpression{Operand: operand, Left: left}
	for _, r := range right {
		rce, ok := r.(ConformanceExpression)
		if !ok {
			return nil, fmt.Errorf("unexpected type in logical expression: %T", r)
		}
		le.Right = append(le.Right, rce)
	}
	return le, nil
}

func (le *LogicalExpression) String() string {

	switch le.Operand {
	case "|":
		var s strings.Builder
		s.WriteRune('(')
		if le.Not {
			s.WriteRune('!')
			s.WriteString(le.Left.String())
			for _, r := range le.Right {
				s.WriteString(" and !")
				s.WriteString(r.String())
			}

		} else {
			s.WriteString(le.Left.String())
			for _, r := range le.Right {
				s.WriteString(" or ")
				s.WriteString(r.String())
			}
		}
		s.WriteRune(')')
		return s.String()
	case "&":
		var s strings.Builder
		s.WriteRune('(')
		s.WriteString(le.Left.String())
		for _, r := range le.Right {
			if le.Not {

				s.WriteString(" or ")
			} else {
				s.WriteString(" and ")

			}
			s.WriteString(r.String())
		}
		s.WriteRune(')')
		return s.String()
	case "^":
		var s strings.Builder
		if le.Not {

			s.WriteString("!")
		}
		s.WriteRune('(')
		s.WriteString(le.Left.String())
		for _, r := range le.Right {
			s.WriteString(" xor ")
			s.WriteString(r.String())
		}
		s.WriteRune(')')
		return s.String()
	default:
		return "unknown operator"
	}
}

func (le *LogicalExpression) Eval(context ConformanceContext) (bool, error) {
	result, err := le.Left.Eval(context)
	if err != nil {
		return false, err
	}
	for _, right := range le.Right {
		r, err := right.Eval(context)
		if err != nil {
			return false, err
		}
		switch le.Operand {
		case "|":
			result = result || r
		case "&":
			result = result && r
		case "^":
			result = (result || r) && !(result && r)
		default:
			return false, fmt.Errorf("unknown operand: %s", le.Operand)
		}

	}
	if le.Not {
		return !result, nil
	}
	return result, nil
}
