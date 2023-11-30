package dm

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderBitmaps(cluster *matter.Cluster, dt *etree.Element) (err error) {
	for _, bm := range cluster.Bitmaps {
		en := dt.CreateElement("bitmap")
		en.CreateAttr("name", bm.Name)
		for _, v := range bm.Bits {
			val := matter.ParseID(v.Bit)
			if !val.Valid() {
				continue
			}
			i := en.CreateElement("bitfield")
			i.CreateAttr("name", v.Name)
			i.CreateAttr("bit", val.IntString())
			i.CreateAttr("summary", v.Summary)
			err = renderConformanceString(v.Conformance, i)
			if err != nil {
				return
			}
		}
	}
	return
}
