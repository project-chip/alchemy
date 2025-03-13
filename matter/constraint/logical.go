package constraint

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type LogicalLimit struct {
	Operand string
	Left    Limit
	Right   []Limit
	Not     bool
}

func NewLogicalLimit(operand string, left Limit, right []any) (*LogicalLimit, error) {
	le := &LogicalLimit{Operand: operand, Left: left}
	for _, r := range right {
		rce, ok := r.(Limit)
		if !ok {
			return nil, fmt.Errorf("unexpected type in logical Limit: %T", r)
		}
		le.Right = append(le.Right, rce)
	}
	return le, nil
}

func (ll *LogicalLimit) ASCIIDocString(dataType *types.DataType) string {

	return ll.string(dataType, true, func(l Limit, dataType *types.DataType) string {
		return l.ASCIIDocString(dataType)
	})
}

func (ll *LogicalLimit) string(dataType *types.DataType, escape bool, value func(l Limit, dataType *types.DataType) string) string {
	switch ll.Operand {
	case "|", "or":
		var s strings.Builder
		s.WriteRune('(')
		if ll.Not {
			s.WriteRune('!')
			s.WriteString(value(ll.Left, dataType))
			for _, r := range ll.Right {
				s.WriteString(" & !")
				s.WriteString(value(r, dataType))
			}

		} else {
			s.WriteString(value(ll.Left, dataType))
			for _, r := range ll.Right {
				if ll.Operand == "|" {
					if escape {
						s.WriteString(" \\| ")
					} else {
						s.WriteString(" | ")
					}
				} else {
					s.WriteString(" or ")
				}
				s.WriteString(value(r, dataType))
			}
		}
		s.WriteRune(')')
		return s.String()
	case "&", "and":
		var s strings.Builder
		s.WriteRune('(')
		s.WriteString(value(ll.Left, dataType))
		for _, r := range ll.Right {
			if ll.Not {
				if ll.Operand == "and" {
					if escape {
						s.WriteString(" \\| ")
					} else {
						s.WriteString(" | ")
					}
				} else {
					s.WriteString(" or ")
				}
			} else {
				s.WriteRune(' ')
				s.WriteString(ll.Operand)
				s.WriteRune(' ')
			}
			s.WriteString(value(r, dataType))
		}
		s.WriteRune(')')
		return s.String()
	case "^":
		var s strings.Builder
		if ll.Not {
			s.WriteString("!")
		}
		s.WriteRune('(')
		s.WriteString(value(ll.Left, dataType))
		for _, r := range ll.Right {
			s.WriteString(" ^ ")
			s.WriteString(value(r, dataType))
		}
		s.WriteRune(')')
		return s.String()
	default:
		return "unknown operator"
	}
}

func (ll *LogicalLimit) DataModelString(dataType *types.DataType) string {
	return ll.string(dataType, false, func(l Limit, dataType *types.DataType) string {
		return l.DataModelString(dataType)
	})
}

func (ll *LogicalLimit) Min(cc Context) (min types.DataTypeExtreme) {
	switch ll.Operand {
	case "|", "or", "^":
		min = ll.Left.Min(cc)
		for _, r := range ll.Right {
			min = minExtreme(min, r.Min(cc))
		}
	case "&", "and":
		min = ll.Left.Min(cc)
		for _, r := range ll.Right {
			min = maxExtreme(min, r.Min(cc))
		}
	}
	return
}

func (ll *LogicalLimit) Max(cc Context) (max types.DataTypeExtreme) {
	switch ll.Operand {
	case "|", "or", "^":
		max = ll.Left.Max(cc)
		for _, r := range ll.Right {
			max = maxExtreme(max, r.Max(cc))
		}
	case "&", "and":
		max = ll.Left.Max(cc)
		for _, r := range ll.Right {
			max = minExtreme(max, r.Max(cc))
		}
	}
	return
}

func (ll *LogicalLimit) Fallback(cc Context) (max types.DataTypeExtreme) {
	return ll.Min(cc)
}

func (ll *LogicalLimit) Equal(l Limit) bool {
	if ll == nil {
		return l == nil
	} else if l == nil {
		return false
	}
	ole, ok := l.(*LogicalLimit)
	if !ok {
		return false
	}
	if ll.Not != ole.Not {
		return false
	}
	if !ll.Left.Equal(ole.Left) {
		return false
	}
	if len(ll.Right) != len(ole.Right) {
		return false
	}
	for i, re := range ll.Right {
		ore := ole.Right[i]
		if !re.Equal(ore) {
			return false
		}
	}
	return true
}

