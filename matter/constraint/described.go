package constraint

import (
	"encoding/json"

	"github.com/project-chip/alchemy/matter/types"
)

type DescribedConstraint struct {
}

func (c *DescribedConstraint) Type() Type {
	return ConstraintTypeDescribed
}

func (c *DescribedConstraint) ASCIIDocString(dataType *types.DataType) string {
	return "desc"
}

func (c *DescribedConstraint) Equal(o Constraint) bool {
	_, ok := o.(*DescribedConstraint)
	return ok
}

func (c *DescribedConstraint) Min(cc Context) (min types.DataTypeExtreme) {
	return
}

func (c *DescribedConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *DescribedConstraint) Fallback(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *DescribedConstraint) NeedsParens(topLevel bool) bool {
	return false
}

func (c *DescribedConstraint) Clone() Constraint {
	return &DescribedConstraint{}
}

func (c *DescribedConstraint) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "described",
	}
	return json.Marshal(js)
}

type DescribedLimit struct {
}

func (gl *DescribedLimit) ASCIIDocString(dataType *types.DataType) string {
	return "desc"
}

func (gl *DescribedLimit) DataModelString(dataType *types.DataType) string {
	return "desc"
}

func (gl *DescribedLimit) Equal(o Limit) bool {
	_, ok := o.(*DescribedLimit)
	return ok
}

func (gl *DescribedLimit) Min(c Context) (min types.DataTypeExtreme) {
	return
}

func (gl *DescribedLimit) Max(c Context) (max types.DataTypeExtreme) {
	return
}

func (gl *DescribedLimit) Fallback(c Context) (max types.DataTypeExtreme) {
	return
}

func (gl *DescribedLimit) NeedsParens(topLevel bool) bool {
	return false
}

func (gl *DescribedLimit) Clone() Limit {
	return &DescribedLimit{}
}
