package zcl

import (
	"fmt"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

var matterToZapMap = map[string]string{
	"bool":        "boolean",
	"uint8":       "INT8U",
	"uint16":      "INT16U",
	"uint32":      "INT32U",
	"uint64":      "INT64U",
	"enum8":       "ENUM8",
	"enum16":      "ENUM16",
	"enum32":      "ENUM32",
	"map8":        "BITMAP8",
	"map16":       "BITMAP16",
	"map32":       "BITMAP32",
	"string":      "CHAR_STRING",
	"octstr":      "OCTET_STRING",
	"elapsed-s":   "elapsed_s",
	"epoch-s":     "epoch_s",
	"epoch-us":    "epoch_us",
	"fabric-idx":  "fabric_idx",
	"node-id":     "NODE_ID",
	"vendor-id":   "VENDOR_ID",
	"group-id":    "group_id",
	"endpoint-id": "endpoint_id",
}

var zapToMatterMap map[string]string

func init() {
	zapToMatterMap = make(map[string]string, len(matterToZapMap))
	for k, v := range matterToZapMap {
		zapToMatterMap[strings.ToLower(v)] = k
	}
}

func ConvertDataTypeToZap(s string) string {
	if z, ok := matterToZapMap[s]; ok {
		return z
	}
	return s
}

func ConvertZapToDataType(s string) string {
	if z, ok := zapToMatterMap[strings.ToLower(s)]; ok {
		return z
	}
	return s
}

func renderDataTypes(cluster *matter.Cluster, clusters []*matter.Cluster, cx *etree.Element, errata *errata) {
	var clusterIDs []string
	for _, cluster := range clusters {
		id := cluster.ID
		cid, err := parse.HexOrDec(id)
		if err == nil {
			id = fmt.Sprintf("%#04x", cid)
		}
		clusterIDs = append(clusterIDs, id)
	}
	for _, s := range errata.dataTypeOrder {
		switch s {
		case matter.SectionDataTypeBitmap:
			renderBitmaps(cluster.Bitmaps, clusterIDs, cx)
		case matter.SectionDataTypeEnum:
			renderEnums(cluster.Enums, clusterIDs, cx)
		case matter.SectionDataTypeStruct:
			renderStructs(cluster.Structs, clusterIDs, cx)
		}
	}
}

func renderEnums(enums []*matter.Enum, clusterIDs []string, cx *etree.Element) {
	for _, v := range enums {
		en := cx.CreateElement("enum")
		en.CreateAttr("name", v.Name)
		en.CreateAttr("type", ConvertDataTypeToZap(v.Type))
		for _, cid := range clusterIDs {
			en.CreateElement("cluster").CreateAttr("code", cid)
		}
		for _, ev := range v.Values {
			evx := en.CreateElement("item")
			evx.CreateAttr("name", ev.Name)
			val := ev.Value
			valNum, err := parse.HexOrDec(val)
			if err == nil {
				val = fmt.Sprintf("%#02x", valNum)
			}
			evx.CreateAttr("value", val)
		}

	}
}

func renderBitmaps(bitmaps []*matter.Bitmap, clusterIDs []string, cx *etree.Element) {
	for _, v := range bitmaps {
		en := cx.CreateElement("bitmap")
		name := matter.StripDataTypeSuffixes(v.Name)
		en.CreateAttr("name", name)
		en.CreateAttr("type", ConvertDataTypeToZap(v.Type))
		for _, cid := range clusterIDs {
			en.CreateElement("cluster").CreateAttr("code", cid)
		}
		for _, ev := range v.Bits {
			bit, err := parse.HexOrDec(ev.Bit)
			if err != nil {
				continue
			}
			evx := en.CreateElement("item")
			evx.CreateAttr("name", ev.Name)
			evx.CreateAttr("value", fmt.Sprintf("%#02x", 1<<(bit-1)))
		}

	}
}

func renderStructs(structs []*matter.Struct, clusterIDs []string, cx *etree.Element) {
	for _, v := range structs {
		en := cx.CreateElement("struct")
		en.CreateAttr("name", v.Name)
		for _, cid := range clusterIDs {
			en.CreateElement("cluster").CreateAttr("code", cid)
		}
		for _, f := range v.Fields {
			fx := en.CreateElement("item")
			if f.Quality.Has(matter.QualityNullable) {
				fx.CreateAttr("isNullable", "true")
			}
			fx.CreateAttr("name", f.Name)
			writeDataType(fx, f.Type)
		}

	}
}

func writeDataType(x *etree.Element, dt *matter.DataType) {
	if dt == nil {
		return
	}
	dts := ConvertDataTypeToZap(dt.Name)
	if dt.IsArray {
		x.CreateAttr("type", "ARRAY")
		x.CreateAttr("entryType", dts)
	} else {
		x.CreateAttr("type", dts)
	}
}
