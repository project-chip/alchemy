package constraint

import (
	"encoding/json"
	"fmt"

	"github.com/project-chip/alchemy/matter/types"
	"github.com/shopspring/decimal"
)

type TemperatureLimit struct {
	Value decimal.Decimal `json:"value"`
}

func (c *TemperatureLimit) ASCIIDocString(dataType *types.DataType) string {
	return c.Value.String() + "°C"
}

func (c *TemperatureLimit) DataModelString(dataType *types.DataType) string {
	return fmt.Sprintf("%d", c.limit(dataType).Int64)
}

func (c *TemperatureLimit) Equal(o Limit) bool {
	if oc, ok := o.(*TemperatureLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *TemperatureLimit) limit(dataType *types.DataType) types.DataTypeExtreme {
	var i int64
	switch dataType.BaseType {
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

func (c *TemperatureLimit) Min(cc Context) (min types.DataTypeExtreme) {
	dt := cc.DataType()
	if dt == nil {
		return
	}
	return c.limit(dt)
}

func (c *TemperatureLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *TemperatureLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *TemperatureLimit) Clone() Limit {
	return &TemperatureLimit{Value: c.Value.Copy()}
}

func (c *TemperatureLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "temperature",
		"value": c.Value,
	}
	return json.Marshal(js)
}
