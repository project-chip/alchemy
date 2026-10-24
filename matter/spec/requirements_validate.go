package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func validateClusterRequirements(spec *Specification, deviceType *matter.DeviceType, clusterRequirements []*matter.ClusterRequirement) {
	crcv := make(conformanceValidation)
	for _, cr := range clusterRequirements {
		validateClusterRequirement(spec, deviceType, cr)
		crcv.add(cr, cr.Conformance)
	}
	crcv.check(spec)
}

func validateClusterRequirement(spec *Specification, deviceType *matter.DeviceType, cr *matter.ClusterRequirement) {
	if cr.Cluster == nil {
		slog.Error("unknown cluster ID for cluster requirement on device type",
			slog.String("clusterId", cr.ClusterRef.ID.HexString()),
			slog.String("clusterName", cr.ClusterRef.Name),
			slog.String("deviceType", deviceType.Name),
			log.Path("source", cr))
		spec.addError(&UnknownEntityError[*matter.ClusterRequirement]{EntityType: types.EntityTypeCluster, EntityReference: &cr.ClusterRef, Source: cr})
		return
	}
	checkEntityReference(spec, cr.ClusterRef, cr.Cluster, cr)
	validateConformance(spec, cr, cr.Conformance)
}

func validateElementRequirements(spec *Specification, dt *matter.DeviceType, ers []*matter.ElementRequirement, referencedClusters map[*matter.Cluster]struct{}) {
	ercv := make(conformanceValidation)
	for _, er := range ers {
		if er.Cluster == nil {
			slog.Error("unknown cluster ID for element requirement on device type",
				slog.String("clusterId", er.ClusterRef.ID.HexString()),
				slog.String("clusterName", er.ClusterRef.Name),
				slog.String("deviceType", dt.Name),
				log.Path("source", er))
			spec.addError(&UnknownEntityError[*matter.ElementRequirement]{EntityType: types.EntityTypeCluster, EntityReference: &er.ClusterRef, Source: er})
		} else {
			checkEntityReference(spec, er.ClusterRef, er.Cluster, er)
		}

		validateAccess(spec, er, er.Access)
		validateConformance(spec, er, er.Conformance)
		validateElementRequirement(spec, dt, er, referencedClusters)
		ercv.add(er, er.Conformance)
	}
	ercv.check(spec)
}

func validateElementRequirement(spec *Specification, dt *matter.DeviceType, er *matter.ElementRequirement, referencedClusters map[*matter.Cluster]struct{}) {
	validateConformance(spec, er, er.Conformance)
	validateConstraint(spec, er, er.Constraint)
	if er.Cluster == nil {
		return
	}
	_, ok := referencedClusters[er.Cluster]
	if !ok {
		slog.Error("Element Requirement references non-required cluster", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterRef.ID.HexString()), slog.String("clusterName", er.ClusterRef.Name), log.Path("source", er))
		spec.addError(UnreferencedEntityError[*matter.Cluster, *matter.ElementRequirement]{Entity: er.Cluster, Source: er})

		return
	}
	if er.Entity == nil {
		slog.Error("Element Requirement references unknown element", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterRef.ID.HexString()), slog.String("clusterName", er.ClusterRef.Name), slog.String("elementRef", er.ElementRef.Name), log.Path("source", er))
		spec.addError(&UnknownEntityError[*matter.ElementRequirement]{EntityType: er.Element, ElementReference: &er.ElementRef, Source: er})
		return
	}
	if er.Entity.EntityType() != er.Element {
		slog.Error("Element Requirement references wrong entity type", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterRef.ID.HexString()), slog.String("clusterName", er.ClusterRef.Name), slog.String("entityName", matter.EntityName(er.Entity)), slog.String("expectedType", er.Element.String()), slog.String("foundType", er.Entity.EntityType().String()), log.Path("source", er))
		spec.addError(&ReferenceTypeMismatch{Element: er.Element, Entity: er.Entity, Source: er.ElementRef.XRef})
	}
	if !er.Cluster.Contains(er.Entity) {
		slog.Error("Element Requirement references entity that does not belong to cluster", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterRef.ID.HexString()), slog.String("clusterName", er.ClusterRef.Name), slog.String("entityName", matter.EntityName(er.Entity)), log.Path("source", er))
		spec.addError(&UnreferencedEntityError[types.Entity, *matter.ElementRequirement]{Entity: er.Entity, Source: er})
	}
}
