package constraint

import (
	"encoding/json"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type MinConstraint struct {
	Minimum Limit `json:"min"`
}

func (c *MinConstraint) Type() Type {
	return ConstraintTypeMin
}

func (c *MinConstraint) ASCIIDocString(dataType *types.DataType) string {
	requiresParens := c.Minimum.NeedsParens(false)
	var s strings.Builder
	s.WriteString("min ")
	if requiresParens {
		s.WriteString("(")
	}
	s.WriteString(c.Minimum.ASCIIDocString(dataType))
	if requiresParens {
		s.WriteString(")")
	}
	return s.String()
}

func (c *MinConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*MinConstraint); ok {
		return oc.Minimum.Equal(c.Minimum)
	}
	return false
}

func (c *MinConstraint) Min(cc Context) (min types.DataTypeExtreme) {
	return c.Minimum.Min(cc)
}

func (c *MinConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	dt := cc.DataType()
	if dt != nil {
		max = types.Max(dt.BaseType, cc.Nullability())
	}
	return
}

func (c *MinConstraint) Fallback(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *MinConstraint) NeedsParens(topLevel bool) bool {
	return false
}

func (c *MinConstraint) Clone() Constraint {
	return &MinConstraint{Minimum: c.Minimum.Clone()}
}

func (c *MinConstraint) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "min",
		"min":  c.Minimum,
	}
	return json.Marshal(js)
}

func (c *MinConstraint) UnmarshalJSON(data []byte) (err error) {
	var js struct {
		Min json.RawMessage `json:"min"`
	}
	err = json.Unmarshal(data, &js)
	if err != nil {
		return
	}

	c.Minimum, err = UnmarshalLimit(js.Min)
	return
}

type MinOfLimit struct {
	Minimums LimitSet `json:"minimums"`
}

func (c *MinOfLimit) ASCIIDocString(dataType *types.DataType) string {

	var b strings.Builder
	b.WriteString("minOf(")
	c.Minimums.ASCIIDocString(dataType, &b)
	b.WriteString(")")
	return b.String()
}

func (c *MinOfLimit) DataModelString(dataType *types.DataType) string {

	var b strings.Builder
	b.WriteString("minOf(")
	c.Minimums.DataModelString(dataType, &b)
	b.WriteString(")")
	return b.String()
}

func (c *MinOfLimit) Equal(o Limit) bool {
	if oc, ok := o.(*MinOfLimit); ok {
		return oc.Minimums.Equal(c.Minimums)
	}
	return false
}

func (c *MinOfLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return c.Minimums.Min(cc)
}

func (c *MinOfLimit) Max(cc Context) (max types.DataTypeExtreme) {
	var to types.DataTypeExtreme

	for _, l := range c.Minimums {
		t := l.Max(cc)
		if !to.Defined() {
			to = t
			continue
		}
		if !t.Defined() {
			continue
		}
		to = minExtreme(to, t)
	}
	return to
}

func (c *MinOfLimit) Fallback(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *MinOfLimit) NeedsParens(topLevel bool) bool {
	return false
}

func (c *MinOfLimit) Clone() Limit {
	return &MinOfLimit{Minimums: c.Minimums.Clone()}
}

func (c *MinOfLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":     "minOf",
		"minimums": c.Minimums,
	}
	return json.Marshal(js)
}

func (c *MinOfLimit) UnmarshalJSON(data []byte) (err error) {
	var js struct {
		Mins json.RawMessage `json:"minimums"`
	}
	err = json.Unmarshal(data, &js)
	if err != nil {
		return
	}

	err = c.Minimums.UnmarshalJSON(js.Mins)
	return
}
