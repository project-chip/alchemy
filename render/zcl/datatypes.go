package zcl

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func renderDataTypes(cluster *matter.Cluster, cx *etree.Element) {
	clusterID := cluster.ID
	cid, err := parse.ID(cluster.ID)
	if err == nil {
		clusterID = fmt.Sprintf("%#04x", cid)
	}
	for _, dt := range cluster.DataTypes {
		switch v := dt.(type) {
		case *matter.Enum:
			en := cx.CreateElement("enum")
			en.CreateAttr("name", v.Name)
			en.CreateElement("cluster").CreateAttr("code", clusterID)
			for _, ev := range v.Values {
				evx := en.CreateElement("item")
				evx.CreateAttr("name", ev.Name)
				evx.CreateAttr("value", fmt.Sprintf("%#04x", ev.Value))
			}
		case *matter.Bitmap:
			en := cx.CreateElement("bitmap")
			en.CreateAttr("name", v.Name)
			en.CreateElement("cluster").CreateAttr("code", clusterID)
			for _, ev := range v.Bits {
				evx := en.CreateElement("item")
				evx.CreateAttr("name", ev.Name)
				evx.CreateAttr("value", fmt.Sprintf("%#04x", 1<<(ev.Bit-1)))
			}
		}
	}
}

func massageDataType(s string) string {
	switch s {
	case "uint8":
		return "INT8U"
	case "uint16":
		return "INT16U"
	case "uint32":
		return "INT32U"
	case "fabric-idx":
		return "fabric_idx"
	case "node-id":
		return "NODE_ID"
	case "bool":
		return "boolean"
	}
	return s
}
