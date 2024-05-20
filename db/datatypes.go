package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
)

func (h *Host) indexDataTypeModels(cxt context.Context, parent *sectionInfo, cluster *matter.Cluster) error {
	h.indexBitmaps(cluster, parent)
	h.indexEnums(cluster, parent)
	h.indexStructs(cluster, parent)
	return nil
}

func (h *Host) indexBitmaps(cluster *matter.Cluster, parent *sectionInfo) {
	for _, bm := range cluster.Bitmaps {
		row := newDBRow()
		row.values[matter.TableColumnName] = bm.Name
		row.values[matter.TableColumnClass] = "bitmap"
		bi := &sectionInfo{id: h.nextID(bitmapTable), parent: parent, values: row, children: make(map[string][]*sectionInfo)}
		parent.children[bitmapTable] = append(parent.children[bitmapTable], bi)
		for _, bmv := range bm.Bits {
			bmr := newDBRow()
			bmr.values[matter.TableColumnBit] = bmv
			bmr.values[matter.TableColumnName] = bmv.Name()
			bmr.values[matter.TableColumnSummary] = bmv.Summary()
			bmr.values[matter.TableColumnConformance] = bmv.Conformance().ASCIIDocString()
			bv := &sectionInfo{id: h.nextID(bitmapValue), parent: bi, values: bmr}
			bi.children[bitmapValue] = append(bi.children[bitmapValue], bv)
		}
	}
}

func (h *Host) indexEnums(cluster *matter.Cluster, parent *sectionInfo) {
	for _, en := range cluster.Enums {
		row := newDBRow()
		row.values[matter.TableColumnName] = en.Name
		row.values[matter.TableColumnType] = en.Type.Name
		ei := &sectionInfo{id: h.nextID(enumTable), parent: parent, values: row, children: make(map[string][]*sectionInfo)}
		parent.children[enumTable] = append(parent.children[enumTable], ei)
		for _, env := range en.Values {
			bmr := newDBRow()
			bmr.values[matter.TableColumnValue] = env.Value
			bmr.values[matter.TableColumnName] = env.Name
			bmr.values[matter.TableColumnSummary] = env.Summary
			bmr.values[matter.TableColumnConformance] = env.Conformance.ASCIIDocString()
			bv := &sectionInfo{id: h.nextID(enumValue), parent: ei, values: bmr}
			ei.children[enumValue] = append(ei.children[enumValue], bv)
		}
	}
}

func (h *Host) readField(f *matter.Field, parent *sectionInfo, tableName string, entityType types.EntityType) {
	sr := newDBRow()

	var t string
	if f.Type != nil {
		if f.Type.IsArray() {
			t = fmt.Sprintf("list[%s]", f.Type.EntryType.Name)
		} else {
			t = f.Type.Name
		}
	} else {
		t = "unknown"
	}
	sr.values[matter.TableColumnID] = f.ID.IntString()
	sr.values[matter.TableColumnName] = f.Name
	sr.values[matter.TableColumnType] = t
	if f.Constraint != nil {
		sr.values[matter.TableColumnConstraint] = f.Constraint.ASCIIDocString(f.Type)
	} else {
		sr.values[matter.TableColumnConstraint] = ""
	}
	sr.values[matter.TableColumnQuality] = f.Quality.String()
	sr.values[matter.TableColumnDefault] = f.Default
	sr.values[matter.TableColumnAccess] = ascii.AccessToASCIIDocString(f.Access, entityType)
	if f.Conformance != nil {
		sr.values[matter.TableColumnConformance] = f.Conformance.ASCIIDocString()
	}
	sv := &sectionInfo{id: h.nextID(tableName), parent: parent, values: sr}
	parent.children[tableName] = append(parent.children[tableName], sv)
}

func (h *Host) indexDataTypes(cxt context.Context, doc *ascii.Doc, ds *sectionInfo, dts *ascii.Section) (err error) {
	if ds.children == nil {
		ds.children = make(map[string][]*sectionInfo)
	}
	for _, s := range parse.Skim[*ascii.Section](dts.Elements()) {
		switch s.SecType {
		case matter.SectionDataTypeBitmap, matter.SectionDataTypeEnum, matter.SectionDataTypeStruct:
			var t string
			switch s.SecType {
			case matter.SectionDataTypeBitmap:
				t = "bitmap"
			case matter.SectionDataTypeEnum:
				t = "enum"
			case matter.SectionDataTypeStruct:
				t = "struct"
			}
			name := strings.TrimSuffix(s.Name, " Type")
			name = matter.StripDataTypeSuffixes(name)
			ci := &sectionInfo{
				parent: ds,
				values: &dbRow{
					values: map[matter.TableColumn]any{
						matter.TableColumnClass: t,
						matter.TableColumnName:  name,
					},
				},
				children: make(map[string][]*sectionInfo),
			}
			switch s.SecType {
			case matter.SectionDataTypeBitmap:
				ci.id = h.nextID(bitmapTable)
				err = h.readTableSection(cxt, doc, ci, s, bitmapValue)
				ds.children[bitmapTable] = append(ds.children[bitmapTable], ci)
			case matter.SectionDataTypeEnum:
				ci.id = h.nextID(enumTable)
				err = h.readTableSection(cxt, doc, ci, s, enumValue)
				ds.children[enumTable] = append(ds.children[enumTable], ci)
			case matter.SectionDataTypeStruct:
				ci.id = h.nextID(structTable)
				err = h.readTableSection(cxt, doc, ci, s, structField)
				ds.children[structTable] = append(ds.children[structTable], ci)
			}
			if err != nil {
				return
			}
		}
	}
	return nil
}
