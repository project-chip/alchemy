package dm

import (
	"fmt"
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
		size := bm.Size() / 4
		for _, v := range bm.Bits {
			val := matter.ParseID(v.Bit)
			var mask uint64
			if !val.Valid() {
				var e error
				mask, e = v.Mask()
				if e != nil {
					continue
				}
			}
			i := en.CreateElement("bitfield")
			i.CreateAttr("name", v.Name)
			if mask > 0 {
				i.CreateAttr("mask", fmt.Sprintf("0x%0*X", size, mask))
			} else {
				i.CreateAttr("bit", val.IntString())
			}
			i.CreateAttr("summary", v.Summary)
			err = renderConformanceString(cluster, v.Conformance, i)
			if err != nil {
				return
			}
		}
	}
	return
}
