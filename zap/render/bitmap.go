package render

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func renderBitmaps(bitmaps map[*matter.Bitmap]struct{}, clusterIDs []string, cx *etree.Element) {
	for bm := range bitmaps {
		en := cx.CreateElement("bitmap")
		en.CreateAttr("name", zap.CleanName(bm.Name))
		en.CreateAttr("type", zap.ConvertDataTypeNameToZap(bm.Type.Name))

		for _, cid := range clusterIDs {
			en.CreateElement("cluster").CreateAttr("code", cid)
		}
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
			evx.CreateAttr("mask", fmt.Sprintf("%#02X", mask))

		}

	}
}
