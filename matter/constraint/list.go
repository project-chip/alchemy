package constraint

import (
	"encoding/json"
	"fmt"

	"github.com/project-chip/alchemy/matter/types"
)

type ListConstraint struct {
	Constraint      Constraint
	EntryConstraint Constraint
}

func (c *ListConstraint) Type() Type {
	return ConstraintTypeList
}

func (c *ListConstraint) ASCIIDocString(dataType *types.DataType) string {
	return fmt.Sprintf("%s[%s]", c.Constraint.ASCIIDocString(dataType), c.EntryConstraint.ASCIIDocString(dataType))
}

func (c *ListConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*ListConstraint); ok {
		return oc.Constraint.Equal(c.Constraint) && oc.EntryConstraint.Equal(c.EntryConstraint)
	}
	return false
}

func (c *ListConstraint) Min(cc Context) (min types.DataTypeExtreme) {
	return c.Constraint.Min(cc)
}

func (c *ListConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Constraint.Max(cc)
}

func (c *ListConstraint) Fallback(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *ListConstraint) NeedsParens(topLevel bool) bool {
	return false
}

func (c *ListConstraint) Clone() Constraint {
	return &ListConstraint{Constraint: c.Constraint.Clone(), EntryConstraint: c.EntryConstraint.Clone()}
}

func (c *ListConstraint) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":       "list",
		"constraint": c.Constraint,
	}
	if c.EntryConstraint != nil {
		js["entryConstraint"] = c.EntryConstraint
	}
	return json.Marshal(js)
}

func (c *ListConstraint) UnmarshalJSON(data []byte) (err error) {
	var js struct {
		Constraint      json.RawMessage `json:"constraint"`
		EntryConstraint json.RawMessage `json:"entryConstraint"`
	}
	err = json.Unmarshal(data, &js)
	if err != nil {
		return
	}

	c.Constraint, err = UnmarshalConstraint(js.Constraint)
	if err != nil {
		return
	}
	if len(js.EntryConstraint) > 0 {
		c.EntryConstraint, err = UnmarshalConstraint(js.EntryConstraint)
	}
	return
}
