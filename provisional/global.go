package provisional

import (
	"iter"

	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func compareGlobals(specs spec.SpecPullRequest, violations map[string][]spec.Violation) {
	compareGlobalEntities(specs, iterateBits, violations)
	compareGlobalEntities(specs, iterateEnumValues, violations)
	compareGlobalEntities(specs, iterateStructFields, violations)
	compareGlobalEntities(specs, iterateCommandFields, violations)
	compareGlobalEntities(specs, iterateEventFields, violations)
}

func compareGlobalEntities[Parent ComparableEntity, Child ComparableEntity](specs spec.SpecPullRequest, iterator func(p Parent) iter.Seq[Child], violations map[string][]spec.Violation) {
	baseEntities := types.FilterSet[Parent](specs.Base.GlobalObjects)
	baseInProgressEntities := types.FilterSet[Parent](specs.BaseInProgress.GlobalObjects)
	headEntities := types.FilterSet[Parent](specs.Head.GlobalObjects)
	headInProgressEntities := types.FilterSet[Parent](specs.HeadInProgress.GlobalObjects)

	for _, e := range headInProgressEntities {
		state := EntityState[Parent]{HeadInProgress: e}
		state.Head = findExistingEntity(e, internal.Iterate(headEntities))
		state.BaseInProgress = findExistingEntity(e, internal.Iterate(baseInProgressEntities))
		state.Base = findExistingEntity(e, internal.Iterate(baseEntities))
		compareStates(specs.HeadInProgress, violations, state)

		for c := range iterator(e) {
			compareEntity(specs.HeadInProgress, violations, c, state, iterator)
		}
	}
}
