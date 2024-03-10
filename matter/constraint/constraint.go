package constraint

import (
	"encoding/json"
	"fmt"

	"github.com/hasty/alchemy/matter/types"
)

func ParseString(constraint string) (Constraint, error) {
	if len(constraint) == 0 {
		return &GenericConstraint{Value: constraint}, nil
	}
	c, err := Parse("", []byte(constraint))
	if err != nil {
		return nil, err
	}
	return c.(Constraint), nil
}

type ConstraintType uint8

const (
	ConstraintTypeUnknown ConstraintType = iota
	ConstraintTypeAll                    // Special section type for everything that comes before any known sections
	ConstraintTypeDescribed
	ConstraintTypeExact
	ConstraintTypeGeneric
	ConstraintTypeList
	ConstraintTypeMax
	ConstraintTypeMin
	ConstraintTypeRange
	ConstraintTypeSet
)

var nameConstraintType map[string]ConstraintType

var constraintTypeNames = map[ConstraintType]string{
	ConstraintTypeUnknown:   "unknown",
	ConstraintTypeAll:       "all",
	ConstraintTypeDescribed: "described",
	ConstraintTypeExact:     "exact",
	ConstraintTypeGeneric:   "generic",
	ConstraintTypeList:      "list",
	ConstraintTypeMax:       "max",
	ConstraintTypeMin:       "min",
	ConstraintTypeRange:     "range",
	ConstraintTypeSet:       "set",
}

func init() {
	nameConstraintType = make(map[string]ConstraintType, len(constraintTypeNames))
	for i, q := range constraintTypeNames {
		nameConstraintType[q] = i
	}
}

func (ct ConstraintType) MarshalJSON() ([]byte, error) {
	v, ok := constraintTypeNames[ct]
	if !ok {
		return nil, fmt.Errorf("unknown constraint type %d", ct)
	}
	return json.Marshal(v)
}

func (c *ConstraintType) UnmarshalJSON(bytes []byte) error {
	var name string
	err := json.Unmarshal(bytes, &name)
	if err != nil {
		return err
	}
	v, ok := nameConstraintType[name]
	if !ok {
		return fmt.Errorf("unknown constraint type %s", name)
	}
	*c = v
	return nil
}

type Constraint interface {
	Type() ConstraintType
	AsciiDocString(dataType *types.DataType) string
	Equal(o Constraint) bool
	Min(c Context) (min types.DataTypeExtreme)
	Max(c Context) (max types.DataTypeExtreme)
	Default(c Context) (max types.DataTypeExtreme)
	Clone() Constraint
}

type ConstraintLimit interface {
	AsciiDocString(dataType *types.DataType) string
	DataModelString(dataType *types.DataType) string
	Equal(o ConstraintLimit) bool
	Min(c Context) (min types.DataTypeExtreme)
	Max(c Context) (max types.DataTypeExtreme)
	Default(c Context) (max types.DataTypeExtreme)
	Clone() ConstraintLimit
}
