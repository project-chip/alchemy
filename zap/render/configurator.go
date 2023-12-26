package render

import (
	"context"
	"slices"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

func (r *renderer) renderModels(cxt context.Context, doc *ascii.Doc, root *etree.Element) (err error) {
	c := root.CreateElement("configurator")
	dom := c.CreateElement("domain")
	dom.CreateAttr("name", "CHIP")

	var clusters []*matter.Cluster
	for c := range r.configurator.Clusters {
		clusters = append(clusters, c)
	}

	slices.SortFunc(clusters, func(a, b *matter.Cluster) int {
		return a.ID.Compare(b.ID)
	})

	if len(clusters) > 0 {
		r.renderFeatures(cxt, clusters[0].Features, clusters, c)
	}

	r.renderBitmaps(r.configurator.Bitmaps, c)
	r.renderEnums(r.configurator.Enums, c)
	r.renderStructs(r.configurator.Structs, c)

	for _, cluster := range clusters {
		r.renderCluster(cxt, doc, cluster, c)
	}

	if err != nil {
		return
	}
	return
}
