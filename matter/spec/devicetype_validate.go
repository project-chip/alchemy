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

	requiredDeviceTypes := validateDeviceTypeRequirements(spec, dt, dt.DeviceTypeRequirements)

	validateConditionRequirements(spec, dt, requiredDeviceTypes, dt.ConditionRequirements)
	validateTagRequirements(spec, dt, dt.TagRequirements)

	validateDeviceTypeClusterRequirements(spec, dt, requiredDeviceTypes, dt.ComposedDeviceTypeClusterRequirements)
	validateDeviceTypeElementRequirements(spec, dt, dt.ComposedDeviceTypeElementRequirements, requiredDeviceTypes, referencedClusters)
	validateDeviceTypeTagRequirements(spec, dt, dt.ComposedDeviceTagRequirements, requiredDeviceTypes)
}

func checkDeviceReference[SourceT types.Entity](spec *Specification, deviceTypeRef matter.EntityReference, deviceType *matter.DeviceType, source SourceT) {
	if deviceType == spec.BaseDeviceType && deviceTypeRef.Name == "Base" {
		// This is a common abbreviation
		return
	}
	checkEntityReference(spec, deviceTypeRef, deviceType, source)
}

func validateDeviceTypeRequirements(spec *Specification, deviceType *matter.DeviceType, deviceTypeRequirements []*matter.DeviceTypeRequirement) (requiredDeviceTypes map[*matter.DeviceType]*matter.DeviceTypeRequirement) {
	requiredDeviceTypes = make(map[*matter.DeviceType]*matter.DeviceTypeRequirement)
	cv := make(conformanceValidation)
	for _, dtr := range deviceTypeRequirements {
		validateConformance(spec, dtr, dtr.Conformance)
		cv.add(dtr, dtr.Conformance)
		if dtr.DeviceType == nil {
			slog.Error("unknown device type ID for device type requirement on composing device type",
				slog.String("deviceTypeId", dtr.DeviceTypeRef.ID.HexString()),
				slog.String("deviceTypeName", dtr.DeviceTypeRef.Name),
				slog.String("deviceType", deviceType.Name),
				log.Path("source", dtr))
			spec.addError(&UnknownEntityError[*matter.DeviceTypeRequirement]{EntityType: types.EntityTypeDeviceType, EntityReference: &dtr.DeviceTypeRef, Source: dtr})
		} else {
			checkDeviceReference(spec, dtr.DeviceTypeRef, dtr.DeviceType, dtr)
			requiredDeviceTypes[dtr.DeviceType] = dtr
		}
	}
	cv.check(spec)
	return
}

func validateTagRequirements(spec *Specification, deviceType *matter.DeviceType, tagRequirements []*matter.TagRequirement) {
	cv := make(conformanceValidation)
	for _, tr := range tagRequirements {
		validateConformance(spec, tr, tr.Conformance)
		validateTagRequirement(spec, deviceType, tr)
		cv.add(tr, tr.Conformance)
	}
	cv.check(spec)
}

func validateTagRequirement(spec *Specification, deviceType *matter.DeviceType, tr *matter.TagRequirement) {
	if tr.Namespace == nil {
		slog.Error("unknown namespace for tag requirement on device type",
			slog.String("namespaceId", tr.NamespaceRef.ID.HexString()),
			slog.String("namespaceName", tr.NamespaceRef.Name),
			slog.String("deviceType", deviceType.Name),
			log.Path("source", tr))
		spec.addError(&UnknownEntityError[*matter.TagRequirement]{EntityType: types.EntityTypeNamespace, EntityReference: &tr.NamespaceRef, Source: tr})
	} else {
		checkEntityReference(spec, tr.NamespaceRef, tr.Namespace, tr)
	}
	if tr.SemanticTag == nil {
		slog.Error("unknown semantic tag for tag requirement on device type",
			slog.String("semanticTagId", tr.SemanticTagRef.ID.HexString()),
			slog.String("semanticTagName", tr.SemanticTagRef.Name),
			slog.String("deviceType", deviceType.Name),
			log.Path("source", tr))
		spec.addError(&UnknownEntityError[*matter.TagRequirement]{EntityType: types.EntityTypeSemanticTag, EntityReference: &tr.SemanticTagRef, Source: tr})
	} else {
		checkEntityReference(spec, tr.SemanticTagRef, tr.SemanticTag, tr)
		if tr.Namespace != nil {
			inNamespace := false
			for _, t := range tr.Namespace.SemanticTags {
				if t == tr.SemanticTag {
					inNamespace = true
					break
				}
			}
			if !inNamespace {
				spec.addError(&NamespaceNameMismatchTagRequirementError{Namespace: tr.Namespace, Requirement: tr})
			}
		}
	}
}

