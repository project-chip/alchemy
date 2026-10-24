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

func validateConditionRequirements(spec *Specification, deviceType *matter.DeviceType, deviceTypes map[*matter.DeviceType]*matter.DeviceTypeRequirement, conditionRequirements []*matter.ConditionRequirement) {
	cv := make(conformanceValidation)
	for _, cr := range conditionRequirements {
		validateConformance(spec, cr, cr.Conformance)
		cv.add(cr, cr.Conformance)
		if cr.DeviceType == nil {
			slog.Error("unknown device type ID for condition requirement on device type",
				slog.String("deviceTypeId", cr.DeviceTypeRef.ID.HexString()),
				slog.String("deviceTypeName", cr.DeviceTypeRef.Name),
				slog.String("deviceType", deviceType.Name),
				log.Path("source", cr))
			spec.addError(&UnknownEntityError[*matter.ConditionRequirement]{EntityType: types.EntityTypeDeviceType, EntityReference: &cr.DeviceTypeRef, Source: cr})
		} else {
			checkDeviceReference(spec, cr.DeviceTypeRef, cr.DeviceType, cr)
			if _, ok := deviceTypes[cr.DeviceType]; !ok && cr.DeviceType != spec.RootNodeDeviceType && cr.DeviceType != spec.BaseDeviceType {
				slog.Error("Condition requirement on device type refers to unincluded device type",
					slog.String("deviceTypeId", cr.DeviceTypeRef.ID.HexString()),
					slog.String("deviceTypeName", cr.DeviceTypeRef.Name),
					slog.String("deviceType", deviceType.Name),
					log.Path("source", cr))
				spec.addError(&UnreferencedEntityError[*matter.DeviceType, *matter.ConditionRequirement]{Entity: cr.DeviceType, Source: cr})
				continue
			}
		}
		if cr.Condition == nil {
			slog.Error("unknown condition for condition requirement on device type",
				slog.String("deviceTypeId", cr.DeviceTypeRef.ID.HexString()),
				slog.String("deviceTypeName", cr.DeviceTypeRef.Name),
				slog.String("deviceType", deviceType.Name),
				slog.String("condition", cr.ConditionRef.Name),
				log.Path("source", cr))
			spec.addError(&UnknownEntityError[*matter.ConditionRequirement]{EntityType: types.EntityTypeCondition, EntityReference: &cr.ConditionRef, Source: cr})
		} else {
			checkEntityReference(spec, cr.ConditionRef, cr.Condition, cr)
			if _, ok := deviceTypes[cr.DeviceType]; !ok && cr.DeviceType != spec.RootNodeDeviceType && cr.DeviceType != spec.BaseDeviceType {
				slog.Error("Condition requirement on device type refers to unincluded device type",
					slog.String("deviceTypeId", cr.DeviceTypeRef.ID.HexString()),
					slog.String("deviceTypeName", cr.DeviceTypeRef.Name),
					slog.String("deviceType", deviceType.Name),
					log.Path("source", cr))
				spec.addError(&UnreferencedEntityError[*matter.DeviceType, *matter.ConditionRequirement]{Entity: cr.DeviceType, Source: cr})
			}
		}
	}
	cv.check(spec)
}
