package dm

import (
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
)

func renderEnums(cluster *matter.Cluster, dt *etree.Element) (err error) {
	enums := make([]*matter.Enum, len(cluster.Enums))
	copy(enums, cluster.Enums)
	slices.SortFunc(enums, func(a, b *matter.Enum) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, e := range enums {
		en := dt.CreateElement("enum")
		en.CreateAttr("name", e.Name)
		for _, v := range e.Values {
			var val, from, to *matter.Number
			var valFormat, fromFormat, toFormat types.NumberFormat
			//val, valFormat = matter.ParseFormattedNumber(v.Value)
			if v.Value.Valid() {
				val = v.Value
				valFormat = v.Value.Format()
			} else {
				from, fromFormat, to, toFormat = matter.ParseFormattedIDRange(v.Value.Text())
				if !from.Valid() {
					continue
				}
			}

			i := en.CreateElement("item")
			if val.Valid() {
				switch valFormat {
				case types.NumberFormatAuto, types.NumberFormatInt:
					i.CreateAttr("value", val.IntString())
				case types.NumberFormatHex:
					i.CreateAttr("value", val.HexString())
				}
			} else {
				switch fromFormat {
				case types.NumberFormatAuto, types.NumberFormatInt:
					i.CreateAttr("from", from.IntString())
				case types.NumberFormatHex:
					i.CreateAttr("from", from.HexString())
				}
				switch toFormat {
				case types.NumberFormatAuto, types.NumberFormatInt:
					i.CreateAttr("to", to.IntString())
				case types.NumberFormatHex:
					i.CreateAttr("to", to.HexString())
				}
			}
			i.CreateAttr("name", v.Name)
			if len(v.Summary) > 0 {
				i.CreateAttr("summary", v.Summary)
			}
			err = renderConformanceString(cluster, v.Conformance, i)
			if err != nil {
				return
			}

		}
	}
	return
}
