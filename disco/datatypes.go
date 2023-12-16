package disco

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

type DataTypeEntry struct {
	name                string
	ref                 string
	dataType            string
	dataTypeCategory    matter.DataTypeCategory
	section             *ascii.Section
	typeCell            *types.TableCell
	definitionTable     *types.Table
	indexColumn         matter.TableColumn
	definitionColumnMap ascii.ColumnIndex
	existing            bool
}

func getExistingDataTypes(cxt *discoContext, top *ascii.Section) {
	dataTypesSection := ascii.FindSectionByType(top, matter.SectionDataTypes)
	if dataTypesSection == nil {
		return
	}
	for _, ss := range parse.FindAll[*ascii.Section](dataTypesSection.Elements) {
		name := matter.StripDataTypeSuffixes(ss.Name)
		nameKey := strings.ToLower(name)
		dataType := ss.GetDataType()
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

func (b *Ball) getPotentialDataTypes(cxt *discoContext, section *ascii.Section, rows []*types.TableRow, columnMap ascii.ColumnIndex) error {
	sectionDataMap, err := b.getDataTypes(columnMap, rows, section)
	if err != nil {
		return err
	}
	for name, dataType := range sectionDataMap {
		if dataType.section != nil {
			cxt.potentialDataTypes[name] = append(cxt.potentialDataTypes[name], dataType)
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
		cv, err := ascii.GetTableCellValue(row.Cells[nameIndex])
		if err != nil {
			continue
		}
		dtv, err := ascii.GetTableCellValue(row.Cells[typeIndex])
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

func (b *Ball) promoteDataTypes(cxt *discoContext, top *ascii.Section) error {
	if !b.options.promoteDataTypes {
		return nil
	}

	fields := make(map[matter.DataTypeCategory]map[string]*DataTypeEntry)
	for _, infos := range cxt.potentialDataTypes {
		if len(infos) > 1 {
			disambiguateDataTypes(cxt, infos)
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
			err := b.promoteDataType(top, suffix, f, idColumn)
			if err != nil {
				return err
			}
		}
	}
	return nil
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

func (b *Ball) promoteDataType(top *ascii.Section, suffix string, dataTypeFields map[string]*DataTypeEntry, firstColumnType matter.TableColumn) error {
	if dataTypeFields == nil {
		return nil
	}
	var dataTypesSection *ascii.Section
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
		_, columnMap, _, err := ascii.MapTableColumns(b.doc, ascii.TableRows(table))
		if err != nil {
			return fmt.Errorf("failed mapping table columns for data type definition table in section %s: %w", dt.section.Name, err)
		}
		if valueIndex, ok := columnMap[firstColumnType]; !ok || valueIndex > 0 {
			continue
		}

		title, err := types.NewStringElement(dt.name + suffix)
		if err != nil {
			return err
		}

		if dataTypesSection == nil {
			dataTypesSection, err = ensureDataTypesSection(top)
			if err != nil {
				return err
			}
		}

		var removedTable bool
		parse.Filter(dt.section, func(i interface{}) (remove bool, shortCircuit bool) {
			if t, ok := i.(*types.Table); ok && table == t {
				removedTable = true
				return true, true
			}
			return false, false
		})

		if !removedTable {
			return fmt.Errorf("unable to relocate enum value table")
		}

		dataTypeSection, _ := types.NewSection(dataTypesSection.Base.Level+1, []interface{}{title})

		se, _ := types.NewStringElement(fmt.Sprintf("This data type is derived from %s", dt.dataType))
		p, _ := types.NewParagraph(nil, se)
		dataTypeSection.AddElement(p)
		bl, _ := types.NewBlankLine()
		dataTypeSection.AddElement(bl)
		dataTypeSection.AddElement(table)
		dataTypeSection.AddElement(bl)
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

		s, err := ascii.NewSection(dataTypesSection, dataTypeSection)

		if err != nil {
			return err
		}

		dataTypesSection.AppendSection(s)

		table.Attributes.Unset(types.AttrID)
		table.Attributes.Unset(types.AttrTitle)

		icr, _ := types.NewInternalCrossReference(newId, "")
		err = setCellValue(dt.typeCell, []interface{}{icr})
		if err != nil {
			return err
		}
	}
	return nil
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
	ts, _ := types.NewSection(top.Base.Level+1, []interface{}{title})
	bl, _ := types.NewBlankLine()
	ts.AddElement(bl)
	dataTypesSection, err = ascii.NewSection(top, ts)
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
	parents := make([]interface{}, len(infos))
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
				return fmt.Errorf("duplicate reference: %s with invalid parent", dataTypeNames[i])
			}
			parentSections[i] = parentSection
			refParentId := strings.TrimSpace(matter.StripReferenceSuffixes(ascii.ReferenceName(parentSection.Base)))
			dataTypeNames[i] = refParentId + " " + dataTypeNames[i]
			dataTypeRefs[i] = refParentId + "_" + dataTypeNames[i]
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
