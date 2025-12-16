package zap

import (
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (c *Configurator) addExtraTypes(errata *errata.SDK, entities []types.Entity) {
	if errata.ExtraTypes == nil {
		return
	}
	// Extra types added by Errata do not have any references from any other data types in the spec,
	// and thus need to be added manually to their associated clusters
	for _, m := range entities {
		switch v := m.(type) {
		case *matter.ClusterGroup:
			for _, cl := range v.Clusters {
				c.addExtraTypesForCluster(cl, errata.ExtraTypes)
			}
		case *matter.Cluster:
			c.addExtraTypesForCluster(v, errata.ExtraTypes)
		}
	}
}

func (c *Configurator) addExtraTypesForCluster(cluster *matter.Cluster, extraTypes *errata.SDKTypes) {
	for _, bm := range cluster.Bitmaps {
		if _, ok := extraTypes.Bitmaps[bm.Name]; ok {
			c.addEntityType(cluster, bm)
		}
	}
	for _, en := range cluster.Enums {
		if _, ok := extraTypes.Enums[en.Name]; ok {
			c.addEntityType(cluster, en)
		}
	}
	for _, s := range cluster.Structs {
		if _, ok := extraTypes.Structs[s.Name]; ok {
			c.addEntityType(cluster, s)
		}
	}
}
