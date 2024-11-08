package constraint

import (
	"encoding/json"
	"fmt"

	"github.com/project-chip/alchemy/matter/types"
)

type MinConstraint struct {
	Minimum Limit `json:"min"`
}

func (c *MinConstraint) Type() Type {
	return ConstraintTypeMin
}

func (c *MinConstraint) ASCIIDocString(dataType *types.DataType) string {
	return fmt.Sprintf("min %s", c.Minimum.ASCIIDocString(dataType))
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
	return
}

func (c *MinConstraint) Fallback(cc Context) (max types.DataTypeExtreme) {
	return
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
