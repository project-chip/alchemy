package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func validateDeviceTypes(spec *Specification) {
	iu := make(idUniqueness[*matter.DeviceType])
	nu := make(nameUniqueness[*matter.DeviceType])
	for _, dt := range spec.DeviceTypes {
		iu.check(spec, dt.ID, dt)
		nu.check(spec, dt)
		validateDeviceType(spec, dt)
	}
}

func validateDeviceType(spec *Specification, dt *matter.DeviceType) {
	if dt.SupersetOf != "" && dt.SubsetDeviceType == nil {
		slog.Error("unknown superset device type", slog.String("deviceType", dt.Name), slog.String("superset", dt.SupersetOf), log.Path("source", dt))
		spec.addError(&UnknownSupersetError{DeviceType: dt})
	}

	validateConditions(spec, dt.Conditions)

	referencedClusters := make(map[*matter.Cluster]struct{})
	buildReferencedClusters(spec.BaseDeviceType, referencedClusters)
	buildReferencedClusters(dt, referencedClusters)

	validateClusterRequirements(spec, dt, dt.ClusterRequirements)
	validateElementRequirements(spec, dt, dt.ElementRequirements, referencedClusters)

	requiredDeviceTypes := validateDeviceTypeRequirements(spec, dt.DeviceTypeRequirements)

	validateConditionRequirements(spec, dt, requiredDeviceTypes, dt.ConditionRequirements)
	validateTagRequirements(spec, dt.TagRequirements)

	validateDeviceTypeClusterRequirements(spec, dt, requiredDeviceTypes, dt.ComposedDeviceTypeClusterRequirements)
	validateDeviceTypeElementRequirements(spec, dt, dt.ComposedDeviceTypeElementRequirements, requiredDeviceTypes, referencedClusters)
	validateDeviceTypeTagRequirements(spec, dt.ComposedDeviceTagRequirements, requiredDeviceTypes)
}

func checkDeviceReference[SourceT types.Entity](spec *Specification, deviceTypeRef matter.EntityReference, deviceType *matter.DeviceType, referencedDeviceTypes map[*matter.DeviceType]*matter.DeviceTypeRequirement, source SourceT) bool {
	if deviceType == spec.BaseDeviceType && deviceTypeRef.Name == "Base" {
		// This is a common abbreviation
		return true
	}
	if checkEntityReference(spec, deviceTypeRef, deviceType, types.EntityTypeDeviceType, source) && referencedDeviceTypes != nil {
		if _, ok := referencedDeviceTypes[deviceType]; !ok && deviceType != spec.RootNodeDeviceType && deviceType != spec.BaseDeviceType {
			slog.Error("Composing Device Type Element Requirement device type not reachable", slog.String("deviceType", deviceType.Name), slog.String("deviceTypeRef", deviceTypeRef.Name), slog.String("referencedName", deviceType.Name), log.Path("source", source))
			spec.addError(&UnreferencedEntityError[*matter.DeviceType, SourceT]{Entity: deviceType, Source: source})
		}
		return true
	}
	return false
}

func validateDeviceTypeRequirements(spec *Specification, deviceTypeRequirements []*matter.DeviceTypeRequirement) (requiredDeviceTypes map[*matter.DeviceType]*matter.DeviceTypeRequirement) {
	requiredDeviceTypes = make(map[*matter.DeviceType]*matter.DeviceTypeRequirement)
	cv := make(conformanceValidation)
	for _, dtr := range deviceTypeRequirements {
		validateConformance(spec, dtr, dtr.Conformance)
		cv.add(dtr, dtr.Conformance)
		checkDeviceReference(spec, dtr.DeviceTypeRef, dtr.DeviceType, nil, dtr)
		if dtr.DeviceType != nil {
			requiredDeviceTypes[dtr.DeviceType] = dtr
		}
	}
	cv.check(spec)
	return
}

func validateTagRequirements(spec *Specification, tagRequirements []*matter.TagRequirement) {
	cv := make(conformanceValidation)
	for _, tr := range tagRequirements {
		validateConformance(spec, tr, tr.Conformance)
		validateTagRequirement(spec, tr)
		cv.add(tr, tr.Conformance)
	}
	cv.check(spec)
}

func validateTagRequirement(spec *Specification, tr *matter.TagRequirement) {
	checkEntityReference(spec, tr.SemanticTagRef, tr.SemanticTag, types.EntityTypeSemanticTag, tr)
	if !checkEntityReference(spec, tr.NamespaceRef, tr.Namespace, types.EntityTypeNamespace, tr) {
		return
	}
	if tr.SemanticTag == nil {
		return
	}
	inNamespace := false
	for _, t := range tr.Namespace.SemanticTags {
		if t == tr.SemanticTag {
			inNamespace = true
			break
		}
	}
	if !inNamespace {
		spec.addError(&NamespaceUnknownTagError{Namespace: tr.Namespace, Requirement: tr})
	}
}

func validateDeviceTypeClusterRequirements(spec *Specification, deviceType *matter.DeviceType, referencedDeviceTypes map[*matter.DeviceType]*matter.DeviceTypeRequirement, clusterRequirements []*matter.DeviceTypeClusterRequirement) {
	cv := make(conformanceValidation)
	for _, cer := range clusterRequirements {
		cv.add(cer.ClusterRequirement, cer.ClusterRequirement.Conformance)
		if !checkDeviceReference(spec, cer.DeviceTypeRef, cer.DeviceType, referencedDeviceTypes, cer) {
			continue
		}
		validateClusterRequirement(spec, deviceType, cer.ClusterRequirement)
	}
	cv.check(spec)
}

func validateDeviceTypeElementRequirements(spec *Specification, deviceType *matter.DeviceType, elementRequirements []*matter.DeviceTypeElementRequirement, referencedDeviceTypes map[*matter.DeviceType]*matter.DeviceTypeRequirement, referencedClusters map[*matter.Cluster]struct{}) {
	cv := make(conformanceValidation)
	for _, cer := range elementRequirements {
		cv.add(cer, cer.ElementRequirement.Conformance)
		if !checkDeviceReference(spec, cer.DeviceTypeRef, cer.DeviceType, referencedDeviceTypes, cer) {
			continue
		}
		validateElementRequirement(spec, deviceType, cer.ElementRequirement, referencedClusters)
	}
	cv.check(spec)
}

func validateDeviceTypeTagRequirements(spec *Specification, tagRequirements []*matter.DeviceTypeTagRequirement, referencedDeviceTypes map[*matter.DeviceType]*matter.DeviceTypeRequirement) {
	cv := make(conformanceValidation)
	for _, cer := range tagRequirements {
		cv.add(cer, cer.TagRequirement.Conformance)
		checkDeviceReference(spec, cer.DeviceTypeRef, cer.DeviceType, referencedDeviceTypes, cer)
		validateTagRequirement(spec, cer.TagRequirement)
	}
	cv.check(spec)
}

func buildReferencedClusters(deviceType *matter.DeviceType, referencedClusters map[*matter.Cluster]struct{}) {
	parent := deviceType.SubsetDeviceType
	if parent != nil {
		buildReferencedClusters(parent, referencedClusters)
	}
	for _, rc := range deviceType.ClusterRequirements {
		if rc.Cluster != nil {
			referencedClusters[rc.Cluster] = struct{}{}
		}
	}
	for _, dtr := range deviceType.DeviceTypeRequirements {
		if dtr.DeviceType != nil {
			buildReferencedClusters(dtr.DeviceType, referencedClusters)
		}
	}
}
