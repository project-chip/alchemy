package constraint

import (
	"encoding/json"
	"fmt"

	"github.com/project-chip/alchemy/matter/types"
)

type MaxConstraint struct {
	Maximum Limit
}

func (c *MaxConstraint) Type() Type {
	return ConstraintTypeMax
}

func (c *MaxConstraint) ASCIIDocString(dataType *types.DataType) string {
	return fmt.Sprintf("max %s", c.Maximum.ASCIIDocString(dataType))
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
		min = types.Min(dt.BaseType, cc.Nullable())
	}
	return
}

func (c *MaxConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Maximum.Max(cc)
}

func (c *MaxConstraint) Fallback(cc Context) (max types.DataTypeExtreme) {
	return
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
