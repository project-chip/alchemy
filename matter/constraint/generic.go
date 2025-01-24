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

type GenericLimit struct {
	Value string `json:"value"`
}

func (gl *GenericLimit) ASCIIDocString(dataType *types.DataType) string {
	return gl.Value
}

func (gl *GenericLimit) DataModelString(dataType *types.DataType) string {
	return gl.Value
}

func (gl *GenericLimit) Equal(o Limit) bool {
	ogl, ok := o.(*GenericLimit)
	if ok {
		return ogl.Value == gl.Value
	}
	return false
}

func (gl *GenericLimit) Min(c Context) (min types.DataTypeExtreme) {
	return
}

func (gl *GenericLimit) Max(c Context) (max types.DataTypeExtreme) {
	return
}

func (gl *GenericLimit) Fallback(c Context) (max types.DataTypeExtreme) {
	return
}

func (gl *GenericLimit) Clone() Limit {
	return &GenericLimit{Value: gl.Value}
}
