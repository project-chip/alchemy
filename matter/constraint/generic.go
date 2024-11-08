package constraint

import (
	"encoding/json"

	"github.com/project-chip/alchemy/matter/types"
)

type GenericConstraint struct {
	Value string `json:"value"`
}

func (c *GenericConstraint) Type() Type {
	return ConstraintTypeGeneric
}

func (c *GenericConstraint) ASCIIDocString(dataType *types.DataType) string {
	return c.Value
}

func (c *GenericConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*GenericConstraint); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *GenericConstraint) Min(cc Context) (min types.DataTypeExtreme) {
	return
}

func (c *GenericConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *GenericConstraint) Fallback(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *GenericConstraint) Clone() Constraint {
	return &GenericConstraint{Value: c.Value}
}

func (c *GenericConstraint) MarshalJSON() ([]byte, error) {
	js := struct {
		constraintJSONBase
		GenericConstraint
	}{
		constraintJSONBase{Type: "generic"},
		*c,
	}
	return json.Marshal(js)
}
