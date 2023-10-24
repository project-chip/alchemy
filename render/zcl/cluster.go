package zcl

import (
	"context"
	"fmt"
	"strconv"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/matter"
	"github.com/iancoleman/strcase"
)

func renderCluster(cxt context.Context, cluster *matter.Cluster, w *etree.Element) error {
	cx := w.CreateElement("cluster")
	dom := cx.CreateElement("domain")
	dom.SetText("General")
	cx.CreateElement("name").SetText(cluster.Name)
	cx.CreateElement("code").SetText(fmt.Sprintf("%#04X", cluster.ID))
	define := strcase.ToScreamingDelimited(cluster.Name+" Cluster", '_', "", true)
	cx.CreateElement("define").SetText(define)
	client := cx.CreateElement("client")
	client.CreateAttr("init", "false")
	client.CreateAttr("tick", "false")
	client.SetText("true")
	server := cx.CreateElement("server")
	server.CreateAttr("init", "false")
	server.CreateAttr("tick", "false")
	server.SetText("true")
	cx.CreateElement("description").SetText("")

	if len(cluster.Attributes) > 0 {
		cx.CreateComment("Attributes")
	}
	for _, a := range cluster.Attributes {
		mandatory := (a.Conformance == "M")
		readAccess, writeAccess, _, _, _ := matter.ParseAccessValues(a.Access)

		if !mandatory && len(a.Conformance) > 0 {
			cx.CreateComment(fmt.Sprintf("Conformance feature %s - for now optional", a.Conformance))
		}
		attr := cx.CreateElement("attribute")
		attr.CreateAttr("side", "server")
		attr.CreateAttr("code", strconv.Itoa(a.ID))
		define = strcase.ToScreamingDelimited(a.Name, '_', "", true)
		attr.CreateAttr("define", define)
		attr.CreateAttr("type", massageDataType(a.Type))
		if a.Quality.Has(matter.QualityNullable) {
			attr.CreateAttr("isNullable", "true")
		}
		if a.Quality.Has(matter.QualityReportable) {
			attr.CreateAttr("reportable", "true")
		}
		if a.Quality.Has(matter.QualityFixed) || readAccess == "V" && writeAccess == "" {
			attr.CreateAttr("writeable", "false")
			attr.SetText(a.Name)
		} else {
			attr.CreateElement("description").SetText(a.Name)
			if readAccess != "" {
				ax := attr.CreateElement("access")
				ax.CreateAttr("op", "read")
				ax.CreateAttr("privilege", readAccess)
			}
			if writeAccess != "" {
				ax := attr.CreateElement("access")
				ax.CreateAttr("op", "write")
				ax.CreateAttr("privilege", writeAccess)
				attr.CreateAttr("writeable", "true")
			} else {
				attr.CreateAttr("writeable", "false")
			}
		}
		if a.Conformance != "M" {
			attr.CreateAttr("optional", "true")
		}

	}
	for _, e := range cluster.Events {
		readAccess, writeAccess, _, fabricSensitive, _ := matter.ParseAccessValues(e.Access)

		ex := cx.CreateElement("event")
		ex.CreateAttr("side", "server")
		ex.CreateAttr("code", fmt.Sprintf("%#04x", e.ID))
		ex.CreateAttr("name", e.Name)
		if fabricSensitive != 0 {
			ex.CreateAttr("isFabricSensitive", "true")
		}
		ex.CreateElement("description").SetText(e.Description)
		for _, f := range e.Fields {
			fx := ex.CreateElement("field")
			fx.CreateAttr("id", strconv.Itoa(f.ID))
			fx.CreateAttr("name", f.Name)
			fx.CreateAttr("type", massageDataType(f.Type))
		}
		if readAccess != "" {
			ax := ex.CreateElement("access")
			ax.CreateAttr("op", "read")
			ax.CreateAttr("privilege", readAccess)
		}
		if writeAccess != "" {
			ax := ex.CreateElement("access")
			ax.CreateAttr("op", "write")
			ax.CreateAttr("privilege", writeAccess)
		}
	}

	return nil
}
