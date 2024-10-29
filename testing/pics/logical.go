package pics

import (
	"fmt"
	"strings"
)

type LogicalOperator uint8

const (
	LogicalOperatorNone LogicalOperator = iota
	LogicalOperatorAnd
	LogicalOperatorOr
)

func (co LogicalOperator) String() string {
	switch co {
	case LogicalOperatorAnd:
		return "&&"
	case LogicalOperatorOr:
		return "||"
	default:
		return ""
	}
}

func (co LogicalOperator) PythonString() string {
	switch co {
	case LogicalOperatorAnd:
		return "and"
	case LogicalOperatorOr:
		return "or"
	default:
		return ""
	}
}

type LogicalExpression struct {
	Operand LogicalOperator
	Left    Expression
	Right   []Expression
	Not     bool
}

func NewLogicalExpression(operator LogicalOperator, left Expression, right []any) (*LogicalExpression, error) {
	le := &LogicalExpression{Operand: operator, Left: left}
	for _, r := range right {
		rce, ok := r.(Expression)
		if !ok {
			return nil, fmt.Errorf("unexpected type in logical expression: %T", r)
		}
		le.Right = append(le.Right, rce)
	}
	return le, nil
}

func (pe LogicalExpression) String() string {
	var sb strings.Builder
	if pe.Not {
		sb.WriteString("!")
	}
	sb.WriteRune('(')
	sb.WriteString(pe.Left.String())
	sb.WriteRune(' ')
	for _, re := range pe.Right {
		sb.WriteString(pe.Operand.String())
		sb.WriteRune(' ')
		sb.WriteString(re.String())
	}
	sb.WriteRune(')')
	return sb.String()
}

func (pe LogicalExpression) PythonString() string {
	var sb strings.Builder
	if pe.Not {
		sb.WriteString("not ")
	}
	sb.WriteString(pe.Left.PythonString())
	for _, re := range pe.Right {
		sb.WriteRune(' ')
		sb.WriteString(pe.Operand.PythonString())
		sb.WriteRune(' ')
		sb.WriteString(re.PythonString())
	}
	return sb.String()
}

func (pe LogicalExpression) PythonBuilder(aliases map[string]string, sb *strings.Builder) {
	if pe.Not {
		sb.WriteString("not ")
	}
	pe.Left.PythonBuilder(aliases, sb)
	for _, re := range pe.Right {
		sb.WriteRune(' ')
		sb.WriteString(pe.Operand.PythonString())
		sb.WriteRune(' ')
		re.PythonBuilder(aliases, sb)
	}
}
