package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter/types"
)

type LengthLimit struct {
	Value string
}

func (ll *LengthLimit) AsciiDocString(dataType *types.DataType) string {
	return fmt.Sprintf("len(%s)", ll.Value)
}

func (c *LengthLimit) DataModelString(dataType *types.DataType) string {
	return c.Value
}

func (ll *LengthLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*LengthLimit); ok {
		return oc.Value == ll.Value
	}
	return false
}

func (c *LengthLimit) Min(cc Context) (min types.DataTypeExtreme) {
	rc := cc.ReferenceConstraint(c.Value)
	if rc == nil {
		return
	}
	return rc.Min(cc)
}

func (c *LengthLimit) Max(cc Context) (max types.DataTypeExtreme) {
	rc := cc.ReferenceConstraint(c.Value)
	if rc == nil {
		return
	}
	return rc.Max(cc)
}

func (c *LengthLimit) Default(cc Context) (def types.DataTypeExtreme) {
	return cc.Default(c.Value)
}

func (c *LengthLimit) Clone() ConstraintLimit {
	return &LengthLimit{Value: c.Value}
}
