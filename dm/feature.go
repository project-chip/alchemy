package dm

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderFeatures(cluster *matter.Cluster, c *etree.Element) (err error) {
	if cluster.Features == nil || len(cluster.Features.Bits) == 0 {
		return
	}
	features := c.CreateElement("features")
	for _, b := range cluster.Features.Bits {
		f, ok := b.(*matter.Feature)
		if !ok {
			err = fmt.Errorf("feature bits contains non-feature bit %s on cluster %s ", b.Name(), cluster.Name)
			return
		}
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
