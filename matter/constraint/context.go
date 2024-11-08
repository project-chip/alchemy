package constraint

import (
	"github.com/project-chip/alchemy/matter/types"
)

type Context interface {
	DataType() *types.DataType
	Fallback(name string) types.DataTypeExtreme
	ReferenceConstraint(ref string) Constraint
}
