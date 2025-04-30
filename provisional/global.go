package provisional

import (
	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func compareGlobals(specs specs, violations *Violations) {
	compareGlobalEntities[*matter.Bitmap](specs, violations)
	compareGlobalEntities[*matter.Enum](specs, violations)
	compareGlobalEntities[*matter.Struct](specs, violations)
	compareGlobalEntities[*matter.Command](specs, violations)
	compareGlobalEntities[*matter.Event](specs, violations)
}

func compareGlobalEntities[T ComparableEntity](specs specs, violations *Violations) (states []EntityState[T]) {
	baseEntities := types.FilterSet[T](specs.Base.GlobalObjects)
	baseInProgressEntities := types.FilterSet[T](specs.BaseInProgress.GlobalObjects)
	headEntities := types.FilterSet[T](specs.Head.GlobalObjects)
	headInProgressEntities := types.FilterSet[T](specs.HeadInProgress.GlobalObjects)

	for _, e := range headInProgressEntities {
		state := EntityState[T]{HeadInProgress: e}
		state.Head = findExistingEntity(e, internal.Iterate(headEntities))
		state.BaseInProgress = findExistingEntity(e, internal.Iterate(baseInProgressEntities))
		state.Base = findExistingEntity(e, internal.Iterate(baseEntities))
		compareStates(specs.HeadInProgress, violations, state)
	}
	return
}
