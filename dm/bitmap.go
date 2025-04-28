package dm

import (
	"fmt"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func renderBitmaps(bitmaps []*matter.Bitmap, dt *etree.Element) (err error) {
	bs := make([]*matter.Bitmap, len(bitmaps))
	copy(bs, bitmaps)
	slices.SortStableFunc(bs, func(a, b *matter.Bitmap) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, bm := range bs {
		en := dt.CreateElement("bitmap")
		en.CreateAttr("name", bm.Name)
		size := bm.Size() / 4
		for _, v := range bm.Bits {
			err = renderBit(en, v, size, bm)
			if err != nil {
				return
			}
		}
	}
	return
}

func renderBit(en *etree.Element, v matter.Bit, size int, parentEntity types.Entity) (err error) {
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
	i.CreateAttr("summary", scrubDescription(v.Summary()))
	err = renderConformanceElement(v.Conformance(), i, parentEntity)
	return
}
