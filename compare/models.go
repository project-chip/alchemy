package compare

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"slices"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
)

func compareModels(specModels map[string][]types.Entity, zapModels map[string][]types.Entity) (diffs []any, err error) {
	for path, sm := range specModels {
		if filepath.Base(path) != "thread-network-diagnostics-cluster.xml" {
			//continue
		}
		zm, ok := zapModels[path]
		if !ok {
			slog.Warn("missing from ZAP models", slog.String("path", path))
			continue
		}
		slog.Debug("found in ZAP models", slog.String("path", path))

		specClusters := make(map[uint64]*matter.Cluster)
		for _, m := range sm {
			switch v := m.(type) {
			case *matter.Cluster:
				specClusters[v.ID.Value()] = v
			default:
				slog.Warn("unexpected spec model", slog.String("path", path), slog.String("type", fmt.Sprintf("%T", m)))
			}
		}
		zapClusters := make(map[uint64]*matter.Cluster)
		for _, m := range zm {
			switch v := m.(type) {
			case *matter.Cluster:
				zapClusters[v.ID.Value()] = v
			default:
				slog.Warn("unexpected ZAP model", slog.String("path", path), slog.String("type", fmt.Sprintf("%T", m)))
			}

		}
		delete(zapModels, path)
		for cid, sc := range specClusters {
			if zc, ok := zapClusters[cid]; ok {
				var clusterDiffs *ClusterDifferences
				clusterDiffs, err = compareClusters(sc, zc)
				if err != nil {
					slog.Warn("unable to compare clusters", slog.String("path", path), slog.Uint64("clusterId", cid), slog.Any("error", err))
					err = nil
				} else if clusterDiffs != nil {
					diffs = append(diffs, clusterDiffs)
				}
				delete(zapClusters, cid)
			} else {
				slog.Debug("missing from spec models", slog.Uint64("clusterId", cid), slog.String("path", path))
			}
		}
		for cid := range zapClusters {
			slog.Debug("missing from spec clusters", slog.Uint64("clusterId", cid), slog.String("path", path))
		}
	}

	var missingZapModels []string
	for path := range zapModels {
		missingZapModels = append(missingZapModels, path)
	}
	slices.Sort(missingZapModels)
	for _, path := range missingZapModels {
		slog.Warn("missing from spec models", slog.String("path", path))
	}
	slices.SortFunc(diffs, func(a, b any) int {
		acd, ok := a.(*ClusterDifferences)
		if ok {
			bcd, ok := b.(*ClusterDifferences)
			if ok {
				return strings.Compare(acd.Name, bcd.Name)
			}
		}
		if a == b {
			return 0
		}
		return 1
	})
	return
}
