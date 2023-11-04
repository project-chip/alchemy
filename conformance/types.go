package conformance

import (
	"fmt"
	"strings"
)

type Conformance interface {
	fmt.Stringer
}

type OtherwiseConformance struct {
	Conformances []Conformance
}

func (oc *OtherwiseConformance) String() string {
	var s strings.Builder
	for _, c := range oc.Conformances {
		if s.Len() > 0 {
			s.WriteString(", otherwise ")
		}
		s.WriteString(c.String())
	}
	return s.String()
}

type ChoiceIdentifier struct {
	Conformance Conformance
	Choice      *Choice
}

func (ci *ChoiceIdentifier) String() string {
	var s strings.Builder
	s.WriteString("optional if ")
	s.WriteString(ci.Conformance.String())
	if ci.Choice != nil {
		s.WriteString("(")
		s.WriteString(ci.Choice.String())
		s.WriteString(")")
	}
	return s.String()
}

type IdentifierConformance struct {
	ID  string
	Not bool
}

func (id *IdentifierConformance) String() string {
	if id.Not {
		return fmt.Sprintf("not %s", id.ID)
	}
	return id.ID
}

func (id *IdentifierConformance) Eval(context map[string]any) (bool, error) {
	v, ok := context[id.ID]
	if !ok {
		return id.Not, nil
	}
	if b, ok := v.(bool); ok {
		return b != id.Not, nil
	}
	return false, fmt.Errorf("unexpected value when interpreting identifier %s: %v", id.ID, v)
}

type Optional struct {
	Choice *Choice
}

func (o *Optional) String() string {
	if o.Choice == nil {
		return "Optional"
	}
	return fmt.Sprintf("Optional; %s", o.Choice.String())
}

func (id *Optional) Eval(context map[string]any) (bool, error) {
	return true, nil
}

type Mandatory struct {
}

type Deprecated struct {
}

type Provisional struct {
}

type Otherwise struct {
}

type Disallowed struct {
}

type EqualityConformance struct {
	Not   bool
	Left  Conformance
	Right Conformance
}

func (o *EqualityConformance) String() string {
	if o.Not {
		return fmt.Sprintf("(%s != %s)", o.Left, o.Right)
	}

	return fmt.Sprintf("(%s == %s)", o.Left, o.Right)
}

type LogicalConformance struct {
	Operand string
	Left    Conformance
	Right   Conformance
}

func (lc *LogicalConformance) String() string {
	switch lc.Operand {
	case "|":
		return fmt.Sprintf("(%s or %s)", lc.Left.String(), lc.Right.String())
	case "&":
		return fmt.Sprintf("(%s and %s)", lc.Left.String(), lc.Right.String())
	case "^":
		return fmt.Sprintf("(%s xor %s)", lc.Left.String(), lc.Right.String())
	default:
		return "unknown operator"
	}
}

type Choice struct {
	Set   string
	Limit ChoiceLimit
}

func (c *Choice) String() string {
	if c.Limit != nil {
		return c.Limit.String(c.Set)
	}
	return fmt.Sprintf("set: %s", c.Set)
}

type ChoiceLimit interface {
	String(set string) string
}

type ChoiceExactLimit struct {
	Limit int
}

func (c *ChoiceExactLimit) String(set string) string {
	return fmt.Sprintf("exactly %d of set %s", c.Limit, set)
}

type ChoiceMinLimit struct {
	Min int
}

func (c *ChoiceMinLimit) String(set string) string {
	return fmt.Sprintf("at least %d of set %s", c.Min, set)
}

type ChoiceMaxLimit struct {
	Max int
}

func (c *ChoiceMaxLimit) String(set string) string {
	return fmt.Sprintf("at most %d of set %s", c.Max, set)
}

type ChoiceRangeLimit struct {
	Min int
	Max int
}

func (c *ChoiceRangeLimit) String(set string) string {
	return fmt.Sprintf("between %d and %d of set %s", c.Min, c.Max, set)
}
