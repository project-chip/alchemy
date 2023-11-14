package constraint

import (
	"math/big"

	"github.com/hasty/alchemy/matter"
)

type TemperatureLimit struct {
	Value *big.Float
}

func (c *TemperatureLimit) AsciiDocString() string {
	return c.Value.String() + "°C"
}

func (c *TemperatureLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*TemperatureLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *TemperatureLimit) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	val := big.NewFloat(0)
	val.Mul(c.Value, big.NewFloat(100))
	i, _ := val.Int64()
	return matter.ConstraintExtreme{
			Type:  matter.ConstraintExtremeTypeInt64,
			Int64: i,
		},
		matter.ConstraintExtreme{
			Type:  matter.ConstraintExtremeTypeInt64,
			Int64: i,
		}
}
