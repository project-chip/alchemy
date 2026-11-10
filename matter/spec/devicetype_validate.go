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

	validateClusterRequirements(spec, dt, dt.ClusterRequirements)
	validateElementRequirements(spec, dt, dt.ElementRequirements)
	validateTagRequirements(spec, dt.TagRequirements)

	validateDeviceTypeRequirements(spec, dt.DeviceTypeRequirements)
	validateConditionRequirements(spec, dt.ConditionRequirements)
	validateDeviceTypeClusterRequirements(spec, dt.ComposedDeviceTypeClusterRequirements)
	validateDeviceTypeElementRequirements(spec, dt, dt.ComposedDeviceTypeElementRequirements)
	validateDeviceTypeTagRequirements(spec, dt.ComposedDeviceTagRequirements)
}

func checkDeviceReference[SourceT types.Entity](spec *Specification, deviceTypeRef matter.EntityReference, deviceType *matter.DeviceType, source SourceT) bool {
	if deviceType == spec.BaseDeviceType && deviceTypeRef.Name == "Base" {
		// This is a common abbreviation
		return false
	}
	if !checkEntityReference(spec, deviceTypeRef, deviceType, types.EntityTypeDeviceType, source) {
		return false
	}
	return true
}

func checkComposedDeviceReference[SourceT types.Entity](spec *Specification, deviceTypeRef matter.EntityReference, deviceType *matter.DeviceType, deviceTypeRequirement *matter.DeviceTypeRequirement, source SourceT) {
	if !checkDeviceReference(spec, deviceTypeRef, deviceType, source) {
		return
	}
	if deviceTypeRequirement == nil && deviceType != spec.BaseDeviceType && deviceType != spec.RootNodeDeviceType {
		slog.Error("Composing Device Type Element Requirement device type not reachable", slog.String("deviceType", deviceType.Name), slog.String("deviceTypeRef", deviceTypeRef.Name), slog.String("referencedName", deviceType.Name), log.Path("source", source))
		spec.addError(&UnreferencedEntityError[*matter.DeviceType, SourceT]{Entity: deviceType, Source: source})
	}
}

func checkComposedDeviceClusterReference[SourceT types.Entity](spec *Specification, clusterRef matter.EntityReference, cluster *matter.Cluster, clusterRequirement *matter.ClusterRequirement, source SourceT) {
	if !checkEntityReference(spec, clusterRef, cluster, types.EntityTypeCluster, source) {
		return
	}
	if clusterRequirement == nil {
		slog.Error("Composing Device Type Cluster Requirement cluster not reachable", slog.String("cluster", cluster.Name), slog.String("clusterRef", clusterRef.Name), log.Path("source", source))
		spec.addError(&UnreferencedEntityError[*matter.Cluster, SourceT]{Entity: cluster, Source: source})
	}
}

func validateDeviceTypeRequirements(spec *Specification, deviceTypeRequirements []*matter.DeviceTypeRequirement) (requiredDeviceTypes map[*matter.DeviceType]*matter.DeviceTypeRequirement) {
	requiredDeviceTypes = make(map[*matter.DeviceType]*matter.DeviceTypeRequirement)
	cv := make(conformanceValidation)
	for _, dtr := range deviceTypeRequirements {
		validateConformance(spec, dtr, dtr.Conformance)
		cv.add(dtr, dtr.Conformance)
		validateConstraint(spec, dtr, dtr.Constraint)
		checkDeviceReference(spec, dtr.DeviceTypeRef, dtr.DeviceType, dtr)
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
	validateConstraint(spec, tr, tr.Constraint)
	if !checkEntityReference(spec, tr.NamespaceRef, tr.Namespace, types.EntityTypeNamespace, tr) {
		return
	}
	checkEntityReference(spec, tr.SemanticTagRef, tr.SemanticTag, types.EntityTypeSemanticTag, tr)
	if tr.SemanticTag == nil || tr.Namespace == nil {
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

func validateDeviceTypeClusterRequirements(spec *Specification, clusterRequirements []*matter.DeviceTypeClusterRequirement) {
	cv := make(conformanceValidation)
	for _, cer := range clusterRequirements {
		cv.add(cer, cer.Conformance)
		checkComposedDeviceReference(spec, cer.DeviceTypeRef, cer.DeviceType, cer.DeviceTypeRequirement, cer)
		checkComposedDeviceClusterReference(spec, cer.ClusterRef, cer.Cluster, cer.ClusterRequirement, cer)
		validateConformance(spec, cer, cer.Conformance)
	}
	cv.check(spec)
}

func validateDeviceTypeElementRequirements(spec *Specification, deviceType *matter.DeviceType, elementRequirements []*matter.DeviceTypeElementRequirement) {
	cv := make(conformanceValidation)
	for _, cer := range elementRequirements {
		cv.add(cer, cer.ElementRequirement.Conformance)
		checkComposedDeviceReference(spec, cer.DeviceTypeRef, cer.DeviceType, cer.DeviceTypeRequirement, cer)
		validateElementRequirement(spec, deviceType, cer.ElementRequirement)
	}
	cv.check(spec)
}

func validateDeviceTypeTagRequirements(spec *Specification, tagRequirements []*matter.DeviceTypeTagRequirement) {
	cv := make(conformanceValidation)
	for _, cer := range tagRequirements {
		cv.add(cer, cer.TagRequirement.Conformance)
		validateConformance(spec, cer.TagRequirement, cer.TagRequirement.Conformance)
		checkComposedDeviceReference(spec, cer.DeviceTypeRef, cer.DeviceType, cer.DeviceTypeRequirement, cer)
		validateTagRequirement(spec, cer.TagRequirement)
	}
	cv.check(spec)
}

func buildReferencedClusters(spec *Specification, deviceType *matter.DeviceType) (referencedClusters map[*matter.Cluster]*matter.ClusterRequirement) {
	referencedClusters = map[*matter.Cluster]*matter.ClusterRequirement{}
	addReferencedClusters(spec.BaseDeviceType, referencedClusters)
	addReferencedClusters(spec.RootNodeDeviceType, referencedClusters)
	addReferencedClusters(deviceType, referencedClusters)
	return
}

func addReferencedClusters(deviceType *matter.DeviceType, referencedClusters map[*matter.Cluster]*matter.ClusterRequirement) {
	parent := deviceType.SubsetDeviceType
	if parent != nil {
		addReferencedClusters(parent, referencedClusters)
	}
	for _, rc := range deviceType.ClusterRequirements {
		if rc.Cluster != nil {
			referencedClusters[rc.Cluster] = rc
		}
	}
	for _, dtr := range deviceType.DeviceTypeRequirements {
		if dtr.DeviceType != nil && dtr.Location == matter.DeviceTypeRequirementLocationDeviceEndpoint {
			addReferencedClusters(dtr.DeviceType, referencedClusters)
		}
	}
}
