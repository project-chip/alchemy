package zcl

import (
	"context"
	"fmt"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
	"github.com/iancoleman/strcase"
)

func renderCluster(cxt context.Context, cluster *matter.Cluster, w *etree.Element, errata *errata) error {

	cx := w.CreateElement("cluster")
	dom := cx.CreateElement("domain")
	dom.SetText("General")
	cx.CreateElement("name").SetText(cluster.Name)
	code := cx.CreateElement("code")
	id, err := parse.ID(cluster.ID)
	if err == nil {
		code.SetText(fmt.Sprintf("%#04x", id))
	} else {
		code.SetText(cluster.ID)

	}
	var define string
	var clusterPrefix string

	define = strcase.ToScreamingDelimited(cluster.Name+" Cluster", '_', "", true)
	if !errata.suppressClusterDefinePrefix {
		clusterPrefix = strcase.ToScreamingDelimited(cluster.Name, '_', "", true) + "_"
		if len(errata.clusterDefinePrefix) > 0 {
			clusterPrefix = errata.clusterDefinePrefix
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
			renderCommands(cluster, cx)
		case matter.SectionEvents:
			renderEvents(cluster, cx)
		}
	}

	return nil
}

func renderFeatures(cxt context.Context, cluster *matter.Cluster, w *etree.Element, errata *errata) {
	id := cluster.ID
	cid, err := parse.ID(id)
	if err == nil {
		id = fmt.Sprintf("%#04x", cid)
	}
	if len(cluster.Features) == 0 {
		return
	}
	fb := w.CreateElement("bitmap")
	fb.CreateAttr("name", "Feature")
	fb.CreateAttr("type", "BITMAP32")
	fb.CreateElement("cluster").CreateAttr("code", id)
	for _, f := range cluster.Features {
		bit, err := parse.ID(f.Bit)
		if err != nil {
			continue
		}
		fx := fb.CreateElement("field")
		fx.CreateAttr("name", f.Name)
		bit = (1 << bit)
		fx.CreateAttr("mask", fmt.Sprintf("%#x", bit))
	}
}
