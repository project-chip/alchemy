package dm

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
)

func renderAttributes(cluster *matter.Cluster, c *etree.Element) (err error) {
	if len(cluster.Attributes) == 0 {
		return
	}
	attributes := c.CreateElement("attributes")
	for _, a := range cluster.Attributes {
		if conformance.IsZigbee(a.Conformance) {
			continue
		}
		ax := attributes.CreateElement("attribute")
		ax.CreateAttr("id", a.ID.HexString())
		ax.CreateAttr("name", a.Name)
		renderDataType(a, ax)
		if len(a.Default) > 0 {
			ax.CreateAttr("default", a.Default)
		}
		renderAccess(ax, a)
		renderQuality(ax, a)
		err = renderConformanceString(cluster, a.Conformance, ax)
		if err != nil {
			return
		}

		err = renderConstraint(a.Constraint, a.Type, ax)
		if err != nil {
			return
		}
		renderDefault(cluster.Attributes, a, ax)
	}
	return
}
