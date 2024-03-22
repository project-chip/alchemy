package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

type DataTypeEntry struct {
	name             string
	ref              string
	dataType         string
	dataTypeCategory matter.DataTypeCategory
	section          *ascii.Section
	typeCell         *types.TableCell
	definitionTable  *types.Table
	indexColumn      matter.TableColumn
	existing         bool
}

func getExistingDataTypes(cxt *discoContext, dp *docParse) {
	if dp.dataTypes == nil {
		return
	}

	for _, ss := range parse.FindAll[*ascii.Section](dp.dataTypes.section.Elements) {
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

func (b *Ball) getDataTypes(columnMap ascii.ColumnIndex, rows []*types.TableRow, section *ascii.Section) (map[string]*DataTypeEntry, error) {
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
		cv, err := ascii.RenderTableCell(row.Cells[nameIndex])
		if err != nil {
			continue
		}
		dtv, err := ascii.RenderTableCell(row.Cells[typeIndex])
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
				typeCell:         row.Cells[typeIndex],
			}
		}
	}
	for _, el := range section.Elements {
		if s, ok := el.(*ascii.Section); ok {
			name := strings.TrimSpace(matter.StripReferenceSuffixes(s.Name))

			dataType, ok := sectionDataMap[strings.ToLower(name)]
			if !ok {
				continue
			}
			table := ascii.FindFirstTable(s)
			if table == nil {
				continue
			}
			_, columnMap, _, err := ascii.MapTableColumns(b.doc, ascii.TableRows(table))
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

func (b *Ball) promoteDataTypes(cxt *discoContext, top *ascii.Section) (promoted bool, err error) {
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

func (b *Ball) promoteDataType(top *ascii.Section, suffix string, dataTypeFields map[string]*DataTypeEntry, firstColumnType matter.TableColumn, dtc matter.DataTypeCategory) (promoted bool, err error) {
	if dataTypeFields == nil {
		return
	}
	var dataTypesSection *ascii.Section
	var entityType mattertypes.EntityType
	switch dtc {
	case matter.DataTypeCategoryBitmap:
		entityType = mattertypes.EntityTypeBitmapValue
	case matter.DataTypeCategoryEnum:
		entityType = mattertypes.EntityTypeEnumValue
	case matter.DataTypeCategoryStruct:
		entityType = mattertypes.EntityTypeField
	}
	for _, dt := range dataTypeFields {
		if dt.existing {
			continue
		}

		if dt.section == nil {
			continue
		}
		table := ascii.FindFirstTable(dt.section)
		if table == nil {
			continue
		}
		rows := ascii.TableRows(table)
		var headerRowIndex int
		var columnMap ascii.ColumnIndex
		var extraColumns []ascii.ExtraColumn
		headerRowIndex, columnMap, extraColumns, err = ascii.MapTableColumns(b.doc, rows)
		if err != nil {
			err = fmt.Errorf("failed mapping table columns for data type definition table in section %s: %w", dt.section.Name, err)
			return
		}
		if valueIndex, ok := columnMap[firstColumnType]; !ok || valueIndex > 0 {
			continue
		}

		summaryIndex, hasSummaryColumn := columnMap[matter.TableColumnSummary]
		if !hasSummaryColumn {
			descriptionIndex, hasDescriptionColumn := columnMap[matter.TableColumnDescription]
			if hasDescriptionColumn {
				// Use the description column as the summary
				delete(columnMap, matter.TableColumnDescription)
				columnMap[matter.TableColumnSummary] = descriptionIndex
				summaryIndex = descriptionIndex
				err = b.renameTableHeaderCells(rows, headerRowIndex, columnMap, nil)
				if err != nil {
					return
				}
			} else if len(extraColumns) > 0 {
				// Hrm, no summary or description on this promoted data type table
				// Take the first extra column and rename it
				summaryIndex = extraColumns[0].Offset
				columnMap[matter.TableColumnSummary] = summaryIndex
				err = b.renameTableHeaderCells(rows, headerRowIndex, columnMap, nil)
				if err != nil {
					return
				}
			} else {
				summaryIndex, err = b.appendColumn(rows, columnMap, headerRowIndex, matter.TableColumnSummary, nil, entityType)
				if err != nil {
					return
				}
			}
		}
		_, hasNameColumn := columnMap[matter.TableColumnName]
		if !hasNameColumn {
			var nameIndex int
			nameIndex, err = b.appendColumn(rows, columnMap, headerRowIndex, matter.TableColumnName, nil, entityType)
			if err != nil {
				return
			}
			err = copyCells(rows, headerRowIndex, summaryIndex, nameIndex, matter.Case)
			if err != nil {
				return
			}
		}

		var title *types.StringElement
		title, err = types.NewStringElement(dt.name + suffix + " Type")
		if err != nil {
			return
		}

		if dataTypesSection == nil {
			dataTypesSection, err = ensureDataTypesSection(top)
			if err != nil {
				return
			}
		}

		var removedTable bool
		parse.Filter(dt.section, func(i any) (remove bool, shortCircuit bool) {
			if t, ok := i.(*types.Table); ok && table == t {
				removedTable = true
				return true, true
			}
			return false, false
		})

		if !removedTable {
			err = fmt.Errorf("unable to relocate enum value table")
			return
		}

		dataTypeSection, _ := types.NewSection(dataTypesSection.Base.Level+1, []any{title})

		se, _ := types.NewStringElement(fmt.Sprintf("This data type is derived from %s", dt.dataType))
		p, _ := types.NewParagraph(nil, se)
		err = dataTypeSection.AddElement(p)
		if err != nil {
			return
		}
		bl, _ := types.NewBlankLine()
		err = dataTypeSection.AddElement(bl)
		if err != nil {
			return
		}
		err = dataTypeSection.AddElement(table)
		if err != nil {
			return
		}
		err = dataTypeSection.AddElement(bl)
		if err != nil {
			return
		}
		newAttr := make(types.Attributes)
		var newId string
		if id, ok := table.Attributes.GetAsString(types.AttrID); ok {
			// Reuse the table's ID if it has one, so existing links get updated
			newId = id
			newAttr[types.AttrID] = newId
		} else {
			newId = "ref_" + dt.ref + suffix
			newAttr[types.AttrID] = newId + ", " + dt.name + suffix
		}

		dataTypeSection.AddAttributes(newAttr)

		var s *ascii.Section
		s, err = ascii.NewSection(top.Doc, dataTypesSection, dataTypeSection)

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

		table.Attributes.Unset(types.AttrID)
		table.Attributes.Unset(types.AttrTitle)

		icr, _ := types.NewInternalCrossReference(newId, "")
		err = setCellValue(dt.typeCell, []any{icr})
		if err != nil {
			return
		}
		promoted = true
	}
	return
}

func ensureDataTypesSection(top *ascii.Section) (*ascii.Section, error) {
	dataTypesSection := ascii.FindSectionByType(top, matter.SectionDataTypes)
	if dataTypesSection != nil {
		return dataTypesSection, nil
	}
	title, err := types.NewStringElement(matter.SectionTypeName(matter.SectionDataTypes))
	if err != nil {
		return nil, err
	}
	ts, _ := types.NewSection(top.Base.Level+1, []any{title})
	bl, _ := types.NewBlankLine()
	err = ts.AddElement(bl)
	if err != nil {
		return nil, err
	}
	dataTypesSection, err = ascii.NewSection(top.Doc, top, ts)
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
	parentSections := make([]*ascii.Section, len(infos))
	for {
		for i := range infos {
			parentSection := findRefSection(parents[i])
			if parentSection == nil {
				return fmt.Errorf("duplicate reference: %s in %T with invalid parent", dataTypeNames[i], parents[i])
			}
			parentSections[i] = parentSection
			refParentId := strings.TrimSpace(matter.StripReferenceSuffixes(ascii.ReferenceName(parentSection.Base)))
			dataTypeNames[i] = refParentId + dataTypeNames[i]
			dataTypeRefs[i] = refParentId + dataTypeNames[i]
		}
		ids := make(map[string]struct{})
		var duplicateIds bool
		for _, refId := range dataTypeRefs {
			if _, ok := ids[refId]; ok {
				duplicateIds = true
			}
			ids[refId] = struct{}{}
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
