package constraint

import (
	"github.com/hasty/alchemy/matter/types"
)

type Context interface {
	DataType() *types.DataType
	Default(name string) types.DataTypeExtreme
	ReferenceConstraint(ref string) Constraint
}
