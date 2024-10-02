package dm

import (
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
)

func renderEvents(doc *spec.Doc, cluster *matter.Cluster, c *etree.Element) (err error) {
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

	slices.SortStableFunc(evs, func(a, b *matter.Event) int {
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

		if e.Access.Read != matter.PrivilegeUnknown || e.Access.IsFabricSensitive() {
			a := cx.CreateElement("access")
			a.CreateAttr("readPrivilege", strings.ToLower(matter.PrivilegeNamesShort[e.Access.Read]))
			if e.Access.IsFabricSensitive() {
				a.CreateAttr("fabricSensitive", "true")
			}
		}
		err = RenderConformanceElement(doc, cluster, e.Conformance, cx)
		if err != nil {
			return
		}

		err = renderFields(doc, cluster, e.Fields, cx)
		if err != nil {
			return
		}
	}

	return
}
