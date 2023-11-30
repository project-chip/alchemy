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
			val := matter.ParseID(v.Value)
			if !val.Valid() {
				continue
			}
			i := en.CreateElement("item")
			i.CreateAttr("value", val.IntString())
			i.CreateAttr("name", v.Name)
			if len(v.Summary) > 0 {
				i.CreateAttr("summary", v.Summary)
			}
			err = renderConformanceString(v.Conformance, i)
			if err != nil {
				return
			}

		}
	}
	return
}
