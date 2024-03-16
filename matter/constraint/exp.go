package constraint

import (
	"encoding/json"
	"math/big"
	"strconv"

	"github.com/hasty/alchemy/matter/types"
)

type ExpLimit struct {
	Value int64
	Exp   int64
}

func (c *ExpLimit) AsciiDocString(dataType *types.DataType) string {
	return strconv.FormatInt(c.Value, 10) + "^" + strconv.FormatInt(c.Exp, 10) + "^"
}

func (c *ExpLimit) DataModelString(dataType *types.DataType) string {
	e := c.minmax()
	return e.DataModelString(dataType)
}

func (c *ExpLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*ExpLimit); ok {
		return oc.Value == c.Value && oc.Exp == c.Exp
	}
	return false
}

func (c *ExpLimit) minmax() (minmax types.DataTypeExtreme) {

	negative := c.Value < 0
	base := c.Value
	if negative {
		base *= -1
	}
	v := big.NewInt(base)
	e := big.NewInt(c.Exp)
	var t big.Int
	t.Exp(v, e, nil)
	if t.IsInt64() {
		i := t.Int64()
		if negative {
			i *= -1
		}
		minmax = types.DataTypeExtreme{
			Type:   types.DataTypeExtremeTypeInt64,
			Format: types.NumberFormatInt,
			Int64:  i,
		}
		return
	}
	if t.IsUint64() {
		u := t.Uint64()
		minmax = types.DataTypeExtreme{
			Type:   types.DataTypeExtremeTypeInt64,
			Format: types.NumberFormatHex,
			UInt64: u,
		}
		return
	}
	return
}

func (c *ExpLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return c.minmax()
}

func (c *ExpLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.minmax()
}

func (c *ExpLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.minmax()
}

func (c *ExpLimit) Clone() ConstraintLimit {
	return &ExpLimit{Value: c.Value, Exp: c.Exp}
}

func (c *ExpLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "exponent",
		"value": c.Value,
		"exp":   c.Exp,
	}
	return json.Marshal(js)
}
