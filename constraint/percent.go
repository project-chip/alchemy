package constraint

import (
	"github.com/hasty/alchemy/matter"
	"github.com/shopspring/decimal"
)

type PercentLimit struct {
	Value      decimal.Decimal
	Hundredths bool
}

func (c *PercentLimit) AsciiDocString(dataType *matter.DataType) string {
	return c.Value.String() + "%"
}

func (c *PercentLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*PercentLimit); ok {
		return oc.Value == c.Value && oc.Hundredths == c.Hundredths
	}
	return false
}

func (c *PercentLimit) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	val := c.Value
	if c.Hundredths {
		val = val.Mul(decimal.NewFromInt(100))
	}
	v := val.IntPart()
	return matter.ConstraintExtreme{
			Type:   matter.ConstraintExtremeTypeInt64,
			Format: matter.ConstraintExtremeFormatInt,
			Int64:  v},
		matter.ConstraintExtreme{
			Type:   matter.ConstraintExtremeTypeInt64,
			Format: matter.ConstraintExtremeFormatInt,
			Int64:  val.IntPart(),
		}
}
