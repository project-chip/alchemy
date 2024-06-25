package constraint

import (
	"encoding/json"

	"github.com/project-chip/alchemy/matter/types"
	"github.com/shopspring/decimal"
)

type PercentLimit struct {
	Value      decimal.Decimal `json:"value"`
	Hundredths bool            `json:"hundredths,omitempty"`
}

func (c *PercentLimit) ASCIIDocString(dataType *types.DataType) string {
	return c.Value.String() + "%"
}

func (c *PercentLimit) DataModelString(dataType *types.DataType) string {
	return c.Value.String()
}

func (c *PercentLimit) Equal(o Limit) bool {
	if oc, ok := o.(*PercentLimit); ok {
		return oc.Value == c.Value && oc.Hundredths == c.Hundredths
	}
	return false
}

func (c *PercentLimit) Min(cc Context) (min types.DataTypeExtreme) {
	val := c.Value
	if c.Hundredths {
		val = val.Mul(decimal.NewFromInt(100))
	}
	v := val.IntPart()
	return types.DataTypeExtreme{
		Type:   types.DataTypeExtremeTypeInt64,
		Format: types.NumberFormatInt,
		Int64:  v}
}

func (c *PercentLimit) Max(cc Context) (max types.DataTypeExtreme) {
	val := c.Value
	if c.Hundredths {
		val = val.Mul(decimal.NewFromInt(100))
	}
	v := val.IntPart()
	return types.DataTypeExtreme{
		Type:   types.DataTypeExtremeTypeInt64,
		Format: types.NumberFormatInt,
		Int64:  v}

}

func (c *PercentLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *PercentLimit) Clone() Limit {
	return &PercentLimit{Value: c.Value.Copy(), Hundredths: c.Hundredths}
}

func (c *PercentLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":      "percent",
		"value":     c.Value,
		"hundreths": c.Hundredths,
	}
	return json.Marshal(js)
}
