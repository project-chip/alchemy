package render

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func renderBitmaps(bitmaps []*matter.Bitmap, clusterIDs []string, cx *etree.Element) {
	for _, bm := range bitmaps {
		en := cx.CreateElement("bitmap")
		en.CreateAttr("name", bm.Name)
		en.CreateAttr("type", zap.ConvertDataTypeNameToZap(bm.Type))

		for _, cid := range clusterIDs {
			en.CreateElement("cluster").CreateAttr("code", cid)
		}
		for _, bv := range bm.Bits {
			if conformance.IsZigbee(bv.Conformance) {
				continue
			}

			mask, err := bv.Mask()
			if err != nil {
				continue
			}
			evx := en.CreateElement("field")
			name := zap.CleanName(bv.Name)
			evx.CreateAttr("name", name)
			evx.CreateAttr("mask", fmt.Sprintf("%#02X", mask))

		}

	}
}
