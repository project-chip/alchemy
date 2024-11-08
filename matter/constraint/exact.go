package constraint

import (
	"encoding/json"

	"github.com/project-chip/alchemy/matter/types"
)

type ExactConstraint struct {
	Value Limit `json:"value"`
}

func (c *ExactConstraint) Type() Type {
	return ConstraintTypeExact
}

func (c *ExactConstraint) ASCIIDocString(dataType *types.DataType) string {
	return c.Value.ASCIIDocString(dataType)
}

func (c *ExactConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*ExactConstraint); ok {
		return oc.Value.Equal(c.Value)
	}
	return false
}

func (c *ExactConstraint) Min(cc Context) (min types.DataTypeExtreme) {
	return c.Value.Min(cc)
}

func (c *ExactConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Value.Max(cc)
}

func (c *ExactConstraint) Fallback(cc Context) (max types.DataTypeExtreme) {
	return c.Value.Fallback(cc)
}

func (c *ExactConstraint) Clone() Constraint {
	return &ExactConstraint{Value: c.Value.Clone()}
}

func (c *ExactConstraint) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "exact",
		"value": c.Value,
	}
	return json.Marshal(js)
}

func (c *ExactConstraint) UnmarshalJSON(data []byte) (err error) {
	var js struct {
		Value json.RawMessage `json:"value"`
	}
	err = json.Unmarshal(data, &js)
	if err != nil {
		return
	}

	c.Value, err = UnmarshalLimit(js.Value)
	return
}
