package constraint

import (
	"encoding/json"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type MaxConstraint struct {
	Maximum Limit
}

func (c *MaxConstraint) Type() Type {
	return ConstraintTypeMax
}

func (c *MaxConstraint) ASCIIDocString(dataType *types.DataType) string {
	requiresParens := c.Maximum.NeedsParens(false)
	var s strings.Builder
	s.WriteString("max ")
	if requiresParens {
		s.WriteString("(")
	}
	s.WriteString(c.Maximum.ASCIIDocString(dataType))
	if requiresParens {
		s.WriteString(")")
	}
	return s.String()
}

func (c *MaxConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*MaxConstraint); ok {
		return oc.Maximum.Equal(c.Maximum)
	}
	return false
}

func (c *MaxConstraint) Min(cc Context) (min types.DataTypeExtreme) {
	dt := cc.DataType()
	if dt != nil {
		min = types.Min(dt.BaseType, cc.Nullability())
	}
	return
}

func (c *MaxConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Maximum.Max(cc)
}

func (c *MaxConstraint) Fallback(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *MaxConstraint) NeedsParens(topLevel bool) bool {
	return false
}

func (c *MaxConstraint) Clone() Constraint {
	return &MaxConstraint{Maximum: c.Maximum.Clone()}
}

func (c *MaxConstraint) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "max",
		"max":  c.Maximum,
	}
	return json.Marshal(js)
}

func (c *MaxConstraint) UnmarshalJSON(data []byte) (err error) {
	var js struct {
		Max json.RawMessage `json:"max"`
	}
	err = json.Unmarshal(data, &js)
	if err != nil {
		return
	}

	c.Maximum, err = UnmarshalLimit(js.Max)
	return
}

type MaxOfLimit struct {
	Maximums LimitSet `json:"maximums"`
}

func (c *MaxOfLimit) ASCIIDocString(dataType *types.DataType) string {

	var b strings.Builder
	b.WriteString("maxOf(")
	c.Maximums.ASCIIDocString(dataType, &b)
	b.WriteString(")")
	return b.String()
}

func (c *MaxOfLimit) DataModelString(dataType *types.DataType) string {

	var b strings.Builder
	b.WriteString("maxOf(")
	c.Maximums.DataModelString(dataType, &b)
	b.WriteString(")")
	return b.String()
}

func (c *MaxOfLimit) Equal(o Limit) bool {
	if oc, ok := o.(*MinOfLimit); ok {
		return oc.Minimums.Equal(c.Maximums)
	}
	return false
}

func (c *MaxOfLimit) Min(cc Context) (min types.DataTypeExtreme) {
	var to types.DataTypeExtreme

	for _, l := range c.Maximums {
		t := l.Min(cc)
		if !to.Defined() {
			to = t
			continue
		}
		if !t.Defined() {
			continue
		}
		to = maxExtreme(to, t)
	}
	return to
}

func (c *MaxOfLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Maximums.Max(cc)
}

func (c *MaxOfLimit) Fallback(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *MaxOfLimit) NeedsParens(topLevel bool) bool {
	return false
}

func (c *MaxOfLimit) Clone() Limit {
	return &MaxOfLimit{Maximums: c.Maximums.Clone()}
}

func (c *MaxOfLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":     "maxOf",
		"maximums": c.Maximums,
	}
	return json.Marshal(js)
}

func (c *MaxOfLimit) UnmarshalJSON(data []byte) (err error) {
	var js struct {
		Mins json.RawMessage `json:"maximums"`
	}
	err = json.Unmarshal(data, &js)
	if err != nil {
		return
	}

	err = c.Maximums.UnmarshalJSON(js.Mins)
	return
}
