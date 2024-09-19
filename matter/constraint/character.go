package constraint

import (
	"encoding/json"
	"fmt"

	"github.com/project-chip/alchemy/matter/types"
)

type CharacterLimit struct {
	ByteCount      Limit `json:"byteCount"`
	CodepointCount Limit `json:"codepointCount"`
}

func (c *CharacterLimit) ASCIIDocString(dataType *types.DataType) string {
	return fmt.Sprintf("%s{%s}", c.ByteCount.ASCIIDocString(dataType), c.CodepointCount.ASCIIDocString(dataType))
}

func (c *CharacterLimit) DataModelString(dataType *types.DataType) string {
	return c.ByteCount.DataModelString(dataType)
}

func (c *CharacterLimit) Equal(o Limit) bool {
	if oc, ok := o.(*CharacterLimit); ok {
		return oc.ByteCount.Equal(c.ByteCount) && oc.CodepointCount.Equal(c.CodepointCount)
	}
	return false
}

func (c *CharacterLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return c.ByteCount.Min(cc)
}

func (c *CharacterLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.ByteCount.Max(cc)
}

func (c *CharacterLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.ByteCount.Default(cc)
}

func (c *CharacterLimit) Clone() Limit {
	return &CharacterLimit{ByteCount: c.ByteCount.Clone(), CodepointCount: c.CodepointCount.Clone()}
}

func (c *CharacterLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":       "character",
		"bytes":      c.ByteCount,
		"codepoints": c.CodepointCount,
	}
	return json.Marshal(js)
}
