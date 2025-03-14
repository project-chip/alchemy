package constraint

import (
	"github.com/project-chip/alchemy/matter/types"
)

type Context interface {
	DataType() *types.DataType
	Nullable() bool
	Fallback(entity types.Entity, field Limit) types.DataTypeExtreme
	MinEntityValue(entity types.Entity, field Limit) types.DataTypeExtreme
	MaxEntityValue(entity types.Entity, field Limit) types.DataTypeExtreme
}
