package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal"
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
	checkEntityReference(spec, cr.ClusterRef, cr.Cluster, types.EntityTypeCluster, cr)
	validateConformance(spec, cr, cr.Conformance)
}

func validateElementRequirements(spec *Specification, dt *matter.DeviceType, ers []*matter.ElementRequirement, referencedClusters map[*matter.Cluster]struct{}) {
	ercv := make(conformanceValidation)
	for _, er := range ers {
		validateElementRequirement(spec, dt, er, referencedClusters)
		ercv.add(er, er.Conformance)
	}
	ercv.check(spec)
}

func validateElementRequirement(spec *Specification, dt *matter.DeviceType, er *matter.ElementRequirement, referencedClusters map[*matter.Cluster]struct{}) {
	validateConformance(spec, er, er.Conformance)
	validateConstraint(spec, er, er.Constraint)
	checkEntityReference(spec, er.ClusterRef, er.Cluster, types.EntityTypeCluster, er)
	validateAccess(spec, er, er.Access)
	if checkEntityRefCrossReference(spec, er, er.ElementRef.XRef, er.Entity, er.Entity, er.Element, er) {
		if internal.IsNil(er.Entity) {
			slog.Error("unknown entity on element reference",
				slog.String("entityName", er.ElementRef.Name),
				slog.String("entityType", er.Element.String()),
				log.Path("source", er))
			spec.addError(&UnknownEntityError[*matter.ElementRequirement]{EntityType: er.Element, ElementReference: &er.ElementRef, Source: er})
		}
	}

	if er.Cluster == nil || er.Entity == nil {
		return
	}
	_, ok := referencedClusters[er.Cluster]
	if !ok {
		slog.Error("Element Requirement references non-required cluster", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterRef.ID.HexString()), slog.String("clusterName", er.ClusterRef.Name), log.Path("source", er))
		spec.addError(UnreferencedEntityError[*matter.Cluster, *matter.ElementRequirement]{Entity: er.Cluster, Source: er})
		return
	}
	if !er.Cluster.Contains(er.Entity) {
		slog.Error("Element Requirement references entity that does not belong to cluster", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterRef.ID.HexString()), slog.String("clusterName", er.ClusterRef.Name), slog.String("entityName", matter.EntityName(er.Entity)), log.Path("source", er))
		spec.addError(&UnreferencedEntityError[types.Entity, *matter.ElementRequirement]{Entity: er.Entity, Source: er})
	}
}
