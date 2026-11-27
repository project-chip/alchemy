package dm

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

// resolveConformance returns c if it is non-blank; otherwise, for an entity in a
// derived cluster, it walks ParentCluster without mutating the derived entity to
// find the inherited conformance and its owning parent entity.
func resolveConformance(c conformance.Conformance, entity types.Entity, parentEntity types.Entity) (conformance.Conformance, types.Entity) {
	if c != nil && !conformance.IsBlank(c) {
		return c, parentEntity
	}
	var cluster *matter.Cluster
	switch p := parentEntity.(type) {
	case *matter.Cluster:
		cluster = p
	case interface{ Cluster() *matter.Cluster }:
		cluster = p.Cluster()
	}
	if cluster == nil || cluster.ParentCluster == nil {
		return c, parentEntity
	}

	visited := map[*matter.Cluster]bool{cluster: true}
	for base := cluster.ParentCluster; base != nil && !visited[base]; base = base.ParentCluster {
		visited[base] = true
		if bc, bp := lookupBaseConformance(base, entity, parentEntity); bc != nil && !conformance.IsBlank(bc) {
			return bc, bp
		}
	}
	return c, parentEntity
}

func lookupBaseConformance(base *matter.Cluster, entity types.Entity, parentEntity types.Entity) (conformance.Conformance, types.Entity) {
	switch e := entity.(type) {
	case *matter.Field:
		switch p := parentEntity.(type) {
		case *matter.Cluster:
			if bf := findField(base.Attributes, e); bf != nil {
				return bf.Conformance, nil
			}
		case *matter.Struct:
			if s, ok := base.Structs.Identifier(p.Name); ok {
				bs := s.(*matter.Struct)
				if bf := findField(bs.Fields, e); bf != nil {
					return bf.Conformance, bs
				}
			}
		case *matter.Command:
			if bc := findCommand(base.Commands, p); bc != nil {
				if bf := findField(bc.Fields, e); bf != nil {
					return bf.Conformance, bc
				}
			}
		case *matter.Event:
			if be := findEvent(base.Events, p); be != nil {
				if bf := findField(be.Fields, e); bf != nil {
					return bf.Conformance, be
				}
			}
		}
	case *matter.Command:
		if bc := findCommand(base.Commands, e); bc != nil {
			return bc.Conformance, bc
		}
	case *matter.Event:
		if be := findEvent(base.Events, e); be != nil {
			return be.Conformance, be
		}
	case *matter.Feature:
		if base.Features != nil {
			if f, ok := base.Features.Identifier(e.Code); ok {
				if bf, ok := f.(*matter.Feature); ok {
					return bf.Conformance(), nil
				}
			}
		}
	case *matter.EnumValue:
		if p, ok := parentEntity.(*matter.Enum); ok {
			if en, ok := base.Enums.Identifier(p.Name); ok {
				ben := en.(*matter.Enum)
				if bv := findEnumValue(ben.Values, e); bv != nil {
					return bv.Conformance, ben
				}
			}
		}
	case matter.Bit:
		if p, ok := parentEntity.(*matter.Bitmap); ok {
			if bm, ok := base.Bitmaps.Identifier(p.Name); ok {
				bbm := bm.(*matter.Bitmap)
				if bb := findBit(bbm.Bits, e); bb != nil {
					return bb.Conformance(), bbm
				}
			}
		}
	}
	return nil, nil
}

func findField(fields matter.FieldSet, target *matter.Field) *matter.Field {
	if target.ID.Valid() {
		for _, f := range fields {
			if f.ID.Valid() && f.ID.Equals(target.ID) {
				return f
			}
		}
	}
	if e, ok := fields.Identifier(target.Name); ok {
		if f, ok := e.(*matter.Field); ok && (!target.ID.Valid() || !f.ID.Valid() || target.ID.Equals(f.ID)) {
			return f
		}
	}
	return nil
}

func findCommand(commands matter.CommandSet, target *matter.Command) *matter.Command {
	if e, ok := commands.Identifier(target.Name); ok {
		if c, ok := e.(*matter.Command); ok && (!target.ID.Valid() || !c.ID.Valid() || target.ID.Equals(c.ID)) {
			return c
		}
	}
	return nil
}

func findEvent(events matter.EventSet, target *matter.Event) *matter.Event {
	if target.ID.Valid() {
		for _, ev := range events {
			if ev.ID.Valid() && ev.ID.Equals(target.ID) {
				return ev
			}
		}
	}
	if e, ok := events.Identifier(target.Name); ok {
		if ev, ok := e.(*matter.Event); ok && (!target.ID.Valid() || !ev.ID.Valid() || target.ID.Equals(ev.ID)) {
			return ev
		}
	}
	return nil
}

func findEnumValue(values matter.EnumValueSet, target *matter.EnumValue) *matter.EnumValue {
	if target.Value.Valid() {
		for _, v := range values {
			if v.Value.Valid() && v.Value.Equals(target.Value) {
				return v
			}
		}
	}
	if e, ok := values.Identifier(target.Name); ok {
		if v, ok := e.(*matter.EnumValue); ok && (!target.Value.Valid() || !v.Value.Valid() || target.Value.Equals(v.Value)) {
			return v
		}
	}
	return nil
}

func findBit(bits matter.BitSet, target matter.Bit) matter.Bit {
	for _, b := range bits {
		if len(target.Bit()) > 0 && b.Bit() == target.Bit() {
			return b
		}
	}
	if e, ok := bits.Identifier(target.Name()); ok {
		if b, ok := e.(matter.Bit); ok {
			return b
		}
	}
	return nil
}
