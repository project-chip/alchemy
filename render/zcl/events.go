package zcl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func renderEvents(cluster *matter.Cluster, cx *etree.Element) {
	for _, e := range cluster.Events {

		ex := cx.CreateElement("event")
		ex.CreateAttr("side", "server")
		id := e.ID
		eid, err := parse.HexOrDec(id)
		if err == nil {
			id = fmt.Sprintf("%#02x", eid)
		}
		ex.CreateAttr("code", id)
		ex.CreateAttr("priority", strings.ToLower(e.Priority))
		ex.CreateAttr("name", e.Name)
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
			id, err := parse.HexOrDec(f.ID)
			if err != nil {
				continue
			}
			fx := ex.CreateElement("field")
			fx.CreateAttr("id", strconv.Itoa(int(id)))
			fx.CreateAttr("name", f.Name)
			writeDataType(fx, f.Type)
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
