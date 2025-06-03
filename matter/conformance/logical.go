package conformance

import (
	"encoding/json"
	"fmt"
	"strings"
)

type LogicalExpression struct {
	Operand string
	Left    Expression
	Right   []Expression
	Not     bool
}

func NewLogicalExpression(operand string, left Expression, right []any) (*LogicalExpression, error) {
	le := &LogicalExpression{Operand: operand, Left: left}
	for _, r := range right {
		rce, ok := r.(Expression)
		if !ok {
			return nil, fmt.Errorf("unexpected type in logical expression: %T", r)
		}
		le.Right = append(le.Right, rce)
	}
	return le, nil
}

func (le *LogicalExpression) ASCIIDocString() string {

	switch le.Operand {
	case "|":
		var s strings.Builder
		s.WriteRune('(')
		if le.Not {
			s.WriteRune('!')
			s.WriteString(le.Left.ASCIIDocString())
			for _, r := range le.Right {
				s.WriteString(" & !")
				s.WriteString(r.ASCIIDocString())
			}

		} else {
			s.WriteString(le.Left.ASCIIDocString())
			for _, r := range le.Right {
				s.WriteString(" \\| ")
				s.WriteString(r.ASCIIDocString())
			}
		}
		s.WriteRune(')')
		return s.String()
	case "&":
		var s strings.Builder
		s.WriteRune('(')
		s.WriteString(le.Left.ASCIIDocString())
		for _, r := range le.Right {
			if le.Not {
				s.WriteString(" \\| ")
			} else {
				s.WriteString(" & ")

			}
			s.WriteString(r.ASCIIDocString())
		}
		s.WriteRune(')')
		return s.String()
	case "^":
		var s strings.Builder
		if le.Not {
			s.WriteString("!")
		}
		s.WriteRune('(')
		s.WriteString(le.Left.ASCIIDocString())
		for _, r := range le.Right {
			s.WriteString(" ^ ")
			s.WriteString(r.ASCIIDocString())
		}
		s.WriteRune(')')
		return s.String()
	default:
		return "unknown operator"
	}
}

func (le *LogicalExpression) Description() string {

	switch le.Operand {
	case "|":
		var s strings.Builder
		s.WriteRune('(')
		if le.Not {
			s.WriteString("not ")
			s.WriteString(le.Left.Description())
			for _, r := range le.Right {
				s.WriteString(" and not")
				s.WriteString(r.Description())
			}

		} else {
			s.WriteString(le.Left.Description())
			for _, r := range le.Right {
				s.WriteString(" or ")
				s.WriteString(r.Description())
			}
		}
		s.WriteRune(')')
		return s.String()
	case "&":
		var s strings.Builder
		s.WriteRune('(')
		s.WriteString(le.Left.Description())
		for _, r := range le.Right {
			if le.Not {

				s.WriteString(" or ")
			} else {
				s.WriteString(" and ")

			}
			s.WriteString(r.Description())
		}
		s.WriteRune(')')
		return s.String()
	case "^":
		var s strings.Builder
		if le.Not {

			s.WriteString("!")
		}
		s.WriteRune('(')
		s.WriteString(le.Left.Description())
		for _, r := range le.Right {
			s.WriteString(" xor ")
			s.WriteString(r.Description())
		}
		s.WriteRune(')')
		return s.String()
	default:
		return "unknown operator"
	}
}

func (le *LogicalExpression) Eval(context Context) (ExpressionResult, error) {
	l, err := le.Left.Eval(context)
	if err != nil {
		return nil, err
	}
	for _, right := range le.Right {
		r, err := right.Eval(context)
		if err != nil {
			return nil, err
		}
		switch le.Operand {
		case "|":
			l = &expressionResult{value: l.IsTrue() || r.IsTrue(), confidence: coalesceConfidences(l.Confidence(), r.Confidence())}
		case "&":
			l = &expressionResult{value: l.IsTrue() && r.IsTrue(), confidence: coalesceConfidences(l.Confidence(), r.Confidence())}
		case "^":
			l = &expressionResult{value: (l.IsTrue() || r.IsTrue()) && !(l.IsTrue() && r.IsTrue()), confidence: coalesceConfidences(l.Confidence(), r.Confidence())}
		default:
			return nil, fmt.Errorf("unknown operand: %s", le.Operand)
		}

	}
	if le.Not {
		return &expressionResult{value: !l.IsTrue(), confidence: l.Confidence()}, nil
	}
	return l, nil
}

func (le *LogicalExpression) Equal(e Expression) bool {
	if le == nil {
		return e == nil
	} else if e == nil {
		return false
	}
	ole, ok := e.(*LogicalExpression)
	if !ok {
		return false
	}
	if le.Not != ole.Not {
		return false
	}
	if !le.Left.Equal(ole.Left) {
		return false
	}
	if len(le.Right) != len(ole.Right) {
		return false
	}
	for i, re := range le.Right {
		ore := ole.Right[i]
		if !re.Equal(ore) {
			return false
		}
	}
	return true
}

func (le *LogicalExpression) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":    "logical",
		"operand": le.Operand,
		"left":    le.Left,
		"right":   le.Right,
	}
	if le.Not {
		js["not"] = true
	}
	return json.Marshal(js)
}

func (le *LogicalExpression) Clone() Expression {
	nle := &LogicalExpression{Not: le.Not, Operand: le.Operand, Left: le.Left.Clone()}
	for _, re := range le.Right {
		nle.Right = append(nle.Right, re.Clone())
	}
	return nle
}
