package render

import (
	"context"
	"slices"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func (r *renderer) renderModels(cxt context.Context, root *etree.Element, models []matter.Model) (err error) {
	c := root.CreateElement("configurator")
	dom := c.CreateElement("domain")
	dom.CreateAttr("name", "CHIP")

	bitmapSet := make(map[*matter.Bitmap]struct{})
	enumSet := make(map[*matter.Enum]struct{})
	structSet := make(map[*matter.Struct]struct{})
	clusterSet := make(map[*matter.Cluster]struct{})

	for _, m := range models {
		switch m := m.(type) {
		case *matter.Bitmap:
			bitmapSet[m] = struct{}{}
		case *matter.Enum:
			enumSet[m] = struct{}{}
		case *matter.Struct:
			structSet[m] = struct{}{}
		case *matter.Cluster:
			clusterSet[m] = struct{}{}
			for _, bm := range m.Bitmaps {
				bitmapSet[bm] = struct{}{}
			}
			for _, en := range m.Enums {
				enumSet[en] = struct{}{}
			}
			for _, s := range m.Structs {
				structSet[s] = struct{}{}
			}
		}
	}

	var clusters []*matter.Cluster
	for c := range clusterSet {
		clusters = append(clusters, c)
	}

	slices.SortFunc(clusters, func(a, b *matter.Cluster) int {
		return a.ID.Compare(b.ID)
	})

	if len(clusters) > 0 {
		r.renderFeatures(cxt, clusters[0].Features, clusters, c)
	}

	r.renderBitmaps(bitmapSet, c)
	r.renderEnums(enumSet, c)
	r.renderStructs(structSet, c)

	for _, cluster := range clusters {
		r.renderCluster(cxt, cluster, c)
	}

	if err != nil {
		return
	}
	return
}
