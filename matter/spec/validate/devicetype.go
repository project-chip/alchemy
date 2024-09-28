package validate

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func validateDeviceTypes(spec *spec.Specification) {
	for _, dt := range spec.DeviceTypes {
		requiredClusterIDs := make(map[uint64]*matter.Cluster)
		for _, cr := range dt.ClusterRequirements {
			if !cr.ClusterID.Valid() {
				continue
			}
			clusterID := cr.ClusterID.Value()
			c, ok := spec.ClustersByID[clusterID]
			if !ok {
				slog.Error("Cluster Requirement references unknown cluster ID", slog.String("deviceType", dt.Name), slog.String("clusterId", cr.ClusterID.HexString()))
				requiredClusterIDs[clusterID] = nil
				continue
			}
			requiredClusterIDs[clusterID] = c
			name := stripName(cr.ClusterName)
			clusterName := stripName(c.Name)
			if !strings.EqualFold(name, clusterName) {
				slog.Error("Cluster Requirement mismatch", slog.String("deviceType", dt.Name), slog.String("clusterName", cr.ClusterName), slog.String("referencedName", c.Name))
				continue
			}
		}
		for _, er := range dt.ElementRequirements {
			if !er.ClusterID.Valid() {
				continue
			}
			c, ok := requiredClusterIDs[er.ClusterID.Value()]
			if !ok {
				slog.Error("Element Requirement references non-required cluster", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName))
				continue
			}
			if c == nil {
				slog.Error("Element Requirement references unknown cluster", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName))
				continue
			}
			switch er.Element {
			case types.EntityTypeAttribute:
				found := false
				for _, a := range c.Attributes {
					if strings.EqualFold(a.Name, er.Name) {
						found = true
						break
					}
				}
				if !found {
					slog.Error("Element Requirement references unknown attribute", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("attributeName", er.Name))
				}
			case types.EntityTypeFeature:
				found := false
				for _, fb := range c.Features.Bits {
					f, ok := fb.(*matter.Feature)
					if !ok {
						continue
					}
					if strings.EqualFold(f.Code, er.Name) || strings.EqualFold(f.Name(), er.Name) {
						found = true
						break
					}
				}
				if !found {
					slog.Error("Element Requirement references unknown feature", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("featureName", er.Name))
				}
			case types.EntityTypeCommand:
				found := false
				for _, cmd := range c.Commands {
					if strings.EqualFold(cmd.Name, er.Name) {
						found = true
						break
					}
				}
				if !found {
					slog.Error("Element Requirement references unknown command", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("commandName", er.Name))
				}
			case types.EntityTypeEvent:
				found := false
				for _, e := range c.Events {
					if strings.EqualFold(e.Name, er.Name) {
						found = true
						break
					}
				}
				if !found {
					slog.Error("Element Requirement references unknown event", slog.String("deviceType", dt.Name), slog.String("clusterId", er.ClusterID.HexString()), slog.String("clusterName", er.ClusterName), slog.String("commandName", er.Name))
				}

			default:
				slog.Error("Unknown entity type", slog.String("entityType", er.Element.String()))
			}
		}
	}
}
