package spec

import (
	"github.com/project-chip/alchemy/matter"
)

func validateConditions(spec *Specification, conditions []*matter.Condition) {
	iu := make(idUniqueness[*matter.Condition])
	nu := make(nameUniqueness[*matter.Condition])
	for _, c := range conditions {
		iu.check(spec, c.ID, c)
		nu.check(spec, c)
	}
}

func validateConditionRequirements(spec *Specification, deviceType *matter.DeviceType, deviceTypes map[*matter.DeviceType]*matter.DeviceTypeRequirement, conditionRequirements []*matter.ConditionRequirement) {
	cv := make(conformanceValidation)
	for _, cr := range conditionRequirements {
		validateConformance(spec, cr, cr.Conformance)
		cv.add(cr, cr.Conformance)
		checkDeviceReference(spec, cr.DeviceTypeRef, cr.DeviceType, deviceTypes, cr)
	}
	cv.check(spec)
}
