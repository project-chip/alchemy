package compare

import (
	"log/slog"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
)

func compareAccess(entityType types.EntityType, spec matter.Access, zap matter.Access) (diffs []Diff) {

	defaultAccess := matter.DefaultAccess(entityType)
	switch entityType {
	case types.EntityTypeAttribute:
		diffs = append(diffs, comparePrivilege(entityType, DiffPropertyReadAccess, spec.Read, zap.Read, defaultAccess.Read)...)
		diffs = append(diffs, comparePrivilege(entityType, DiffPropertyWriteAccess, spec.Write, zap.Write, defaultAccess.Write)...)
		if spec.OptionalWrite != zap.OptionalWrite {
			diffs = append(diffs, &BoolDiff{Type: DiffTypeMismatch, Property: DiffPropertyOptionalWrite, Spec: spec.OptionalWrite, ZAP: zap.OptionalWrite})
		}
	case types.EntityTypeField, types.EntityTypeCommandField:
		diffs = append(diffs, comparePrivilege(entityType, DiffPropertyReadAccess, spec.Read, zap.Read, defaultAccess.Read)...)
		diffs = append(diffs, comparePrivilege(entityType, DiffPropertyWriteAccess, spec.Write, zap.Write, defaultAccess.Write)...)
		diffs = append(diffs, compareScoping(entityType, DiffPropertyFabricScoping, spec.FabricScoping, zap.FabricScoping, defaultAccess.FabricScoping)...)
		diffs = append(diffs, compareSensitivity(entityType, DiffPropertyFabricSensitivity, spec.FabricSensitivity, zap.FabricSensitivity, defaultAccess.FabricSensitivity)...)
		if spec.OptionalWrite != zap.OptionalWrite {
			diffs = append(diffs, &BoolDiff{Type: DiffTypeMismatch, Property: DiffPropertyOptionalWrite, Spec: spec.OptionalWrite, ZAP: zap.OptionalWrite})
		}
	case types.EntityTypeCommand:
		diffs = append(diffs, comparePrivilege(entityType, DiffPropertyInvokeAccess, spec.Invoke, zap.Invoke, defaultAccess.Invoke)...)
		diffs = append(diffs, compareScoping(entityType, DiffPropertyFabricScoping, spec.FabricScoping, zap.FabricScoping, defaultAccess.FabricScoping)...)
		diffs = append(diffs, compareSensitivity(entityType, DiffPropertyFabricSensitivity, spec.FabricSensitivity, zap.FabricSensitivity, defaultAccess.FabricSensitivity)...)
		diffs = append(diffs, compareTiming(entityType, DiffPropertyTiming, spec.Timing, zap.Timing, defaultAccess.Timing)...)
	case types.EntityTypeEvent:
		diffs = append(diffs, comparePrivilege(entityType, DiffPropertyReadAccess, spec.Read, zap.Read, defaultAccess.Read)...)
		diffs = append(diffs, compareScoping(entityType, DiffPropertyFabricScoping, spec.FabricScoping, zap.FabricScoping, defaultAccess.FabricScoping)...)
		diffs = append(diffs, compareSensitivity(entityType, DiffPropertyFabricSensitivity, spec.FabricSensitivity, zap.FabricSensitivity, defaultAccess.FabricSensitivity)...)
	default:
		slog.Warn("unexpected entity for access comparison", "entityType", entityType)
	}
	return
}

func comparePrivilege(entityType types.EntityType, prop DiffProperty, spec matter.Privilege, zap matter.Privilege, defaultSpec matter.Privilege) (diffs []Diff) {
	if zap == matter.PrivilegeUnknown && (spec == defaultSpec || (entityType == types.EntityTypeField && spec == matter.PrivilegeView)) {
		return
	}
	if spec == matter.PrivilegeUnknown && zap != matter.PrivilegeUnknown {
		spec = defaultSpec
	}
	if spec != zap {
		diffs = append(diffs, NewPropertyDiff[matter.Privilege](DiffTypeMismatch, prop, spec, zap))
	}
	return
}

func compareScoping(entityType types.EntityType, prop DiffProperty, spec matter.FabricScoping, zap matter.FabricScoping, defaultSpec matter.FabricScoping) (diffs []Diff) {
	if zap == matter.FabricScopingUnknown && spec == defaultSpec {
		return
	}
	if spec == matter.FabricScopingUnknown {
		spec = defaultSpec
	}
	if spec != zap {
		diffs = append(diffs, NewPropertyDiff[matter.FabricScoping](DiffTypeMismatch, prop, spec, zap))
	}
	return
}

func compareSensitivity(entityType types.EntityType, prop DiffProperty, spec matter.FabricSensitivity, zap matter.FabricSensitivity, defaultSpec matter.FabricSensitivity) (diffs []Diff) {
	if zap == matter.FabricSensitivityUnknown && spec == defaultSpec {
		return
	}
	if spec == matter.FabricSensitivityUnknown {
		spec = defaultSpec
	}
	if spec != zap {
		diffs = append(diffs, NewPropertyDiff[matter.FabricSensitivity](DiffTypeMismatch, prop, spec, zap))
	}
	return
}

func compareTiming(entityType types.EntityType, prop DiffProperty, spec matter.Timing, zap matter.Timing, defaultSpec matter.Timing) (diffs []Diff) {
	if zap == matter.TimingUnknown && spec == defaultSpec {
		return
	}
	if spec == matter.TimingUnknown {
		spec = defaultSpec
	}
	if spec != zap {
		diffs = append(diffs, NewPropertyDiff[matter.Timing](DiffTypeMismatch, prop, spec, zap))
	}
	return
}
