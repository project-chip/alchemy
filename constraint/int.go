package constraint

import (
	"strconv"

	"github.com/hasty/alchemy/matter"
)

type IntLimit struct {
	Value int64
}

func (c *IntLimit) AsciiDocString() string {
	return strconv.FormatInt(c.Value, 10)
}

func (c *IntLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*IntLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *IntLimit) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return matter.ConstraintExtreme{
			Type:  matter.ConstraintExtremeTypeInt64,
			Int64: c.Value},
		matter.ConstraintExtreme{
			Type:  matter.ConstraintExtremeTypeInt64,
			Int64: c.Value,
		}
}
