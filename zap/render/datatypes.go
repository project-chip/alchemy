package render

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

func renderDataTypes(cluster *matter.Cluster, clusters []*matter.Cluster, cx *etree.Element, errata *Errata) {
	var clusterIDs []string
	for _, cluster := range clusters {
		clusterIDs = append(clusterIDs, cluster.ID.HexString())
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
		if v.Type != "" {
			en.CreateAttr("type", zap.ConvertDataTypeNameToZap(v.Type))
		} else {
			en.CreateAttr("type", "enum8")
		}
		en.CreateAttr("apiMaturity", "provisional")

		for _, cid := range clusterIDs {
			en.CreateElement("cluster").CreateAttr("code", cid)
		}
		for _, ev := range v.Values {
			if ev.Conformance == "Zigbee" {
				continue
			}
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

func renderStructs(structs []*matter.Struct, clusterIDs []string, cx *etree.Element) {
	for _, v := range structs {
		en := cx.CreateElement("struct")
		en.CreateAttr("name", v.Name)
		en.CreateAttr("apiMaturity", "provisional")
		if v.FabricScoped {
			en.CreateAttr("isFabricScoped", "true")
		}
		for _, cid := range clusterIDs {
			en.CreateElement("cluster").CreateAttr("code", cid)
		}
		for _, f := range v.Fields {
			if f.Conformance == "Zigbee" {
				continue
			}
			fx := en.CreateElement("item")
			fx.CreateAttr("fieldId", f.ID.IntString())
			fx.CreateAttr("name", f.Name)
			writeDataType(fx, f.Type)
			renderConstraint(v.Fields, f, fx)
			if f.Quality.Has(matter.QualityNullable) {
				fx.CreateAttr("isNullable", "true")
			}
			if f.Conformance != "M" {
				fx.CreateAttr("optional", "true")
			}
		}

	}
}

func writeAttributeDataType(x *etree.Element, dt *matter.DataType) {
	if dt == nil {
		return
	}
	dts := zap.ConvertDataTypeNameToZap(dt.Name)
	if dt.IsArray {
		x.CreateAttr("type", "ARRAY")
		x.CreateAttr("entryType", dts)
	} else {
		x.CreateAttr("type", dts)
	}
}

func writeDataType(x *etree.Element, dt *matter.DataType) {
	if dt == nil {
		return
	}
	dts := zap.ConvertDataTypeNameToZap(dt.Name)
	if dt.IsArray {
		x.CreateAttr("array", "true")
	}
	x.CreateAttr("type", dts)
}
