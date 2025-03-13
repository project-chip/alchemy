package constraint

import (
	"encoding/json"
	"math"
	"strings"

	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter/types"
)

type Set []Constraint

func (cs Set) Type() Type {
	return ConstraintTypeSet
}

func (cs Set) ASCIIDocString(dataType *types.DataType) string {
	var b strings.Builder
	for _, con := range cs {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		b.WriteString(text.TrimUnnecessaryParens(con.ASCIIDocString(dataType)))
	}
	return b.String()
}

func (cs Set) Equal(o Constraint) bool {
	ocs, ok := o.(Set)
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

func (cs Set) Min(c Context) (min types.DataTypeExtreme) {
	var from types.DataTypeExtreme

	from = cs[0].Min(c)
	for i := 1; i < len(cs); i++ {
		con := cs[i]
		f := con.Min(c)
		if !f.Defined() {
			continue
		}
		if !from.Defined() {
			from = f
			continue
		}

		from = minExtreme(from, f)
	}

	return from
}

func (cs Set) Max(c Context) (max types.DataTypeExtreme) {
	var to types.DataTypeExtreme

	to = cs[0].Max(c)
	for i := 1; i < len(cs); i++ {
		con := cs[i]
		t := con.Max(c)
		if !t.Defined() {
			continue
		}
		if !to.Defined() {
			to = t
			continue
		}
		to = maxExtreme(to, t)
	}
	return to
}

func (cs Set) Fallback(cc Context) (max types.DataTypeExtreme) {
	return
}

func (cs Set) Clone() Constraint {
	nc := make(Set, 0, len(cs))
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
			return f2
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
			if f2.Int64 > f1.Int64 {
				return f2
			}
			return f1
		case types.DataTypeExtremeTypeUInt64:
			if f2.UInt64 > math.MaxInt64 {
				return f2
			}
			if int64(f2.UInt64) > f1.Int64 {
				return f2
			}
			return f1
		}
	case types.DataTypeExtremeTypeUInt64:
		switch f2.Type {
		case types.DataTypeExtremeTypeUndefined:
			return f2
		case types.DataTypeExtremeTypeUInt64:
			if f2.UInt64 > f1.UInt64 {
				return f2
			}
			return f1
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

func (cs *Set) UnmarshalJSON(data []byte) (err error) {
	var list []json.RawMessage
	err = json.Unmarshal(data, &list)
	if err != nil {
		return
	}
	for _, l := range list {
		var c Constraint
		c, err = UnmarshalConstraint(l)
		if err != nil {
			return
		}
		*cs = append(*cs, c)
	}
	return
}

func UnmarshalConstraintSetJSON(list []json.RawMessage) (set Set, err error) {
	for _, l := range list {
		var c Constraint
		c, err = UnmarshalConstraint(l)
		if err != nil {
			return
		}
		set = append(set, c)
	}
	return
}
