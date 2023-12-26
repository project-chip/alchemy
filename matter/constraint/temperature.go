package constraint

import (
	"github.com/hasty/alchemy/matter"
	"github.com/shopspring/decimal"
)

type TemperatureLimit struct {
	Value decimal.Decimal
}

func (c *TemperatureLimit) AsciiDocString(dataType *matter.DataType) string {
	return c.Value.String() + "°C"
}

func (c *TemperatureLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*TemperatureLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *TemperatureLimit) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	var i int64
	switch cc.Field.Type.BaseType {
	case matter.BaseDataTypeTemperature, matter.BaseDataTypeTemperatureDifference:
		i = c.Value.Mul(decimal.NewFromInt(100)).IntPart()
	case matter.BaseDataTypeUnsignedTemperature, matter.BaseDataTypeSignedTemperature:
		i = c.Value.Mul(decimal.NewFromInt(10)).IntPart()
	}
	return matter.ConstraintExtreme{
		Type:   matter.ConstraintExtremeTypeInt64,
		Format: matter.NumberFormatInt,
		Int64:  i,
	}
}

func (c *TemperatureLimit) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return c.Min(cc)
}

func (c *TemperatureLimit) Default(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return c.Min(cc)
}

func (c *TemperatureLimit) Clone() matter.ConstraintLimit {
	return &TemperatureLimit{Value: c.Value.Copy()}
}
