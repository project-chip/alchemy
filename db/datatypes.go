package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (h *Host) indexDataTypeModels(cxt context.Context, parent *sectionInfo, cluster *matter.Cluster) error {
	for _, bm := range cluster.Bitmaps {
		row := newDBRow()
		row.values[matter.TableColumnName] = bm.Name
		row.values[matter.TableColumnClass] = "bitmap"
		bi := &sectionInfo{id: h.nextId(dataTypeTable), parent: parent, values: row, children: make(map[string][]*sectionInfo)}
		parent.children[dataTypeTable] = append(parent.children[dataTypeTable], bi)
		for _, bmv := range bm.Bits {
			bmr := newDBRow()
			bmr.values[matter.TableColumnBit] = bmv.Bit
			bmr.values[matter.TableColumnName] = bmv.Name
			bmr.values[matter.TableColumnSummary] = bmv.Summary
			bmr.values[matter.TableColumnConformance] = bmv.Conformance
			bv := &sectionInfo{id: h.nextId(bitmapValue), parent: bi, values: bmr}
			bi.children[bitmapValue] = append(bi.children[bitmapValue], bv)
		}
	}
	for _, en := range cluster.Enums {
		row := newDBRow()
		row.values[matter.TableColumnName] = en.Name
		row.values[matter.TableColumnClass] = "enum"
		row.values[matter.TableColumnType] = en.Type
		ei := &sectionInfo{id: h.nextId(dataTypeTable), parent: parent, values: row, children: make(map[string][]*sectionInfo)}
		parent.children[dataTypeTable] = append(parent.children[dataTypeTable], ei)
		for _, env := range en.Values {
			bmr := newDBRow()
			bmr.values[matter.TableColumnValue] = env.Value
			bmr.values[matter.TableColumnName] = env.Name
			bmr.values[matter.TableColumnSummary] = env.Summary
			bmr.values[matter.TableColumnConformance] = env.Conformance
			bv := &sectionInfo{id: h.nextId(enumValue), parent: ei, values: bmr}
			ei.children[enumValue] = append(ei.children[enumValue], bv)
		}
	}
	for _, s := range cluster.Structs {
		row := newDBRow()
		row.values[matter.TableColumnName] = s.Name
		row.values[matter.TableColumnClass] = "struct"
		ei := &sectionInfo{id: h.nextId(dataTypeTable), parent: parent, values: row, children: make(map[string][]*sectionInfo)}
		parent.children[dataTypeTable] = append(parent.children[dataTypeTable], ei)
		for _, env := range s.Fields {
			h.readField(env, ei, structField)
		}
	}
	return nil
}

func (h *Host) readField(f *matter.Field, parent *sectionInfo, tableName string) {
	sr := newDBRow()

	var t string
	if f.Type != nil {
		if f.Type.IsArray {
			t = fmt.Sprintf("list[%s]", f.Type.Name)
		} else {
			t = f.Type.Name
		}
	} else {
		t = "unknown"
	}
	sr.values[matter.TableColumnID] = f.ID
	sr.values[matter.TableColumnName] = f.Name
	sr.values[matter.TableColumnType] = t
	if f.Constraint != nil {
		sr.values[matter.TableColumnConstraint] = f.Constraint.AsciiDocString()
	} else {
		sr.values[matter.TableColumnConstraint] = ""
	}
	sr.values[matter.TableColumnQuality] = f.Quality.String()
	sr.values[matter.TableColumnDefault] = f.Default
	sr.values[matter.TableColumnAccess] = ascii.AccessToAsciiString(f.Access)
	sr.values[matter.TableColumnConformance] = f.Conformance
	sv := &sectionInfo{id: h.nextId(tableName), parent: parent, values: sr}
	parent.children[tableName] = append(parent.children[tableName], sv)
}

func (h *Host) indexDataTypes(cxt context.Context, ds *sectionInfo, dts *ascii.Section) error {
	if ds.children == nil {
		ds.children = make(map[string][]*sectionInfo)
	}
	for _, s := range parse.Skim[*ascii.Section](dts.Elements) {
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
				id:     h.nextId(dataTypeTable),
				parent: ds,
				values: &dbRow{
					values: map[matter.TableColumn]interface{}{
						matter.TableColumnClass: t,
						matter.TableColumnName:  name,
					},
				},
				children: make(map[string][]*sectionInfo),
			}
			ds.children[dataTypeTable] = append(ds.children[dataTypeTable], ci)
			switch s.SecType {
			case matter.SectionDataTypeBitmap:
				h.readTableSection(cxt, ci, s, bitmapValue)
			case matter.SectionDataTypeEnum:
				h.readTableSection(cxt, ci, s, enumValue)
			case matter.SectionDataTypeStruct:
				h.readTableSection(cxt, ci, s, structField)
			}
		}
	}
	return nil
}
