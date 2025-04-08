package constraint

import (
	"encoding/json"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type RangeConstraint struct {
	Minimum Limit `json:"min"`
	Maximum Limit `json:"max"`
}

func (c *RangeConstraint) Type() Type {
	return ConstraintTypeRange
}

func (c *RangeConstraint) ASCIIDocString(dataType *types.DataType) string {
	var s strings.Builder
	minRequiresParens := c.Minimum.NeedsParens(false)
	if minRequiresParens {
		s.WriteString("(")
	}
	s.WriteString(c.Minimum.ASCIIDocString(dataType))
	if minRequiresParens {
		s.WriteString(")")
	}
	s.WriteString(" to ")
	maxRequiresParens := c.Maximum.NeedsParens(false)
	if maxRequiresParens {
		s.WriteString("(")
	}
	s.WriteString(c.Maximum.ASCIIDocString(dataType))
	if maxRequiresParens {
		s.WriteString(")")
	}
	return s.String()
}

func (c *RangeConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*RangeConstraint); ok {
		return oc.Minimum.Equal(c.Minimum) && oc.Maximum.Equal(c.Maximum)
	}
	return false
}

func (c *RangeConstraint) Min(cc Context) (from types.DataTypeExtreme) {
	m := c.Minimum.Min(cc)
	return m
}

func (c *RangeConstraint) Max(cc Context) (to types.DataTypeExtreme) {
	m := c.Maximum.Max(cc)
	return m
}

func (c *RangeConstraint) Fallback(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *RangeConstraint) NeedsParens(topLevel bool) bool {
	return false
}

func (c *RangeConstraint) Clone() Constraint {
	return &RangeConstraint{Minimum: c.Minimum.Clone(), Maximum: c.Maximum.Clone()}
}

func (c *RangeConstraint) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "range",
		"min":  c.Minimum,
		"max":  c.Maximum,
	}
	return json.Marshal(js)
}

func (c *RangeConstraint) UnmarshalJSON(data []byte) (err error) {
	var js struct {
		Max json.RawMessage `json:"max"`
		Min json.RawMessage `json:"min"`
	}
	err = json.Unmarshal(data, &js)
	if err != nil {
		return
	}

	c.Minimum, err = UnmarshalLimit(js.Min)
	if err != nil {
		return
	}
	c.Maximum, err = UnmarshalLimit(js.Max)
	return
}
