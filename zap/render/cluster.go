package render

import (
	"context"
	"fmt"
	"log/slog"
	"slices"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
	"github.com/iancoleman/strcase"
)

func (r *renderer) renderCluster(cxt context.Context, cluster *matter.Cluster, w *etree.Element) error {

	cx := w.CreateElement("cluster")
	cx.CreateAttr("apiMaturity", "provisional")

	cx.CreateElement("name").SetText(cluster.Name)
	dom := cx.CreateElement("domain")
	domainName := matter.DomainNames[r.doc.Domain]
	dom.SetText(domainName)
	cx.CreateElement("code").SetText(cluster.ID.HexString())

	var define string
	var clusterPrefix string

	define = strcase.ToScreamingDelimited(cluster.Name+" Cluster", '_', "", true)
	if !r.errata.SuppressClusterDefinePrefix {
		clusterPrefix = strcase.ToScreamingDelimited(cluster.Name, '_', "", true) + "_"
		if len(r.errata.ClusterDefinePrefix) > 0 {
			clusterPrefix = r.errata.ClusterDefinePrefix
		}
	}

	cx.CreateElement("define").SetText(define)
	client := cx.CreateElement("client")
	client.CreateAttr("init", "false")
	client.CreateAttr("tick", "false")
	client.SetText("true")
	server := cx.CreateElement("server")
	server.CreateAttr("init", "false")
	server.CreateAttr("tick", "false")
	server.SetText("true")
	cx.CreateElement("description").SetText(cluster.Description)

	renderAttributes(cluster, cx, clusterPrefix, r.errata)
	renderCommands(cluster, cx, r.errata)
	renderEvents(cluster, cx)

	return nil
}

func (r *renderer) renderFeatures(cxt context.Context, features matter.FeatureSet, clusters []*matter.Cluster, w *etree.Element) {
	if len(features) == 0 {
		return
	}
	fb := w.CreateElement("bitmap")
	fb.CreateAttr("name", "Feature")
	fb.CreateAttr("type", "bitmap32")
	for _, cluster := range clusters {
		fb.CreateElement("cluster").CreateAttr("code", cluster.ID.HexString())
	}
	for _, f := range features {
		if conformance.IsZigbee(features, f.Conformance) {
			continue
		}
		bit, err := parse.HexOrDec(f.Bit)
		if err != nil {
			continue
		}
		fx := fb.CreateElement("field")
		fx.CreateAttr("name", zap.CleanName(f.Name))
		bit = (1 << bit)
		fx.CreateAttr("mask", fmt.Sprintf("%#x", bit))
	}
}

func (r *renderer) renderClusterCodes(parent *etree.Element, model matter.Model) {
	refs, ok := r.spec.ClusterRefs[model]
	if !ok {
		slog.Warn("unknown cluster ref", "val", model)
		return
	}
	var clusterIDs []string
	for ref := range refs {
		clusterIDs = append(clusterIDs, ref.ID.HexString())
	}
	slices.Sort(clusterIDs)
	for _, cid := range clusterIDs {
		parent.CreateElement("cluster").CreateAttr("code", cid)
	}

}
