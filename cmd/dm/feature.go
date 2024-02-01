package dm

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderFeatures(cluster *matter.Cluster, c *etree.Element) (err error) {
	if len(cluster.Features.Bits) == 0 {
		return
	}
	features := c.CreateElement("features")
	for _, b := range cluster.Features.Bits {
		f := b.(*matter.Feature)
		bit := matter.ParseNumber(f.Bit())
		if !bit.Valid() {
			continue
		}
		feature := features.CreateElement("feature")
		feature.CreateAttr("bit", bit.IntString())
		feature.CreateAttr("code", f.Code)
		feature.CreateAttr("name", f.Name())
		if len(f.Summary()) > 0 {
			feature.CreateAttr("summary", f.Summary())
		}
		err = renderConformanceString(cluster, f.Conformance(), feature)
		if err != nil {
			return
		}
	}
	return
}
