package zcl

import (
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderEvents(cluster *matter.Cluster, cx *etree.Element) {
	for _, e := range cluster.Events {

		ex := cx.CreateElement("event")
		ex.CreateAttr("code", e.ID.HexString())
		ex.CreateAttr("name", e.Name)
		ex.CreateAttr("priority", strings.ToLower(e.Priority))
		ex.CreateAttr("side", "server")
		ex.CreateAttr("apiMaturity", "provisional")
		if e.Access.FabricSensitive {
			ex.CreateAttr("isFabricSensitive", "true")
		}
		if e.Conformance != "M" {
			ex.CreateAttr("optional", "true")
		}

		if len(e.Description) > 0 {
			ex.CreateElement("description").SetText(e.Description)
		} else {
			ex.CreateElement("description").SetText(e.Name)

		}
		for _, f := range e.Fields {
			if f.Conformance == "Zigbee" {
				continue
			}
			if !f.ID.Valid() {
				continue
			}
			fx := ex.CreateElement("field")
			fx.CreateAttr("fieldId", f.ID.IntString())
			fx.CreateAttr("name", f.Name)
			writeDataType(fx, f.Type)
			fx.CreateAttr("apiMaturity", "provisional")
		}
		if e.Access.Read != matter.PrivilegeUnknown {
			ax := ex.CreateElement("access")
			ax.CreateAttr("op", "read")
			ax.CreateAttr("privilege", renderPrivilege(e.Access.Read))
		}
		if e.Access.Write != matter.PrivilegeUnknown {
			ax := ex.CreateElement("access")
			ax.CreateAttr("op", "write")
			ax.CreateAttr("privilege", renderPrivilege(e.Access.Write))
		}
	}
}
