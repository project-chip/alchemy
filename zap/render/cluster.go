package render

import (
	"context"
	"fmt"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/iancoleman/strcase"
)

func renderCluster(cxt context.Context, doc *ascii.Doc, cluster *matter.Cluster, w *etree.Element, errata *Errata) error {

	cx := w.CreateElement("cluster")
	cx.CreateAttr("apiMaturity", "provisional")

	cx.CreateElement("name").SetText(cluster.Name)
	dom := cx.CreateElement("domain")
	domainName := matter.DomainNames[doc.Domain]
	dom.SetText(domainName)
	cx.CreateElement("code").SetText(cluster.ID.HexString())

	var define string
	var clusterPrefix string

	define = strcase.ToScreamingDelimited(cluster.Name+" Cluster", '_', "", true)
	if !errata.SuppressClusterDefinePrefix {
		clusterPrefix = strcase.ToScreamingDelimited(cluster.Name, '_', "", true) + "_"
		if len(errata.ClusterDefinePrefix) > 0 {
			clusterPrefix = errata.ClusterDefinePrefix
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

	for _, s := range errata.clusterOrder {
		switch s {
		case matter.SectionAttributes:
			renderAttributes(cluster, cx, clusterPrefix, errata)
		case matter.SectionCommands:
			renderCommands(cluster, cx, errata)
		case matter.SectionEvents:
			renderEvents(cluster, cx)
		}
	}

	return nil
}

func renderFeatures(cxt context.Context, features []*matter.Feature, clusters []*matter.Cluster, w *etree.Element, errata *Errata) {
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
		if f.Conformance == "Zigbee" {
			continue
		}
		bit, err := parse.HexOrDec(f.Bit)
		if err != nil {
			continue
		}
		fx := fb.CreateElement("field")
		fx.CreateAttr("name", f.Name)
		bit = (1 << bit)
		fx.CreateAttr("mask", fmt.Sprintf("%#x", bit))
	}
}
