package render

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func (r *renderer) renderDataTypes(cluster *matter.Cluster, clusters []*matter.Cluster, cx *etree.Element, errata *zap.Errata) {

	r.renderBitmaps(r.configurator.Bitmaps, cx)
	r.renderEnums(r.configurator.Enums, cx)
	r.renderStructs(r.configurator.Structs, cx)
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
