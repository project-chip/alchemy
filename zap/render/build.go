package render

import (
	"log/slog"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func ZAPTemplateDestinations(sdkRoot string, docPath string, entities []types.Entity, errata *errata.SDK) (destinations map[string][]types.Entity) {
	destinations = make(map[string][]types.Entity)
	if len(errata.ClusterSplit) == 0 {
		newFile := zap.ClusterName(docPath, errata, entities)
		newPath := getZapPath(sdkRoot, newFile)
		destinations[newPath] = entities
		return
	}
	for clusterID, path := range errata.ClusterSplit {
		cid := matter.ParseNumber(clusterID)
		if !cid.Valid() {
			slog.Warn("invalid cluster split ID", "clusterId", clusterID)
			continue
		}
		var clusters []types.Entity
		for _, m := range entities {
			switch m := m.(type) {
			case *matter.ClusterGroup:
				for _, c := range m.Clusters {
					if c.ID.Equals(cid) {
						clusters = append(clusters, c)
					}
				}
			case *matter.Cluster:
				if m.ID.Equals(cid) {
					clusters = append(clusters, m)
				}
			}
		}
		destinations[getZapPath(sdkRoot, path)] = clusters
	}
	return
}