func validateDeviceTypeClusterRequirements(spec *Specification, deviceType *matter.DeviceType, referencedDeviceTypes map[*matter.DeviceType]*matter.DeviceTypeRequirement, clusterRequirements []*matter.DeviceTypeClusterRequirement) {
	cv := make(conformanceValidation)
	for _, cer := range clusterRequirements {
		if cer.DeviceType == nil {
			slog.Error("unknown device type ID for cluster requirement on composing device type",
				slog.String("deviceTypeId", cer.DeviceTypeRef.ID.HexString()),
				slog.String("deviceTypeName", cer.DeviceTypeRef.Name),
				slog.String("deviceType", deviceType.Name),
				log.Path("source", cer.ClusterRequirement))
			spec.addError(&UnknownEntityError[*matter.DeviceTypeClusterRequirement]{EntityType: types.EntityTypeDeviceType, EntityReference: &cer.DeviceTypeRef, Source: cer})
		} else {
			checkDeviceReference(spec, cer.DeviceTypeRef, cer.DeviceType, cer)
			if _, ok := referencedDeviceTypes[cer.DeviceType]; !ok && cer.DeviceType != spec.RootNodeDeviceType && cer.DeviceType != spec.BaseDeviceType {
				slog.Error("Composing Device Type Cluster Requirement device type not reachable",
					slog.String("deviceTypeId", cer.DeviceTypeRef.ID.HexString()),
					slog.String("deviceTypeName", cer.DeviceTypeRef.Name),
					slog.String("deviceType", deviceType.Name),
					log.Path("source", cer.DeviceTypeRequirement))
				spec.addError(&UnreferencedEntityError[*matter.DeviceType, *matter.DeviceTypeClusterRequirement]{Entity: cer.DeviceType, Source: cer})
				continue
			}
		}
		validateClusterRequirement(spec, deviceType, cer.ClusterRequirement)
		cv.add(cer.ClusterRequirement, cer.ClusterRequirement.Conformance)
	}
	cv.check(spec)
}

func validateDeviceTypeElementRequirements(spec *Specification, deviceType *matter.DeviceType, elementRequirements []*matter.DeviceTypeElementRequirement, referencedDeviceTypes map[*matter.DeviceType]*matter.DeviceTypeRequirement, referencedClusters map[*matter.Cluster]struct{}) {
	cv := make(conformanceValidation)
	for _, cer := range elementRequirements {
		if cer.DeviceType == nil {
			slog.Error("unknown device type ID for cluster requirement on composing device type",
				slog.String("deviceTypeId", cer.DeviceTypeRef.ID.HexString()),
				slog.String("deviceTypeName", cer.DeviceTypeRef.Name),
				slog.String("deviceType", deviceType.Name),
				log.Path("source", cer.ElementRequirement))
			spec.addError(&UnknownEntityError[*matter.DeviceTypeElementRequirement]{EntityType: types.EntityTypeDeviceType, EntityReference: &cer.DeviceTypeRef, Source: cer})
		} else {
			checkDeviceReference(spec, cer.DeviceTypeRef, cer.DeviceType, cer)
			if _, ok := referencedDeviceTypes[cer.DeviceType]; !ok {
				slog.Error("Composing Device Type Element Requirement device type not reachable", slog.String("deviceType", deviceType.Name), slog.String("deviceTypeRef", cer.DeviceTypeRef.Name), slog.String("referencedName", cer.DeviceType.Name), log.Path("source", cer))
				spec.addError(&UnreferencedEntityError[*matter.DeviceType, *matter.DeviceTypeElementRequirement]{Entity: cer.DeviceType, Source: cer})
			}
		}
		validateElementRequirement(spec, deviceType, cer.ElementRequirement, referencedClusters)
		cv.add(cer, cer.ElementRequirement.Conformance)
	}
	cv.check(spec)
}

func validateDeviceTypeTagRequirements(spec *Specification, deviceType *matter.DeviceType, tagRequirements []*matter.DeviceTypeTagRequirement, referencedDeviceTypes map[*matter.DeviceType]*matter.DeviceTypeRequirement) {
	cv := make(conformanceValidation)
	for _, cer := range tagRequirements {
		if cer.DeviceType == nil {
			slog.Error("unknown device type ID for cluster requirement on composing device type",
				slog.String("deviceTypeId", cer.DeviceTypeRef.ID.HexString()),
				slog.String("deviceTypeName", cer.DeviceTypeRef.Name),
				slog.String("deviceType", deviceType.Name),
				log.Path("source", cer.TagRequirement))
			spec.addError(&UnknownEntityError[*matter.DeviceTypeTagRequirement]{EntityType: types.EntityTypeDeviceType, EntityReference: &cer.DeviceTypeRef, Source: cer})
		} else {
			checkDeviceReference(spec, cer.DeviceTypeRef, cer.DeviceType, cer)
			if _, ok := referencedDeviceTypes[cer.DeviceType]; !ok {
				slog.Error("Composing Device Type Element Requirement device type not reachable", slog.String("deviceType", deviceType.Name), slog.String("deviceTypeRef", cer.DeviceTypeRef.Name), slog.String("referencedName", cer.DeviceType.Name), log.Path("source", cer))
				spec.addError(&UnreferencedEntityError[*matter.DeviceType, *matter.DeviceTypeTagRequirement]{Entity: cer.DeviceType, Source: cer})
			}
		}
		validateTagRequirement(spec, deviceType, cer.TagRequirement)
		cv.add(cer, cer.TagRequirement.Conformance)
	}
	cv.check(spec)
}
