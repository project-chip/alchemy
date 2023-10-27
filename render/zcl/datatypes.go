package zcl

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func renderDataTypes(cluster *matter.Cluster, cx *etree.Element, errata *errata) {
	clusterID := cluster.ID
	cid, err := parse.ID(cluster.ID)
	if err == nil {
		clusterID = fmt.Sprintf("%#04x", cid)
	}
	for _, s := range errata.dataTypeOrder {
		switch s {
		case matter.SectionDataTypeBitmap:
			renderBitmaps(cluster, cx, clusterID)
		case matter.SectionDataTypeEnum:
			renderEnums(cluster, cx, clusterID)
		case matter.SectionDataTypeStruct:
			renderStructs(cluster, cx, clusterID)
		}
	}
}

func renderEnums(cluster *matter.Cluster, cx *etree.Element, clusterID string) {
	for _, dt := range cluster.DataTypes {
		switch v := dt.(type) {
		case *matter.Enum:
			en := cx.CreateElement("enum")
			en.CreateAttr("name", v.Name)
			en.CreateAttr("type", massageDataType(v.Type))
			en.CreateElement("cluster").CreateAttr("code", clusterID)
			for _, ev := range v.Values {
				evx := en.CreateElement("item")
				evx.CreateAttr("name", ev.Name)
				val := ev.Value
				valNum, err := parse.ID(val)
				if err == nil {
					val = fmt.Sprintf("%#02x", valNum)
				}
				evx.CreateAttr("value", val)
			}

		}
	}
}

func renderBitmaps(cluster *matter.Cluster, cx *etree.Element, clusterID string) {
	for _, dt := range cluster.DataTypes {
		switch v := dt.(type) {

		case *matter.Bitmap:
			en := cx.CreateElement("bitmap")
			name := matter.StripDataTypeSuffixes(v.Name)
			en.CreateAttr("name", name)
			en.CreateAttr("type", massageDataType(v.Type))
			en.CreateElement("cluster").CreateAttr("code", clusterID)
			for _, ev := range v.Bits {
				bit, err := parse.ID(ev.Bit)
				if err != nil {
					continue
				}
				evx := en.CreateElement("item")
				evx.CreateAttr("name", ev.Name)
				evx.CreateAttr("value", fmt.Sprintf("%#02x", 1<<(bit-1)))
			}

		}
	}
}

func renderStructs(cluster *matter.Cluster, cx *etree.Element, clusterID string) {
	for _, dt := range cluster.DataTypes {
		switch v := dt.(type) {

		case *matter.Struct:
			en := cx.CreateElement("struct")
			en.CreateAttr("name", v.Name)
			en.CreateElement("cluster").CreateAttr("code", clusterID)
			for _, f := range v.Fields {
				fx := en.CreateElement("item")
				fx.CreateAttr("name", f.Name)
				writeDataType(fx, f.Type)
			}
		}
	}
}

func writeDataType(x *etree.Element, dt *matter.DataType) {
	if dt == nil {
		return
	}
	dts := massageDataType(dt.Name)
	if dt.IsArray {
		x.CreateAttr("type", "ARRAY")
		x.CreateAttr("entryType", dts)
	} else {
		x.CreateAttr("type", dts)
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
	case "string":
		return "CHAR_STRING"
	case "octstr":
		return "OCTET_STRING"
	case "vendor-id":
		return "VENDOR_ID"
	case "group-id":
		return "group_id"
	case "enum8":
		return "ENUM8"
	case "enum16":
		return "ENUM16"
	case "enum32":
		return "ENUM32"
	case "map8":
		return "BITMAP8"
	case "map16":
		return "BITMAP16"
	case "map32":
		return "BITMAP32"
	}
	return s
}
