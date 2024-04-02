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

type Type uint8

const (
	ConstraintTypeUnknown Type = iota
	ConstraintTypeAll          // Special section type for everything that comes before any known sections
	ConstraintTypeDescribed
	ConstraintTypeExact
	ConstraintTypeGeneric
	ConstraintTypeList
	ConstraintTypeMax
	ConstraintTypeMin
	ConstraintTypeRange
	ConstraintTypeSet
)

var nameConstraintType map[string]Type

var constraintTypeNames = map[Type]string{
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
	nameConstraintType = make(map[string]Type, len(constraintTypeNames))
	for i, q := range constraintTypeNames {
		nameConstraintType[q] = i
	}
}

func (ct Type) MarshalJSON() ([]byte, error) {
	v, ok := constraintTypeNames[ct]
	if !ok {
		return nil, fmt.Errorf("unknown constraint type %d", ct)
	}
	return json.Marshal(v)
}

func (ct *Type) UnmarshalJSON(bytes []byte) error {
	var name string
	err := json.Unmarshal(bytes, &name)
	if err != nil {
		return err
	}
	v, ok := nameConstraintType[name]
	if !ok {
		return fmt.Errorf("unknown constraint type %s", name)
	}
	*ct = v
	return nil
}

type Constraint interface {
	Type() Type
	ASCIIDocString(dataType *types.DataType) string
	Equal(o Constraint) bool
	Min(c Context) (min types.DataTypeExtreme)
	Max(c Context) (max types.DataTypeExtreme)
	Default(c Context) (max types.DataTypeExtreme)
	Clone() Constraint
}

type Limit interface {
	ASCIIDocString(dataType *types.DataType) string
	DataModelString(dataType *types.DataType) string
	Equal(o Limit) bool
	Min(c Context) (min types.DataTypeExtreme)
	Max(c Context) (max types.DataTypeExtreme)
	Default(c Context) (max types.DataTypeExtreme)
	Clone() Limit
}

func AppendConstraint(c Constraint, n ...Constraint) Constraint {
	if c == nil {
		return Set(n)
	}
	switch c := c.(type) {
	// If the only constraint is no constraint, just replace it with the one provided
	case *AllConstraint:
		return Set(n)
	// If the only constraint is an empty generic constraint, just replace it with the one provided
	case *GenericConstraint:
		if c.Value == "" {
			return Set(n)
		}
	case Set:
		return append(c, n...)
	}
	return append(Set([]Constraint{c}), n...)
}
