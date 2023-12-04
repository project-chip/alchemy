package dm

import (
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
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
			var valFormat, fromFormat, toFormat matter.NumberFormat
			val, valFormat = matter.ParseFormattedID(v.Value)
			if !val.Valid() {
				from, fromFormat, to, toFormat = matter.ParseFormattedIDRange(v.Value)
				if !from.Valid() {
					continue
				}
			}

			i := en.CreateElement("item")
			if val.Valid() {
				switch valFormat {
				case matter.NumberFormatAuto, matter.NumberFormatInt:
					i.CreateAttr("value", val.IntString())
				case matter.NumberFormatHex:
					i.CreateAttr("value", val.HexString())
				}
			} else {
				switch fromFormat {
				case matter.NumberFormatAuto, matter.NumberFormatInt:
					i.CreateAttr("from", from.IntString())
				case matter.NumberFormatHex:
					i.CreateAttr("from", from.HexString())
				}
				switch toFormat {
				case matter.NumberFormatAuto, matter.NumberFormatInt:
					i.CreateAttr("to", to.IntString())
				case matter.NumberFormatHex:
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
