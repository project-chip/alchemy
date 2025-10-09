package provisional

import (
	"iter"

	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func compareClusters(specs spec.SpecSet, violations map[string][]Violation) {

	var clusterStates []EntityState[*matter.Cluster]
	//var headInProgressClusters []*matter.Cluster // Clusters which only appear in the head spec when in-progress is set
	//var headClusterPairs []EntityState[*matter.Cluster]
	for cluster := range specs.HeadInProgress.Clusters {
		state := EntityState[*matter.Cluster]{HeadInProgress: cluster}
		state.Head = findExistingClusterInSpec(cluster, specs.Head)
		state.BaseInProgress = findExistingClusterInSpec(cluster, specs.BaseInProgress)
		state.Base = findExistingClusterInSpec(cluster, specs.Base)
		clusterStates = append(clusterStates, state)
	}

	for _, state := range clusterStates {
		compareStates(specs.HeadInProgress, violations, state)
		compareClusterEntities(specs.HeadInProgress, violations, state)
	}
}

func compareClusterEntities(spec *spec.Specification, violations map[string][]Violation, clusterState EntityState[*matter.Cluster]) {

	for f := range clusterState.HeadInProgress.Features.FeatureBits() {
		compareEntity(spec, violations, f, clusterState, func(c *matter.Cluster) iter.Seq[*matter.Feature] {
			return func(yield func(*matter.Feature) bool) {
				for f := range c.Features.FeatureBits() {
					if !yield(f) {
						return
					}
				}
			}
		})
	}
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
}

func compareEntity[T ComparableEntity, Parent types.Entity](spec *spec.Specification, violations map[string][]Violation, e T, parentState EntityState[Parent], iterator func(p Parent) iter.Seq[T]) {
	compareStates(spec, violations, getEntityState(e, parentState, iterator))
}

func getEntityState[T ComparableEntity, Parent types.Entity](e T, parentState EntityState[Parent], iterator func(p Parent) iter.Seq[T]) EntityState[T] {
	state := EntityState[T]{HeadInProgress: e}
	if !isNil(parentState.Head) {
		state.Head = findExistingEntity(e, iterator(parentState.Head))
	}
	if !isNil(parentState.BaseInProgress) {
		state.BaseInProgress = findExistingEntity(e, iterator(parentState.BaseInProgress))
	}
	if !isNil(parentState.Base) {
		state.Base = findExistingEntity(e, iterator(parentState.Base))
	}
	return state
}
