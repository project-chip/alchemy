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

	return ll.string(dataType, true)
}

func (ll *LogicalLimit) string(dataType *types.DataType, escape bool) string {
	switch ll.Operand {
	case "|", "or":
		var s strings.Builder
		s.WriteRune('(')
		if ll.Not {
			s.WriteRune('!')
			s.WriteString(ll.Left.ASCIIDocString(dataType))
			for _, r := range ll.Right {
				s.WriteString(" & !")
				s.WriteString(r.ASCIIDocString(dataType))
			}

		} else {
			s.WriteString(ll.Left.ASCIIDocString(dataType))
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
				s.WriteString(r.ASCIIDocString(dataType))
			}
		}
		s.WriteRune(')')
		return s.String()
	case "&", "and":
		var s strings.Builder
		s.WriteRune('(')
		s.WriteString(ll.Left.ASCIIDocString(dataType))
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
			s.WriteString(r.ASCIIDocString(dataType))
		}
		s.WriteRune(')')
		return s.String()
	case "^":
		var s strings.Builder
		if ll.Not {
			s.WriteString("!")
		}
		s.WriteRune('(')
		s.WriteString(ll.Left.ASCIIDocString(dataType))
		for _, r := range ll.Right {
			s.WriteString(" ^ ")
			s.WriteString(r.ASCIIDocString(dataType))
		}
		s.WriteRune(')')
		return s.String()
	default:
		return "unknown operator"
	}
}

func (ll *LogicalLimit) DataModelString(dataType *types.DataType) string {
	return ll.string(dataType, false)
}

func (ll *LogicalLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined, Format: types.NumberFormatAuto}
}

func (ll *LogicalLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return ll.Min(cc)
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
