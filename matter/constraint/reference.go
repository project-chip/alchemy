package constraint

import (
	"github.com/hasty/alchemy/matter/types"
)

type ReferenceLimit struct {
	Value string
}

func (c *ReferenceLimit) AsciiDocString(dataType *types.DataType) string {
	return c.Value
}

func (c *ReferenceLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*ReferenceLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *ReferenceLimit) Min(cc Context) (min types.DataTypeExtreme) {
	rc := cc.ReferenceConstraint(c.Value)
	if rc == nil {
		return
	}
	return rc.Min(cc)
}

func (c *ReferenceLimit) Max(cc Context) (max types.DataTypeExtreme) {
	rc := cc.ReferenceConstraint(c.Value)
	if rc == nil {
		return
	}
	return rc.Max(cc)
}

func (c *ReferenceLimit) Default(cc Context) (def types.DataTypeExtreme) {
	return cc.Default(c.Value)
}

func (c *ReferenceLimit) Clone() ConstraintLimit {
	return &ReferenceLimit{Value: c.Value}
}
