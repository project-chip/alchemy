package disco

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

type potentialDataType struct {
	name                string
	ref                 string
	dataType            string
	dataTypeCategory    matter.DataTypeCategory
	section             *ascii.Section
	typeCell            *types.TableCell
	definitionTable     *types.Table
	indexColumn         matter.TableColumn
	definitionColumnMap map[matter.TableColumn]int
	existing            bool
}

var dataTypeDefinitionPattern = regexp.MustCompile(`is\s+derived\s+from\s+(?:<<enum-def\s*,\s*)?(enum8|enum16|enum32|map8|map16|map32)(?:\s*>>)?`)

func getExistingDataTypes(cxt *Context, top *ascii.Section) {
	dataTypesSection := findSectionByType(top, matter.SectionDataTypes)
	if dataTypesSection == nil {
		return
	}
	for _, ss := range ascii.FindAll[*ascii.Section](dataTypesSection.Elements) {
		name := stripDataTypeSuffixes(ss.Name)
		nameKey := strings.ToLower(name)
		var dataType string
		se := ascii.FindFirst[*types.StringElement](ss.Elements)
		if se != nil {
			match := dataTypeDefinitionPattern.FindStringSubmatch(se.Content)
			if match != nil {
				dataType = match[1]
			}
		}
		dataTypeCategory := getDataTypeCategory(dataType)
		cxt.potentialDataTypes[nameKey] = append(cxt.potentialDataTypes[nameKey], &potentialDataType{
			name:             name,
			ref:              name,
			section:          ss,
			dataType:         dataType,
			dataTypeCategory: dataTypeCategory,
			existing:         true,
			indexColumn:      getIndexColumnType(dataTypeCategory),
		})
	}
}

func getPotentialDataTypes(cxt *Context, section *ascii.Section, rows []*types.TableRow, columnMap map[matter.TableColumn]int) error {
	sectionDataMap := make(map[string]*potentialDataType)
	nameIndex, ok := columnMap[matter.TableColumnName]
	if !ok {
		return nil
	}
	typeIndex, ok := columnMap[matter.TableColumnType]
	if !ok {
		return nil
	}
	for _, row := range rows {
		cv, err := getCellValue(row.Cells[nameIndex])
		if err != nil {
			continue
		}
		dtv, err := getCellValue(row.Cells[typeIndex])
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
			sectionDataMap[nameKey] = &potentialDataType{
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
			name := strings.TrimSpace(stripReferenceSuffixes(s.Name))

			dataType, ok := sectionDataMap[strings.ToLower(name)]
			if !ok {
				continue
			}
			table := findFirstTable(s)
			if table == nil {
				continue
			}
			_, columnMap, _, err := findColumns(combineRows(table))
			if err != nil {
				return err
			}
			dataType.indexColumn = getIndexColumnType(dataType.dataTypeCategory)

			if valueIndex, ok := columnMap[dataType.indexColumn]; !ok || valueIndex > 0 {
				continue
			}
			dataType.section = s
			dataType.definitionTable = table

		}
	}
	for name, dataType := range sectionDataMap {
		if dataType.section != nil {
			cxt.potentialDataTypes[name] = append(cxt.potentialDataTypes[name], dataType)
		}
	}
	return nil
}

func promoteDataTypes(cxt *Context, top *ascii.Section) error {

	fields := make(map[matter.DataTypeCategory]map[string]*potentialDataType)
	//var dataTypeCount int
	//enumFields := make(map[string]*potentialDataType)
	//bitmapFields := make(map[string]*potentialDataType)
	for _, infos := range cxt.potentialDataTypes {
		if len(infos) > 1 {
			disambiguateDataTypes(cxt, infos)
		}
		for _, info := range infos {
			fieldMap, ok := fields[info.dataTypeCategory]
			if !ok {
				fieldMap = make(map[string]*potentialDataType)
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
			err := promoteDataType(top, suffix, f, idColumn)
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

func promoteDataType(top *ascii.Section, suffix string, dataTypeFields map[string]*potentialDataType, firstColumnType matter.TableColumn) error {
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
		table := findFirstTable(dt.section)
		if table == nil {
			continue
		}
		_, columnMap, _, err := findColumns(combineRows(table))
		if err != nil {
			return err
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
		ascii.Filter(dt.section, func(i interface{}) (remove bool, shortCircuit bool) {
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
	dataTypesSection := findSectionByType(top, matter.SectionDataTypes)
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

func stripDataTypeSuffixes(dataType string) string {
	for _, suffix := range matter.DataTypeSuffixes {
		if strings.HasSuffix(dataType, suffix) {
			dataType = dataType[0 : len(dataType)-len(suffix)]
			break
		}
	}
	return dataType
}

func disambiguateDataTypes(cxt *Context, infos []*potentialDataType) error {
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
			refParentId := strings.TrimSpace(stripReferenceSuffixes(getReferenceName(parentSection.Base)))
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
