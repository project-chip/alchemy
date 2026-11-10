package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func validateConditions(spec *Specification, conditions []*matter.Condition) {
	iu := make(idUniqueness[*matter.Condition])
	nu := make(nameUniqueness[*matter.Condition])
	for _, c := range conditions {
		iu.check(spec, c.ID, c)
		nu.check(spec, c)
	}
}

func validateConditionRequirements(spec *Specification, conditionRequirements []*matter.ConditionRequirement) {
	cv := make(conformanceValidation)
	for _, cr := range conditionRequirements {
		validateConformance(spec, cr, cr.Conformance)
		cv.add(cr, cr.Conformance)
		checkEntityReference(spec, cr.ConditionRef, cr.Condition, types.EntityTypeCondition, cr)
		checkComposedDeviceReference(spec, cr.DeviceTypeRef, cr.DeviceType, cr.DeviceTypeRequirement, cr)
		if cr.Condition != nil && cr.DeviceType != nil && !cr.DeviceType.Contains(cr.Condition) {
			slog.Error("Condition requirement on device type refers to condition not defined in device type", slog.String("condition", cr.Condition.Feature), slog.String("deviceType", cr.DeviceType.Name), log.Path("source", cr))
			spec.addError(&UnreferencedEntityError[*matter.Condition, *matter.ConditionRequirement]{Entity: cr.Condition, Source: cr})
		}
	}
	cv.check(spec)
}
