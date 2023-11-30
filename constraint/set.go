package constraint

import (
	"math"
	"strings"

	"github.com/hasty/alchemy/matter"
)

type ConstraintSet []matter.Constraint

func (cs ConstraintSet) AsciiDocString(dataType *matter.DataType) string {
	var b strings.Builder
	for _, con := range cs {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		b.WriteString(con.AsciiDocString(dataType))
	}
	return b.String()
}

func (cs ConstraintSet) Equal(o matter.Constraint) bool {
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

func (cs ConstraintSet) Min(c *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	var from matter.ConstraintExtreme

	from = cs[0].Min(c)
	for i := 1; i < len(cs); i++ {
		con := cs[i]
		f := con.Min(c)
		from = minExtreme(from, f)
	}

	return from
}

func (cs ConstraintSet) Max(c *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	var to matter.ConstraintExtreme

	to = cs[0].Min(c)
	for i := 1; i < len(cs); i++ {
		con := cs[i]
		t := con.Max(c)
		to = maxExtreme(to, t)
	}
	return to
}

func (c ConstraintSet) Default(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return
}

func minExtreme(f1 matter.ConstraintExtreme, f2 matter.ConstraintExtreme) matter.ConstraintExtreme {
	switch f1.Type {
	case matter.ConstraintExtremeTypeUndefined:
		return f1
	case matter.ConstraintExtremeTypeInt64:
		switch f2.Type {
		case matter.ConstraintExtremeTypeUndefined:
			return f2
		case matter.ConstraintExtremeTypeInt64:
			if f1.Int64 < f2.Int64 {
				return f1
			}
			return f2
		case matter.ConstraintExtremeTypeUInt64:
			if f2.UInt64 > math.MaxInt64 {
				return f1
			}
			if f1.Int64 < int64(f2.UInt64) {
				return f1
			}
			return f2
		}
	case matter.ConstraintExtremeTypeUInt64:
		switch f2.Type {
		case matter.ConstraintExtremeTypeUndefined:
			return f2
		case matter.ConstraintExtremeTypeUInt64:
			if f1.UInt64 < f2.UInt64 {
				return f1
			}
		case matter.ConstraintExtremeTypeInt64:
			if f1.UInt64 > math.MaxInt64 {
				return f2
			}
			if f2.Int64 > int64(f1.UInt64) {
				return f1
			}
			return f2
		}
	}
	return matter.ConstraintExtreme{}
}

func maxExtreme(f1 matter.ConstraintExtreme, f2 matter.ConstraintExtreme) matter.ConstraintExtreme {
	switch f1.Type {
	case matter.ConstraintExtremeTypeUndefined:
		return f1
	case matter.ConstraintExtremeTypeInt64:
		switch f2.Type {
		case matter.ConstraintExtremeTypeUndefined:
			return f2
		case matter.ConstraintExtremeTypeInt64:
			if f1.Int64 > f2.Int64 {
				return f1
			}
			return f2
		case matter.ConstraintExtremeTypeUInt64:
			if f2.UInt64 > math.MaxInt64 {
				return f2
			}
			if f1.Int64 < int64(f2.UInt64) {
				return f2
			}
			return f1
		}
	case matter.ConstraintExtremeTypeUInt64:
		switch f2.Type {
		case matter.ConstraintExtremeTypeUndefined:
			return f2
		case matter.ConstraintExtremeTypeUInt64:
			if f1.UInt64 < f2.UInt64 {
				return f2
			}
		case matter.ConstraintExtremeTypeInt64:
			if f1.UInt64 > math.MaxInt64 {
				return f1
			}
			if f2.Int64 > int64(f1.UInt64) {
				return f2
			}
			return f1
		}
	}
	return matter.ConstraintExtreme{}
}
