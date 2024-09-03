package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type DataTypeEntry struct {
	name             string
	ref              string
	dataType         string
	dataTypeCategory matter.DataTypeCategory
	section          *spec.Section
	typeCell         *asciidoc.TableCell
	definitionTable  *asciidoc.Table
	indexColumn      matter.TableColumn
	existing         bool
}

func getExistingDataTypes(cxt *discoContext, dp *docParse) {
	if dp.dataTypes == nil {
		return
	}

	for _, ss := range parse.FindAll[*spec.Section](dp.dataTypes.section.Elements()) {
		name := matter.StripDataTypeSuffixes(ss.Name)
		nameKey := strings.ToLower(name)
		dataType := ss.GetDataType()
		if dataType == nil {
			slog.Debug("failed to find data type", "sectionName", ss.Name)
			continue
		}
		dataTypeCategory := getDataTypeCategory(dataType.Name)
		cxt.potentialDataTypes[nameKey] = append(cxt.potentialDataTypes[nameKey], &DataTypeEntry{
			name:             name,
			ref:              name,
			section:          ss,
			dataType:         dataType.Name,
			dataTypeCategory: dataTypeCategory,
			existing:         true,
			indexColumn:      getIndexColumnType(dataTypeCategory),
		})
	}
}

func (b *Ball) getPotentialDataTypes(dc *discoContext, dp *docParse) (err error) {
	var subSections []*subSection
	subSections = append(subSections, dp.attributes...)
	subSections = append(subSections, dp.structs...)
	subSections = append(subSections, dp.commands...)
	subSections = append(subSections, dp.events...)

	for _, ss := range subSections {
		err = b.getPotentialDataTypesForSection(dc, dp, ss)
		if err != nil {
			return
		}
	}
	return
}

