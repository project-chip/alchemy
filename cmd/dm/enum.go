package dm

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderEnums(cluster *matter.Cluster, dt *etree.Element) (err error) {
	for _, e := range cluster.Enums {
		en := dt.CreateElement("enum")
		en.CreateAttr("name", e.Name)
		for _, v := range e.Values {
			var val, from, to *matter.ID
			var valFormat, fromFormat, toFormat matter.ConstraintExtremeFormat
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
				case matter.ConstraintExtremeFormatAuto, matter.ConstraintExtremeFormatInt:
					i.CreateAttr("value", val.IntString())
				case matter.ConstraintExtremeFormatHex:
					i.CreateAttr("value", val.HexString())
				}
			} else {
				switch fromFormat {
				case matter.ConstraintExtremeFormatAuto, matter.ConstraintExtremeFormatInt:
					i.CreateAttr("from", from.IntString())
				case matter.ConstraintExtremeFormatHex:
					i.CreateAttr("from", from.HexString())
				}
				switch toFormat {
				case matter.ConstraintExtremeFormatAuto, matter.ConstraintExtremeFormatInt:
					i.CreateAttr("to", to.IntString())
				case matter.ConstraintExtremeFormatHex:
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
