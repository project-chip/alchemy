package constraint

import (
	"math/big"
	"strconv"

	"github.com/hasty/alchemy/matter"
)

type ExpLimit struct {
	Value int64
	Exp   int64
}

func (c *ExpLimit) AsciiDocString(dataType *matter.DataType) string {
	return strconv.FormatInt(c.Value, 10) + "^" + strconv.FormatInt(c.Exp, 10) + "^"
}

func (c *ExpLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*ExpLimit); ok {
		return oc.Value == c.Value && oc.Exp == c.Exp
	}
	return false
}

func (c *ExpLimit) minmax(cc *matter.ConstraintContext) (minmax matter.ConstraintExtreme) {

	negative := c.Value < 0
	base := c.Value
	if negative {
		base *= -1
	}
	v := big.NewInt(base)
	e := big.NewInt(c.Exp)
	var t big.Int
	t.Exp(v, e, nil)
	if !t.IsInt64() {
		return
	}
	i := t.Int64()
	if negative {
		i *= -1
	}
	minmax = matter.ConstraintExtreme{
		Type:   matter.ConstraintExtremeTypeInt64,
		Format: matter.ConstraintExtremeFormatHex,
		Int64:  i,
	}
	return
}

func (c *ExpLimit) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	return c.minmax(cc)
}

func (c *ExpLimit) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return c.minmax(cc)
}

func (c *ExpLimit) Default(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return c.minmax(cc)
}
