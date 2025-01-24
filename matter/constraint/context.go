package constraint

import (
	"github.com/project-chip/alchemy/matter/types"
)

type Context interface {
	DataType() *types.DataType
	Fallback(entity types.Entity, field Limit) types.DataTypeExtreme
	IdentifierConstraint(entity types.Entity, field Limit) Constraint
	ReferenceConstraint(entity types.Entity, field Limit) Constraint
}
