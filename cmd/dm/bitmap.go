package dm

import (
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderBitmaps(cluster *matter.Cluster, dt *etree.Element) (err error) {
	bitmaps := make([]*matter.Bitmap, len(cluster.Bitmaps))
	copy(bitmaps, cluster.Bitmaps)
	slices.SortFunc(bitmaps, func(a, b *matter.Bitmap) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, bm := range bitmaps {
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
			err = renderConformanceString(cluster, v.Conformance, i)
			if err != nil {
				return
			}
		}
	}
	return
}
