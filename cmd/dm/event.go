package dm

import (
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderEvents(cluster *matter.Cluster, c *etree.Element) (err error) {
	if len(cluster.Events) == 0 {
		return
	}
	events := c.CreateElement("events")
	for _, e := range cluster.Events {
		cx := events.CreateElement("event")
		cx.CreateAttr("id", e.ID.ShortHexString())
		cx.CreateAttr("name", e.Name)
		cx.CreateAttr("priority", strings.ToLower(e.Priority))

		if e.Access.Read != matter.PrivilegeUnknown {
			a := cx.CreateElement("access")
			a.CreateAttr("readPrivilege", strings.ToLower(matter.PrivilegeNamesShort[e.Access.Invoke]))
			if e.Access.FabricScoped {
				a.CreateAttr("fabricScoped", "true")
			}
			if e.Access.Timed {
				a.CreateAttr("timed", "true")
			}
		}
		err = renderConformanceString(e.Conformance, cx)
		if err != nil {
			return
		}

		for _, f := range e.Fields {
			if !f.ID.Valid() {
				continue
			}
			i := cx.CreateElement("field")
			i.CreateAttr("id", f.ID.IntString())
			i.CreateAttr("name", f.Name)
			renderDataType(f, i)
			if len(f.Default) > 0 {
				i.CreateAttr("default", f.Default)
			}
			err = renderConformanceString(f.Conformance, i)
			if err != nil {
				return
			}

			err = renderConstraint(f.Constraint, f.Type, i)
			if err != nil {
				return
			}

		}
	}

	return
}