func (le *LogicalLimit) MarshalJSON() ([]byte, error) {
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

func (ll *LogicalLimit) Clone() Limit {
	nle := &LogicalLimit{Not: ll.Not, Operand: ll.Operand, Left: ll.Left.Clone()}
	for _, re := range ll.Right {
		nle.Right = append(nle.Right, re.Clone())
	}
	return nle
}

func NewLogicalConstraint(operand string, left Constraint, right []any) (*LogicalConstraint, error) {
	le := &LogicalConstraint{Operand: operand, Left: left}
	for _, r := range right {
		rce, ok := r.(Constraint)
		if !ok {
			return nil, fmt.Errorf("unexpected type in logical constraint: %T", r)
		}
		le.Right = append(le.Right, rce)
	}
	return le, nil
}

type LogicalConstraint struct {
	Operand string
	Left    Constraint
	Right   []Constraint
	Not     bool
}

func (lc *LogicalConstraint) Type() Type {
	return ConstraintTypeTagList
}

func (lc *LogicalConstraint) ASCIIDocString(dataType *types.DataType) string {
	switch lc.Operand {
	case "|", "or":
		var s strings.Builder
		s.WriteRune('(')
		if lc.Not {
			s.WriteRune('!')
			s.WriteString(lc.Left.ASCIIDocString(dataType))
			for _, r := range lc.Right {
				s.WriteString(" & !")
				s.WriteString(r.ASCIIDocString(dataType))
			}

		} else {
			s.WriteString(lc.Left.ASCIIDocString(dataType))
			for _, r := range lc.Right {
				if lc.Operand == "|" {
					s.WriteString(" | ")
				} else {
					s.WriteString(" or ")
				}
				s.WriteString(r.ASCIIDocString(dataType))
			}
		}
		s.WriteRune(')')
		return s.String()
	case "&", "and":
		var s strings.Builder
		s.WriteRune('(')
		s.WriteString(lc.Left.ASCIIDocString(dataType))
		for _, r := range lc.Right {
			if lc.Not {
				if lc.Operand == "and" {
					s.WriteString(" \\| ")
				} else {
					s.WriteString(" or ")
				}
			} else {
				s.WriteRune(' ')
				s.WriteString(lc.Operand)
				s.WriteRune(' ')
			}
			s.WriteString(r.ASCIIDocString(dataType))
		}
		s.WriteRune(')')
		return s.String()
	case "^":
		var s strings.Builder
		if lc.Not {
			s.WriteString("!")
		}
		s.WriteRune('(')
		s.WriteString(lc.Left.ASCIIDocString(dataType))
		for _, r := range lc.Right {
			s.WriteString(" ^ ")
			s.WriteString(r.ASCIIDocString(dataType))
		}
		s.WriteRune(')')
		return s.String()
	default:
		return "unknown operator"
	}
}

func (lc *LogicalConstraint) Equal(c Constraint) bool {
	if lc == nil {
		return c == nil
	} else if c == nil {
		return false
	}
	olc, ok := c.(*LogicalConstraint)
	if !ok {
		return false
	}
	if lc.Not != olc.Not {
		return false
	}
	if !lc.Left.Equal(olc.Left) {
		return false
	}
	if len(lc.Right) != len(olc.Right) {
		return false
	}
	for i, re := range lc.Right {
		ore := olc.Right[i]
		if !re.Equal(ore) {
			return false
		}
	}
	return true
}

func (lc *LogicalConstraint) Min(c Context) (min types.DataTypeExtreme) {
	switch lc.Operand {
	case "|", "or", "^":
		min = lc.Left.Min(c)
		for _, r := range lc.Right {
			min = minExtreme(min, r.Min(c))
		}
	case "&", "and":
		min = lc.Left.Min(c)
		for _, r := range lc.Right {
			min = maxExtreme(min, r.Min(c))
		}
	}
	return
}

func (lc *LogicalConstraint) Max(c Context) (max types.DataTypeExtreme) {
	switch lc.Operand {
	case "|", "or", "^":
		max = lc.Left.Max(c)
		for _, r := range lc.Right {
			max = minExtreme(max, r.Max(c))
		}
	case "&", "and":
		max = lc.Left.Max(c)
		for _, r := range lc.Right {
			max = minExtreme(max, r.Max(c))
		}
	}
	return
}

func (lc *LogicalConstraint) Clone() Constraint {
	nle := &LogicalConstraint{Not: lc.Not, Operand: lc.Operand, Left: lc.Left.Clone()}
	for _, re := range lc.Right {
		nle.Right = append(nle.Right, re.Clone())
	}
	return nle
}