func (b *Ball) getPotentialDataTypesForSection(cxt *discoContext, dp *docParse, ss *subSection) error {
	if ss.table.element == nil {
		slog.Debug("no data type table found", "sectionName", ss.section.Name)
		return nil
	}
	sectionDataMap, err := b.getDataTypes(ss.table.columnMap, ss.table.rows, ss.section)
	if err != nil {
		return err
	}
	for name, dataType := range sectionDataMap {
		if dataType.section != nil {
			cxt.potentialDataTypes[name] = append(cxt.potentialDataTypes[name], dataType)
		}
	}
	for _, child := range ss.children {
		err = b.getPotentialDataTypesForSection(cxt, dp, child)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Ball) getDataTypes(columnMap spec.ColumnIndex, rows []*asciidoc.TableRow, section *spec.Section) (map[string]*DataTypeEntry, error) {
	sectionDataMap := make(map[string]*DataTypeEntry)
	nameIndex, ok := columnMap[matter.TableColumnName]
	if !ok {
		return nil, nil
	}
	typeIndex, ok := columnMap[matter.TableColumnType]
	if !ok {
		return nil, nil
	}
	for _, row := range rows {
		cv, err := spec.RenderTableCell(row.Cell(nameIndex))
		if err != nil {
			continue
		}
		dtv, err := spec.RenderTableCell(row.Cell(typeIndex))
		if err != nil {
			continue
		}
		name := strings.TrimSpace(cv)
		nameKey := strings.ToLower(name)

		dataType := strings.TrimSpace(dtv)
		dataTypeCategory := getDataTypeCategory(dataType)

		if dataTypeCategory == matter.DataTypeCategoryUnknown {
			continue
		}

		if _, ok := sectionDataMap[nameKey]; !ok {
			sectionDataMap[nameKey] = &DataTypeEntry{
				name:             name,
				ref:              name,
				dataType:         dataType,
				dataTypeCategory: dataTypeCategory,
				typeCell:         row.Cell(typeIndex),
			}
		}
	}
	for _, el := range section.Elements() {
		if s, ok := el.(*spec.Section); ok {
			name := strings.TrimSpace(matter.StripReferenceSuffixes(s.Name))

			dataType, ok := sectionDataMap[strings.ToLower(name)]
			if !ok {
				continue
			}
			table := spec.FindFirstTable(s)
			if table == nil {
				continue
			}
			_, columnMap, _, err := spec.MapTableColumns(b.doc, table.TableRows())
			if err != nil {
				return nil, fmt.Errorf("failed mapping table columns for data type definition table in section %s: %w", s.Name, err)
			}
			dataType.indexColumn = getIndexColumnType(dataType.dataTypeCategory)

			if valueIndex, ok := columnMap[dataType.indexColumn]; !ok || valueIndex > 0 {
				continue
			}
			dataType.section = s
			dataType.definitionTable = table

		}
	}
	return sectionDataMap, nil
}

func (b *Ball) promoteDataTypes(cxt *discoContext, top *spec.Section) (promoted bool, err error) {
	if !b.options.promoteDataTypes {
		return
	}

	fields := make(map[matter.DataTypeCategory]map[string]*DataTypeEntry)
	for _, infos := range cxt.potentialDataTypes {
		if len(infos) > 1 {
			err = disambiguateDataTypes(cxt, infos)
			if err != nil {
				return
			}
		}
		for _, info := range infos {
			fieldMap, ok := fields[info.dataTypeCategory]
			if !ok {
				fieldMap = make(map[string]*DataTypeEntry)
				fields[info.dataTypeCategory] = fieldMap
			}
			fieldMap[info.name] = info
		}
	}

	if len(fields) > 0 {
		for _, dtc := range matter.DataTypeOrder {
			f, ok := fields[dtc]
			if !ok {
				continue
			}
			suffix := matter.DataTypeSuffixes[dtc]
			idColumn := matter.DataTypeIdentityColumn[dtc]
			var didPromotion bool
			didPromotion, err = b.promoteDataType(top, suffix, f, idColumn, dtc)
			if err != nil {
				return
			}
			promoted = didPromotion || promoted
		}
	}
	return
}

func getIndexColumnType(dataTypeCategory matter.DataTypeCategory) matter.TableColumn {
	switch dataTypeCategory {
	case matter.DataTypeCategoryEnum:
		return matter.TableColumnValue
	case matter.DataTypeCategoryBitmap:
		return matter.TableColumnBit
	}
	return matter.TableColumnUnknown
}

func getDataTypeCategory(dataType string) matter.DataTypeCategory {
	switch dataType {
	case "enum8", "enum16", "enum32":
		return matter.DataTypeCategoryEnum
	case "map8", "map16", "map32":
		return matter.DataTypeCategoryBitmap
	}
	return matter.DataTypeCategoryUnknown
}

func (b *Ball) promoteDataType(top *spec.Section, suffix string, dataTypeFields map[string]*DataTypeEntry, firstColumnType matter.TableColumn, dtc matter.DataTypeCategory) (promoted bool, err error) {
	if dataTypeFields == nil {
		return
	}
	var dataTypesSection *spec.Section
	var entityType types.EntityType
	switch dtc {
	case matter.DataTypeCategoryBitmap:
		entityType = types.EntityTypeBitmapValue
	case matter.DataTypeCategoryEnum:
		entityType = types.EntityTypeEnumValue
	case matter.DataTypeCategoryStruct:
		entityType = types.EntityTypeField
	}
	for _, dt := range dataTypeFields {
		if dt.existing {
			continue
		}

		if dt.section == nil {
			continue
		}
		table := spec.FindFirstTable(dt.section)
		if table == nil {
			continue
		}
		ti := tableInfo{element: table, rows: table.TableRows()}
		ti.headerRow, ti.columnMap, ti.extraColumns, err = spec.MapTableColumns(b.doc, ti.rows)
		if err != nil {
			err = fmt.Errorf("failed mapping table columns for data type definition table in section %s: %w", dt.section.Name, err)
			return
		}
		if valueIndex, ok := ti.columnMap[firstColumnType]; !ok || valueIndex > 0 {
			continue
		}

		summaryIndex, hasSummaryColumn := ti.columnMap[matter.TableColumnSummary]
		if !hasSummaryColumn {
			descriptionIndex, hasDescriptionColumn := ti.columnMap[matter.TableColumnDescription]
			if hasDescriptionColumn {
				// Use the description column as the summary
				delete(ti.columnMap, matter.TableColumnDescription)
				ti.columnMap[matter.TableColumnSummary] = descriptionIndex
				summaryIndex = descriptionIndex
				err = b.renameTableHeaderCells(top.Doc, dt.section, &ti, nil)
				if err != nil {
					return
				}
			} else if len(ti.extraColumns) > 0 {
				// Hrm, no summary or description on this promoted data type table
				// Take the first extra column and rename it
				summaryIndex = ti.extraColumns[0].Offset
				ti.columnMap[matter.TableColumnSummary] = summaryIndex
				err = b.renameTableHeaderCells(top.Doc, dt.section, &ti, nil)
				if err != nil {
					return
				}
			} else {
				summaryIndex, err = b.appendColumn(table, ti.columnMap, ti.headerRow, matter.TableColumnSummary, entityType)
				if err != nil {
					return
				}
			}
		}
		_, hasNameColumn := ti.columnMap[matter.TableColumnName]
		if !hasNameColumn {
			var nameIndex int
			nameIndex, err = b.appendColumn(table, ti.columnMap, ti.headerRow, matter.TableColumnName, entityType)
			if err != nil {
				return
			}
			err = copyCells(ti.rows, ti.headerRow, summaryIndex, nameIndex, matter.Case)
			if err != nil {
				return
			}
		}

		dataTypeName := spec.CanonicalName(dt.name + suffix)

		title := asciidoc.NewString(dataTypeName + " Type")

		if dataTypesSection == nil {
			dataTypesSection, err = ensureDataTypesSection(top)
			if err != nil {
				return
			}
		}

		var removedTable bool
		parse.Filter(dt.section, func(i any) (remove bool, shortCircuit bool) {
			if t, ok := i.(*asciidoc.Table); ok && table == t {
				removedTable = true
				return true, true
			}
			return false, false
		})

		if !removedTable {
			err = fmt.Errorf("unable to relocate enum value table")
			return
		}

		dataTypeSection := asciidoc.NewSection(asciidoc.Set{title}, dataTypesSection.Base.Level+1)

		se := asciidoc.NewString(fmt.Sprintf("This data type is derived from %s", dt.dataType))
		p := asciidoc.NewParagraph()
		p.SetElements(asciidoc.Set{se})
		err = dataTypeSection.Append(p)
		if err != nil {
			return
		}
		bl := asciidoc.NewEmptyLine("")
		err = dataTypeSection.Append(bl)
		if err != nil {
			return
		}
		err = dataTypeSection.Append(table)
		if err != nil {
			return
		}
		err = dataTypeSection.Append(bl)
		if err != nil {
			return
		}

		//newAttr := make(asciidoc.AttributeList)
		tableIDAttribute := table.GetAttributeByName(asciidoc.AttributeNameID)
		var newID string
		if tableIDAttribute != nil {
			// Reuse the table's ID if it has one, so existing links get updated
			newID = tableIDAttribute.Raw()
		} else {
			newID = "ref_" + dt.ref + suffix + ", " + dt.name + suffix
		}

		dataTypeSection.AppendAttribute(asciidoc.NewNamedAttribute(string(asciidoc.AttributeNameID), asciidoc.Set{asciidoc.NewString(newID)}, asciidoc.AttributeQuoteTypeDouble))

		var s *spec.Section
		s, err = spec.NewSection(top.Doc, dataTypesSection, dataTypeSection)

		if err != nil {
			return
		}
		switch dt.dataTypeCategory {
		case matter.DataTypeCategoryBitmap:
			s.SecType = matter.SectionDataTypeBitmap
		case matter.DataTypeCategoryEnum:
			s.SecType = matter.SectionDataTypeEnum
		}

		err = dataTypesSection.AppendSection(s)
		if err != nil {
			return
		}

		table.DeleteAttribute(asciidoc.AttributeNameID)
		table.DeleteAttribute(asciidoc.AttributeNameTitle)

		icr := asciidoc.NewCrossReference(newID)
		err = setCellValue(dt.typeCell, asciidoc.Set{icr})
		if err != nil {
			return
		}
		promoted = true
	}
	return
}

func ensureDataTypesSection(top *spec.Section) (*spec.Section, error) {
	dataTypesSection := spec.FindSectionByType(top, matter.SectionDataTypes)
	if dataTypesSection != nil {
		return dataTypesSection, nil
	}
	title := asciidoc.NewString(matter.SectionTypeName(matter.SectionDataTypes))

	ts := asciidoc.NewSection(asciidoc.Set{title}, top.Base.Level+1)
	err := ts.Append(asciidoc.NewEmptyLine(""))
	if err != nil {
		return nil, err
	}
	dataTypesSection, err = spec.NewSection(top.Doc, top, ts)
	if err != nil {
		return nil, err
	}
	dataTypesSection.SecType = matter.SectionDataTypes
	err = top.AppendSection(dataTypesSection)
	if err != nil {
		return nil, err
	}
	return dataTypesSection, nil
}

func disambiguateDataTypes(cxt *discoContext, infos []*DataTypeEntry) error {
	parents := make([]any, len(infos))
	dataTypeNames := make([]string, len(infos))
	dataTypeRefs := make([]string, len(infos))
	for i, info := range infos {
		parents[i] = info.section
		dataTypeNames[i] = info.name
		dataTypeRefs[i] = info.ref
	}
	parentSections := make([]*spec.Section, len(infos))
	for {
		for i := range infos {
			parentSection := findRefSection(parents[i])
			if parentSection == nil {
				return fmt.Errorf("duplicate reference: %s in %T with invalid parent", dataTypeNames[i], parents[i])
			}
			parentSections[i] = parentSection
			refParentID := strings.TrimSpace(matter.StripReferenceSuffixes(spec.ReferenceName(parentSection.Base)))
			dataTypeNames[i] = refParentID + dataTypeNames[i]
			dataTypeRefs[i] = refParentID + dataTypeNames[i]
		}
		ids := make(map[string]struct{})
		var duplicateIds bool
		for _, refID := range dataTypeRefs {
			if _, ok := ids[refID]; ok {
				duplicateIds = true
			}
			ids[refID] = struct{}{}
		}
		if duplicateIds {
			for i, info := range infos {
				parents[i] = parentSections[i].Parent
				dataTypeNames[i] = info.name
				dataTypeRefs[i] = info.ref
			}
		} else {
			break
		}
	}
	for i, info := range infos {
		info.name = dataTypeNames[i]
		info.ref = dataTypeRefs[i]
	}
	return nil
}

func (b *Ball) canonicalizeDataTypeSectionName(dp *docParse, s *spec.Section, dataTypeName string) {
	name := s.Name
	if text.HasCaseInsensitiveSuffix(name, dataTypeName+" type") {
		return
	}
	var newName = name
	if text.HasCaseInsensitiveSuffix(name, dataTypeName) {
		newName = spec.CanonicalName(name) + " Type"
	} else if text.HasCaseInsensitiveSuffix(name, " type") {
		newName = spec.CanonicalName(name[:len(name)-len(" type")])
		newName += dataTypeName + " Type"
	} else {
		newName = spec.CanonicalName(name) + dataTypeName + " Type"
	}
	if name == newName {
		return
	}
	setSectionTitle(s, newName)
	oldName := text.TrimCaseInsensitiveSuffix(name, " type")
	newName = text.TrimCaseInsensitiveSuffix(newName, " type")
	if oldName == newName {
		return
	}
	renameDataType(dp.attributes, oldName, newName)
	renameDataType(dp.commands, oldName, newName)
	renameDataType(dp.events, oldName, newName)
}

func renameDataType(subSections []*subSection, oldName string, newName string) {
	slog.Warn("renaming", "old", oldName, "new", newName)
	for _, ss := range subSections {
		renameDataType(ss.children, oldName, newName)
		if ss.table.element == nil {
			continue
		}
		typeIndex, ok := ss.table.columnMap[matter.TableColumnType]
		if !ok {
			continue
		}
		for i, row := range ss.table.rows {
			if i == ss.table.headerRow {
				continue
			}
			typeCell := row.Cell(typeIndex)
			vc, e := spec.RenderTableCell(typeCell)
			if e != nil {
				continue
			}
			slog.Warn("renaming cell", "replacement", newName, "old", vc, "looking for", oldName)

			if strings.EqualFold(oldName, strings.TrimSpace(vc)) {
				setCellString(typeCell, newName)
			}
		}
	}
}
