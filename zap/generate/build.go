package generate

import (
	"log/slog"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/zap"
)

func getDocDomain(doc *spec.Doc) matter.Domain {
	if doc.Domain != matter.DomainUnknown {
		return doc.Domain
	}
	for _, p := range doc.Parents() {
		d := getDocDomain(p)
		if d != matter.DomainUnknown {
			return d
		}
	}
	return matter.DomainUnknown
}

func ZAPTemplateDestinations(sdkRoot string, docPath string, entities []types.Entity, errata *zap.Errata) (destinations map[string][]types.Entity) {
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
			slog.Warn("invalid cluster split ID", "id", clusterID)
			continue
		}
		var clusters []types.Entity
		for _, m := range entities {
			switch m := m.(type) {
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
