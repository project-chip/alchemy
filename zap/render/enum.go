package render

import (
	"fmt"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

func (r *renderer) renderEnums(enums map[*matter.Enum]struct{}, cx *etree.Element) {

	ens := make([]*matter.Enum, 0, len(enums))
	for en := range enums {
		ens = append(ens, en)
	}

	slices.SortFunc(ens, func(a, b *matter.Enum) int {
		return strings.Compare(a.Name, b.Name)
	})

	for _, v := range ens {
		var valFormat string
		switch v.Type.BaseType {
		case matter.BaseDataTypeEnum16:
			valFormat = "0x%04X"
		default:
			valFormat = "0x%02X"
		}

		en := cx.CreateElement("enum")
		en.CreateAttr("name", zap.CleanName(v.Name))
		if v.Type != nil {
			en.CreateAttr("type", zap.ConvertDataTypeNameToZap(v.Type.Name))
		} else {
			en.CreateAttr("type", "enum8")
		}

		r.renderClusterCodes(en, v)

		for _, ev := range v.Values {
			if conformance.IsZigbee(v.Values, ev.Conformance) {
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
