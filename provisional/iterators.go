package provisional

import (
	"iter"

	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func iterateClusters(spec *spec.Specification) iter.Seq[*matter.Cluster] {
	return internal.IterateKeys(spec.Clusters)
}

func iterateFeatures(c *matter.Cluster) iter.Seq[*matter.Feature] {
	return func(yield func(*matter.Feature) bool) {
		for f := range c.Features.FeatureBits() {
			if !yield(f) {
				return
			}
		}
	}
}

func iterateAttributes(c *matter.Cluster) iter.Seq[*matter.Field] {
	return internal.Iterate(c.Attributes)
}

func iterateBitmaps(c *matter.Cluster) iter.Seq[*matter.Bitmap] {
	return internal.Iterate(c.Bitmaps)
}

func iterateBits(bm *matter.Bitmap) iter.Seq[matter.Bit] {
	return internal.Iterate(bm.Bits)
}

func iterateEnums(c *matter.Cluster) iter.Seq[*matter.Enum] {
	return internal.Iterate(c.Enums)
}

func iterateEnumValues(bm *matter.Enum) iter.Seq[*matter.EnumValue] {
	return internal.Iterate(bm.Values)
}

func iterateStructs(c *matter.Cluster) iter.Seq[*matter.Struct] {
	return internal.Iterate(c.Structs)
}

func iterateStructFields(s *matter.Struct) iter.Seq[*matter.Field] {
	return internal.Iterate(s.Fields)
}

func iterateCommands(c *matter.Cluster) iter.Seq[*matter.Command] {
	return internal.Iterate(c.Commands)
}

func iterateCommandFields(c *matter.Command) iter.Seq[*matter.Field] {
	return internal.Iterate(c.Fields)
}

func iterateEvents(c *matter.Cluster) iter.Seq[*matter.Event] {
	return internal.Iterate(c.Events)
}

func iterateEventFields(e *matter.Event) iter.Seq[*matter.Field] {
	return internal.Iterate(e.Fields)
}
