package constraint

import (
	"math"
	"strings"

	"github.com/hasty/alchemy/matter/types"
)

type ConstraintSet []Constraint

func (c ConstraintSet) Type() ConstraintType {
	return ConstraintTypeSet
}

func (cs ConstraintSet) AsciiDocString(dataType *types.DataType) string {
	var b strings.Builder
	for _, con := range cs {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		b.WriteString(con.AsciiDocString(dataType))
	}
	return b.String()
}

func (cs ConstraintSet) Equal(o Constraint) bool {
	ocs, ok := o.(ConstraintSet)
	if !ok {
		return false
	}
	if len(cs) != len(ocs) {
		return false
	}
	for i, c := range cs {
		if !ocs[i].Equal(c) {
			return false
		}
	}
	return true
}

func (cs ConstraintSet) Min(c Context) (min types.DataTypeExtreme) {
	var from types.DataTypeExtreme

	from = cs[0].Min(c)
	for i := 1; i < len(cs); i++ {
		con := cs[i]
		f := con.Min(c)
		from = minExtreme(from, f)
	}

	return from
}

func (cs ConstraintSet) Max(c Context) (max types.DataTypeExtreme) {
	var to types.DataTypeExtreme

	to = cs[0].Min(c)
	for i := 1; i < len(cs); i++ {
		con := cs[i]
		t := con.Max(c)
		to = maxExtreme(to, t)
	}
	return to
}

func (c ConstraintSet) Default(cc Context) (max types.DataTypeExtreme) {
	return
}

func (cs ConstraintSet) Clone() Constraint {
	nc := make(ConstraintSet, 0, len(cs))
	for _, c := range cs {
		nc = append(nc, c.Clone())
	}
	return nc
}

func minExtreme(f1 types.DataTypeExtreme, f2 types.DataTypeExtreme) types.DataTypeExtreme {
	switch f1.Type {
	case types.DataTypeExtremeTypeUndefined:
		return f1
	case types.DataTypeExtremeTypeInt64:
		switch f2.Type {
		case types.DataTypeExtremeTypeUndefined:
			return f2
		case types.DataTypeExtremeTypeInt64:
			if f1.Int64 < f2.Int64 {
				return f1
			}
			return f2
		case types.DataTypeExtremeTypeUInt64:
			if f2.UInt64 > math.MaxInt64 {
				return f1
			}
			if f1.Int64 < int64(f2.UInt64) {
				return f1
			}
			return f2
		}
	case types.DataTypeExtremeTypeUInt64:
		switch f2.Type {
		case types.DataTypeExtremeTypeUndefined:
			return f2
		case types.DataTypeExtremeTypeUInt64:
			if f1.UInt64 < f2.UInt64 {
				return f1
			}
		case types.DataTypeExtremeTypeInt64:
			if f1.UInt64 > math.MaxInt64 {
				return f2
			}
			if f2.Int64 > int64(f1.UInt64) {
				return f1
			}
			return f2
		}
	}
	return types.DataTypeExtreme{}
}

func maxExtreme(f1 types.DataTypeExtreme, f2 types.DataTypeExtreme) types.DataTypeExtreme {
	switch f1.Type {
	case types.DataTypeExtremeTypeUndefined:
		return f1
	case types.DataTypeExtremeTypeInt64:
		switch f2.Type {
		case types.DataTypeExtremeTypeUndefined:
			return f2
		case types.DataTypeExtremeTypeInt64:
			if f1.Int64 > f2.Int64 {
				return f1
			}
			return f2
		case types.DataTypeExtremeTypeUInt64:
			if f2.UInt64 > math.MaxInt64 {
				return f2
			}
			if f1.Int64 < int64(f2.UInt64) {
				return f2
			}
			return f1
		}
	case types.DataTypeExtremeTypeUInt64:
		switch f2.Type {
		case types.DataTypeExtremeTypeUndefined:
			return f2
		case types.DataTypeExtremeTypeUInt64:
			if f1.UInt64 < f2.UInt64 {
				return f2
			}
		case types.DataTypeExtremeTypeInt64:
			if f1.UInt64 > math.MaxInt64 {
				return f1
			}
			if f2.Int64 > int64(f1.UInt64) {
				return f2
			}
			return f1
		}
	}
	return types.DataTypeExtreme{}
}
