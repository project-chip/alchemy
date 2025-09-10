package constraint

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

func ParseString(constraint string) Constraint {
	c, err := TryParseString(constraint)
	if err != nil {
		return &GenericConstraint{Value: constraint}
	}
	return c
}

func TryParseString(constraint string) (Constraint, error) {
	if len(constraint) == 0 {
		return &GenericConstraint{Value: constraint}, nil
	}
	c, err := Parse("", []byte(constraint))
	if err != nil {
		return nil, err
	}
	return c.(Constraint), nil
}

func ParseLimit(limit string) Limit {
	l, err := TryParseLimit(limit)
	if err != nil {
		return &GenericLimit{Value: limit}
	}
	return l

}

func TryParseLimit(limit string) (Limit, error) {
	if len(limit) == 0 {
		return nil, nil
	}
	if strings.EqualFold(limit, "desc") {
		return &DescribedLimit{}, nil
	}
	l, err := Parse("", []byte(limit), Entrypoint("Limit"))
	if err != nil {
		return nil, err
	}
	return l.(Limit), nil
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
	ConstraintTypeTagList
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
	ConstraintTypeTagList:   "tagList",
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
	NeedsParens(topLevel bool) bool
	Clone() Constraint
}

type Limit interface {
	ASCIIDocString(dataType *types.DataType) string
	DataModelString(dataType *types.DataType) string
	Equal(o Limit) bool
	Min(c Context) (min types.DataTypeExtreme)
	Max(c Context) (max types.DataTypeExtreme)
	Fallback(c Context) (max types.DataTypeExtreme)
	NeedsParens(topLevel bool) bool
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

func IsBlank(c Constraint) bool {
	switch c := c.(type) {
	case *GenericConstraint:
		return c.Value == ""
	case Set:
		if len(c) == 1 {
			mc, ok := c[0].(*GenericConstraint)
			if ok {
				return mc.Value == ""
			}
		} else if len(c) == 0 {
			return true
		}

	}
	return false
}

func IsAllOrEmpty(c Constraint) bool {
	switch c := c.(type) {
	case *AllConstraint:
		return true
	case Set:
		if len(c) == 1 {
			_, ok := c[0].(*AllConstraint)
			if ok {
				return true
			}
		} else if len(c) == 0 {
			return true
		}
	}
	return false
}

func IsGeneric(c Constraint) bool {
	switch c := c.(type) {
	case *GenericConstraint:
		return true
	case Set:
		if len(c) == 1 {
			_, ok := c[0].(*GenericConstraint)
			if ok {
				return true
			}
		} else if len(c) == 0 {
			return true
		}

	}
	return false
}

func IsNull(c Constraint) bool {
	switch c := c.(type) {
	case *ExactConstraint:
		_, isNull := c.Value.(*NullLimit)
		return isNull
	case Set:
		if len(c) == 1 {
			mc, ok := c[0].(*ExactConstraint)
			if ok {
				_, isNull := mc.Value.(*NullLimit)
				return isNull
			}
		}
	}
	return false
}

func IsGenericLimit(c Limit) bool {
	switch c.(type) {
	case *GenericLimit:
		return true
	}
	return false
}

func IsBlankLimit(l Limit) bool {
	switch c := l.(type) {
	case *GenericLimit:
		return c.Value == ""
	case nil:
		return true
	}
	return false
}

func IsNullLimit(l Limit) bool {
	_, isNull := l.(*NullLimit)
	return isNull
}
