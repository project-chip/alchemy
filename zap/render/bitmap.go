package render

import (
	"fmt"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/zap"
)

func (r *renderer) renderBitmaps(bitmaps map[*matter.Bitmap]bool, cx *etree.Element) {
	bms := make([]*matter.Bitmap, 0, len(bitmaps))
	for bm := range bitmaps {
		bms = append(bms, bm)
	}

	slices.SortFunc(bms, func(a, b *matter.Bitmap) int {
		return strings.Compare(a.Name, b.Name)
	})

	for _, bm := range bms {
		en := cx.CreateElement("bitmap")
		en.CreateAttr("name", zap.CleanName(bm.Name))
		en.CreateAttr("type", zap.ConvertDataTypeNameToZap(bm.Type.Name))

		r.renderClusterCodes(en, bm)

		size := bm.Size() / 4

		for _, bv := range bm.Bits {
			if conformance.IsZigbee(bm.Bits, bv.Conformance) {
				continue
			}

			mask, err := bv.Mask()
			if err != nil {
				continue
			}
			evx := en.CreateElement("field")
			name := zap.CleanName(bv.Name)
			evx.CreateAttr("name", zap.CleanName(name))
			evx.CreateAttr("mask", fmt.Sprintf("0x%0*X", size, mask))
		}
		cx.CreateCData("\n")

	}
}
