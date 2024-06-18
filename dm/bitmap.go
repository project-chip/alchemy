package dm

import (
	"fmt"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
)

func renderBitmaps(doc *spec.Doc, cluster *matter.Cluster, dt *etree.Element) (err error) {
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
			err = renderBit(doc, cluster, en, v, size)
			if err != nil {
				return
			}
		}
	}
	return
}

func renderBit(doc *spec.Doc, cluster *matter.Cluster, en *etree.Element, v matter.Bit, size int) (err error) {
	i := en.CreateElement("bitfield")
	i.CreateAttr("name", v.Name())
	val := matter.ParseNumber(v.Bit())
	if val.Valid() {
		i.CreateAttr("bit", val.IntString())
	} else {
		var from, to uint64
		from, to, err = v.Bits()
		if err != nil {
			var mask uint64
			mask, err = v.Mask()
			if err == nil {
				i.CreateAttr("mask", fmt.Sprintf("0x%0*X", size, mask))
			}
			return
		}
		i.CreateAttr("from", fmt.Sprintf("0x%0*X", size, from))
		i.CreateAttr("to", fmt.Sprintf("0x%0*X", size, to))
	}
	i.CreateAttr("summary", v.Summary())
	err = renderConformanceString(doc, cluster, v.Conformance(), i)
	return
}
