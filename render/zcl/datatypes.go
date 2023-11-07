package zcl

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
		en.CreateAttr("type", zap.ConvertDataTypeToZap(v.Type))
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
		en.CreateAttr("type", zap.ConvertDataTypeToZap(v.Type))
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
	dts := zap.ConvertDataTypeToZap(dt.Name)
	if dt.IsArray {
		x.CreateAttr("type", "ARRAY")
		x.CreateAttr("entryType", dts)
	} else {
		x.CreateAttr("type", dts)
	}
}
