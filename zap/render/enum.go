package render

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

func renderEnums(enums []*matter.Enum, clusterIDs []string, cx *etree.Element) {

	for _, v := range enums {
		enumType := v.Type
		if enumType != "" {
			enumType = zap.ConvertDataTypeNameToZap(v.Type)
		} else {
			enumType = "enum8"
		}
		var valFormat string
		switch enumType {
		case "enum16":
			valFormat = "0x%04X"
		default:
			valFormat = "0x%02X"
		}

		en := cx.CreateElement("enum")
		en.CreateAttr("name", v.Name)
		if v.Type != "" {
			en.CreateAttr("type", zap.ConvertDataTypeNameToZap(v.Type))
		} else {
			en.CreateAttr("type", "enum8")
		}

		for _, cid := range clusterIDs {
			en.CreateElement("cluster").CreateAttr("code", cid)
		}
		for _, ev := range v.Values {
			if conformance.IsZigbee(ev.Conformance) {
				continue
			}
			evx := en.CreateElement("item")
			name := zap.CleanName(ev.Name)
			evx.CreateAttr("name", name)
			val := ev.Value
			valNum, err := parse.HexOrDec(val)
			if err == nil {
				val = fmt.Sprintf(valFormat, valNum)
			}
			evx.CreateAttr("value", val)
		}

	}
}
