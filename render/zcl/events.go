package zcl

import (
	"fmt"
	"strconv"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func renderEvents(cluster *matter.Cluster, cx *etree.Element) {
	for _, e := range cluster.Events {
		readAccess, writeAccess, _, fabricSensitive, _ := matter.ParseAccessValues(e.Access)

		ex := cx.CreateElement("event")
		ex.CreateAttr("side", "server")
		id := e.ID
		eid, err := parse.ID(id)
		if err == nil {
			id = fmt.Sprintf("%#04x", eid)
		}
		ex.CreateAttr("code", id)
		ex.CreateAttr("name", e.Name)
		if fabricSensitive != 0 {
			ex.CreateAttr("isFabricSensitive", "true")
		}
		ex.CreateElement("description").SetText(e.Description)
		for _, f := range e.Fields {
			id, err := parse.ID(f.ID)
			if err != nil {
				continue
			}
			fx := ex.CreateElement("field")
			fx.CreateAttr("id", strconv.Itoa(int(id)))
			fx.CreateAttr("name", f.Name)
			writeDataType(fx, f.Type)
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
}
