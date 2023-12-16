package render

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func (r *renderer) renderDataTypes(cluster *matter.Cluster, clusters []*matter.Cluster, cx *etree.Element, errata *zap.Errata) {

	bitmaps := make(map[*matter.Bitmap]struct{})
	enums := make(map[*matter.Enum]struct{})
	structs := make(map[*matter.Struct]struct{})

	addTypes(cluster.Attributes, bitmaps, enums, structs)
	for _, cmd := range cluster.Commands {
		addTypes(cmd.Fields, bitmaps, enums, structs)
	}
	for _, e := range cluster.Events {
		addTypes(e.Fields, bitmaps, enums, structs)
	}
	r.renderBitmaps(bitmaps, cx)
	r.renderEnums(enums, cx)
	r.renderStructs(structs, cx)
}
func addTypes(fs matter.FieldSet, bitmaps map[*matter.Bitmap]struct{}, enums map[*matter.Enum]struct{}, structs map[*matter.Struct]struct{}) {
	for _, f := range fs {
		if f.Type == nil {
			continue
		}
		if conformance.IsZigbee(fs, f.Conformance) {
			continue
		}
		var model any
		if f.Type.IsArray() {
			if f.Type.EntryType != nil {
				model = f.Type.EntryType.Model
			}
		} else {
			model = f.Type.Model
		}
		switch model := model.(type) {
		case *matter.Bitmap:
			bitmaps[model] = struct{}{}
		case *matter.Enum:
			enums[model] = struct{}{}
		case *matter.Struct:
			structs[model] = struct{}{}
		}
	}
}

func writeAttributeDataType(x *etree.Element, fs matter.FieldSet, f *matter.Field) {
	if f.Type == nil {
		return
	}
	dts := zap.FieldToZapDataType(fs, f)
	if f.Type.IsArray() {
		x.CreateAttr("type", "array")
		x.CreateAttr("entryType", dts)
	} else {
		x.CreateAttr("type", dts)
	}
}

func writeDataType(x *etree.Element, fs matter.FieldSet, f *matter.Field) {
	if f.Type == nil {
		return
	}
	dts := zap.FieldToZapDataType(fs, f)
	if f.Type.IsArray() {
		x.CreateAttr("array", "true")
	}
	x.CreateAttr("type", dts)
}
