package dm

import (
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
)

func renderEvents(cluster *matter.Cluster, c *etree.Element) (err error) {
	if len(cluster.Events) == 0 {
		return
	}
	evs := make([]*matter.Event, 0, len(cluster.Events))
	for _, e := range cluster.Events {
		if conformance.IsZigbee(cluster.Commands, e.Conformance) {
			continue
		}
		evs = append(evs, e)
	}

	slices.SortFunc(evs, func(a, b *matter.Event) int {
		return a.ID.Compare(b.ID)
	})
	events := c.CreateElement("events")
	for _, e := range evs {

		cx := events.CreateElement("event")
		cx.CreateAttr("id", e.ID.ShortHexString())
		cx.CreateAttr("name", e.Name)
		if len(e.Priority) > 0 {
			cx.CreateAttr("priority", strings.ToLower(e.Priority))
		}

		if e.Access.Invoke != matter.PrivilegeUnknown || e.Access.IsFabricSensitive() {
			a := cx.CreateElement("access")
			if e.Access.IsFabricSensitive() {
				a.CreateAttr("fabricSensitive", "true")
			}
			a.CreateAttr("readPrivilege", strings.ToLower(matter.PrivilegeNamesShort[e.Access.Invoke]))
		}
		err = renderConformanceString(cluster, e.Conformance, cx)
		if err != nil {
			return
		}

		err = renderFields(cluster, e.Fields, cx)
		if err != nil {
			return
		}
	}

	return
}
