package render

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

func renderBitmaps(bitmaps []*matter.Bitmap, clusterIDs []string, cx *etree.Element) {
	for _, bm := range bitmaps {
		en := cx.CreateElement("bitmap")
		en.CreateAttr("name", bm.Name)
		en.CreateAttr("type", zap.ConvertDataTypeNameToZap(bm.Type))
		en.CreateAttr("apiMaturity", "provisional")

		for _, cid := range clusterIDs {
			en.CreateElement("cluster").CreateAttr("code", cid)
		}
		for _, bv := range bm.Bits {
			if bv.Conformance == "Zigbee" {
				continue
			}

			bit, err := parse.HexOrDec(bv.Bit)
			if err != nil {
				continue
			}
			evx := en.CreateElement("field")
			evx.CreateAttr("name", bv.Name)
			evx.CreateAttr("mask", fmt.Sprintf("%#02x", 1<<(bit)))

		}

	}
}
