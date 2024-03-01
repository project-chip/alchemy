package constraint

import (
	"github.com/hasty/alchemy/matter/types"
	"github.com/shopspring/decimal"
)

type TemperatureLimit struct {
	Value decimal.Decimal
}

func (c *TemperatureLimit) AsciiDocString(dataType *types.DataType) string {
	return c.Value.String() + "°C"
}

func (c *TemperatureLimit) DataModelString(dataType *types.DataType) string {
	return c.Value.String()
}

func (c *TemperatureLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*TemperatureLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *TemperatureLimit) Min(cc Context) (min types.DataTypeExtreme) {
	var i int64
	bt := cc.DataType()
	if bt == nil {
		return
	}
	switch bt.BaseType {
	case types.BaseDataTypeTemperature, types.BaseDataTypeTemperatureDifference:
		i = c.Value.Mul(decimal.NewFromInt(100)).IntPart()
	case types.BaseDataTypeUnsignedTemperature, types.BaseDataTypeSignedTemperature:
		i = c.Value.Mul(decimal.NewFromInt(10)).IntPart()
	}
	return types.DataTypeExtreme{
		Type:   types.DataTypeExtremeTypeInt64,
		Format: types.NumberFormatInt,
		Int64:  i,
	}
}

func (c *TemperatureLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *TemperatureLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *TemperatureLimit) Clone() ConstraintLimit {
	return &TemperatureLimit{Value: c.Value.Copy()}
}
