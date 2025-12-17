package db

import (
	"context"
	"fmt"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (h *Host) indexDataTypeModels(cxt context.Context, parent *sectionInfo, cluster *matter.Cluster) error {
	h.indexBitmaps(cluster, parent)
	h.indexEnums(cluster, parent)
	h.indexStructs(cluster, parent)
	h.indexTypeDefs(cluster, parent)
	return nil
}

func (h *Host) indexBitmaps(cluster *matter.Cluster, parent *sectionInfo) {
	for _, bm := range cluster.Bitmaps {
		row := newDBRow()
		row.values[matter.TableColumnName] = bm.Name
		row.values[matter.TableColumnClass] = "bitmap"
		bi := h.newSectionInfo(bitmapTable, parent, row, bm)
		parent.children[bitmapTable] = append(parent.children[bitmapTable], bi)
		for _, bmv := range bm.Bits {
			bmr := newDBRow()
			bmr.values[matter.TableColumnBit] = bmv
			bmr.values[matter.TableColumnName] = bmv.Name()
			bmr.values[matter.TableColumnSummary] = bmv.Summary()
			bmr.values[matter.TableColumnConformance] = bmv.Conformance().ASCIIDocString()
			bv := h.newSectionInfo(bitmapValue, bi, bmr, bmv)
			bi.children[bitmapValue] = append(bi.children[bitmapValue], bv)
		}
	}
}

func (h *Host) indexEnums(cluster *matter.Cluster, parent *sectionInfo) {
	for _, en := range cluster.Enums {
		row := newDBRow()
		row.values[matter.TableColumnName] = en.Name
		row.values[matter.TableColumnType] = en.Type.Name
		ei := h.newSectionInfo(enumTable, parent, row, en)
		parent.children[enumTable] = append(parent.children[enumTable], ei)
		for _, env := range en.Values {
			en := newDBRow()
			en.values[matter.TableColumnValue] = env.Value
			en.values[matter.TableColumnName] = env.Name
			en.values[matter.TableColumnSummary] = env.Summary
			en.values[matter.TableColumnConformance] = env.Conformance.ASCIIDocString()
			ev := h.newSectionInfo(enumValue, ei, en, env)
			ei.children[enumValue] = append(ei.children[enumValue], ev)
		}
	}
}
func (h *Host) indexTypeDefs(cluster *matter.Cluster, parent *sectionInfo) {
	for _, t := range cluster.TypeDefs {
		row := newDBRow()
		row.values[matter.TableColumnName] = t.Name
		row.values[matter.TableColumnType] = t.Type.Name
		row.values[matter.TableColumnDescription] = t.Description
		ei := h.newSectionInfo(typedefTable, parent, row, t)
		parent.children[typedefTable] = append(parent.children[typedefTable], ei)
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
	sr.values[matter.TableColumnFallback] = f.Fallback
	sr.values[matter.TableColumnAccess] = spec.AccessToASCIIDocString(f.Access, entityType)
	if f.Conformance != nil {
		sr.values[matter.TableColumnConformance] = f.Conformance.ASCIIDocString()
	}
	sv := h.newSectionInfo(tableName, parent, sr, f)
	parent.children[tableName] = append(parent.children[tableName], sv)
}
