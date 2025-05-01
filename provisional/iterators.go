package provisional

import (
	"iter"

	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/matter"
)

func iterateFeatures(c *matter.Cluster) iter.Seq[*matter.Feature] {
	return func(yield func(*matter.Feature) bool) {
		for f := range c.Features.FeatureBits() {
			if !yield(f) {
				return
			}
		}
	}
}

func iterateBits(bm *matter.Bitmap) iter.Seq[matter.Bit] {
	return internal.Iterate(bm.Bits)
}

func iterateEnumValues(bm *matter.Enum) iter.Seq[*matter.EnumValue] {
	return internal.Iterate(bm.Values)
}

func iterateStructFields(s *matter.Struct) iter.Seq[*matter.Field] {
	return internal.Iterate(s.Fields)
}

func iterateCommandFields(c *matter.Command) iter.Seq[*matter.Field] {
	return internal.Iterate(c.Fields)
}

func iterateEventFields(e *matter.Event) iter.Seq[*matter.Field] {
	return internal.Iterate(e.Fields)
}

/*func iterateBitmaps(c *matter.Cluster) iter.Seq[*matter.Feature] {
	return func(yield func(*matter.Bitmap) bool) {
		return internal.Iterate(c.Bitmaps)
	}
}

func iterateBitmaps(yield func(*matter.Bitmap) bool) {
	return internal.Iterate(c.Bitmaps)
}
*/
/*
for _, bm := range clusterState.HeadInProgress.Bitmaps {
	bitmapState := getEntityState(bm, clusterState, func(c *matter.Cluster) iter.Seq[*matter.Bitmap] {
		return internal.Iterate(c.Bitmaps)
	})
	compareStates(spec, violations, bitmapState)
	for _, bmb := range bm.Bits {
		compareEntity(spec, violations, bmb, bitmapState, func(pc *matter.Bitmap) iter.Seq[matter.Bit] {
			return internal.Iterate(pc.Bits)
		})
	}
}
for _, en := range clusterState.HeadInProgress.Enums {
	enumState := getEntityState(en, clusterState, func(c *matter.Cluster) iter.Seq[*matter.Enum] {
		return internal.Iterate(c.Enums)
	})
	compareStates(spec, violations, enumState)
	for _, env := range en.Values {
		compareEntity(spec, violations, env, enumState, func(pc *matter.Enum) iter.Seq[*matter.EnumValue] {
			return internal.Iterate(pc.Values)
		})
	}
}
for _, s := range clusterState.HeadInProgress.Structs {
	structState := getEntityState(s, clusterState, func(c *matter.Cluster) iter.Seq[*matter.Struct] {
		return internal.Iterate(c.Structs)
	})
	compareStates(spec, violations, structState)
	for _, sf := range s.Fields {
		compareEntity(spec, violations, sf, structState, func(pc *matter.Struct) iter.Seq[*matter.Field] {
			return internal.Iterate(pc.Fields)
		})
	}
}
for _, a := range clusterState.HeadInProgress.Attributes {
	compareEntity(spec, violations, a, clusterState, func(c *matter.Cluster) iter.Seq[*matter.Field] {
		return internal.Iterate(c.Attributes)
	})
}
for _, cmd := range clusterState.HeadInProgress.Commands {
	commandState := getEntityState(cmd, clusterState, func(c *matter.Cluster) iter.Seq[*matter.Command] {
		return internal.Iterate(c.Commands)
	})
	compareStates(spec, violations, commandState)
	for _, cf := range cmd.Fields {
		compareEntity(spec, violations, cf, commandState, func(pc *matter.Command) iter.Seq[*matter.Field] {
			return internal.Iterate(pc.Fields)
		})
	}
}
for _, ev := range clusterState.HeadInProgress.Events {
	eventState := getEntityState(ev, clusterState, func(c *matter.Cluster) iter.Seq[*matter.Event] {
		return internal.Iterate(c.Events)
	})
	compareStates(spec, violations, eventState)
	for _, ef := range ev.Fields {
		compareEntity(spec, violations, ef, eventState, func(pev *matter.Event) iter.Seq[*matter.Field] {
			return internal.Iterate(pev.Fields)
		})
	}
}
*/
