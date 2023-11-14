package constraint

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/hasty/alchemy/matter"
)

type PercentLimit struct {
	Value      *big.Float
	Hundredths bool
}

func (c *PercentLimit) AsciiDocString() string {
	if c.Hundredths {
		if c.Value.IsInt() {
			i, _ := c.Value.Uint64()
			return strconv.FormatUint(i/100, 10)
		}
		return fmt.Sprintf("%.2f", c.Value)
	}
	i, _ := c.Value.Uint64()
	return strconv.FormatUint(i, 10)
}

func (c *PercentLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*PercentLimit); ok {
		return oc.Value == c.Value && oc.Hundredths == c.Hundredths
	}
	return false
}

func (c *PercentLimit) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	i, _ := c.Value.Int64()
	return matter.ConstraintExtreme{
			Type:  matter.ConstraintExtremeTypeInt64,
			Int64: i},
		matter.ConstraintExtreme{
			Type:  matter.ConstraintExtremeTypeInt64,
			Int64: i,
		}
}
