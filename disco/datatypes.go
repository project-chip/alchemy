package disco

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func promoteDataTypes(top *ascii.Section, section *ascii.Section, sectionDataMap map[string]*sectionDataType) error {
	//return nil
	enumFields := make(map[string]*sectionDataType)
	bitmapFields := make(map[string]*sectionDataType)
	for field, info := range sectionDataMap {
		fmt.Printf("\"field\": %s -> %s\n", field, info.dataType)
		switch info.dataType {
		case "enum8", "enum16", "enum32":
			enumFields[field] = info
		case "map8", "map16", "map32":
			bitmapFields[field] = info
		}
	}
	var dataTypeCount int
	find(section.Elements, func(el interface{}) bool {
		s, ok := el.(*ascii.Section)
		if !ok {
			return false
		}
		name := strings.TrimSpace(stripReferenceSuffixes(s.Name))
		fmt.Printf("name: %s -> \"%s\"\n", s.Name, strings.ToLower(name))
		if sdt, ok := enumFields[strings.ToLower(name)]; ok {
			fmt.Printf("Enum field! %s\n", name)
			sdt.section = s
			dataTypeCount++
		} else if sdt, ok := bitmapFields[strings.ToLower(name)]; ok {
			fmt.Printf("Bitmap field! %s\n", name)
			sdt.section = s
			dataTypeCount++
		}

		return false
	})
	if dataTypeCount == 0 {
		return nil
	}
	err := promoteDataType(top, section, "Bitmap", bitmapFields, matter.TableColumnBit)
	if err != nil {
		return err
	}
	err = promoteDataType(top, section, "Enum", enumFields, matter.TableColumnValue)
	if err != nil {
		return err
	}
	return nil
}

func promoteDataType(top *ascii.Section, section *ascii.Section, suffix string, dataTypeFields map[string]*sectionDataType, firstColumnType matter.TableColumn) error {
	var dataTypesSection *ascii.Section
	for _, dt := range dataTypeFields {
		if dt.section == nil {
			fmt.Printf("skipping non-section\n")
			continue
		}
		table := findFirstTable(dt.section)
		if table == nil {
			continue
		}
		_, columnMap, _ := findColumns(combineRows(table))
		if valueIndex, ok := columnMap[firstColumnType]; !ok || valueIndex > 0 {
			fmt.Printf("missing value column\n")
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
		filter(dt.section, func(i interface{}) (remove bool, shortCircuit bool) {
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
		if id, ok := table.Attributes.GetAsString(types.AttrID); ok {
			// Reuse the table's ID if it has one, so existing links get updated
			newAttr[types.AttrID] = id
		} else {
			newAttr[types.AttrID] = "ref_" + title.Content + ", " + title.Content
		}

		dataTypeSection.AddAttributes(newAttr)

		s, err := ascii.NewSection(dataTypesSection, dataTypeSection)

		if err != nil {
			return err
		}

		dataTypesSection.AppendSection(s)

		table.Attributes.Unset(types.AttrID)
		table.Attributes.Unset(types.AttrTitle)

		icr, _ := types.NewInternalCrossReference("ref_"+title.Content, "")
		err = setCellValue(dt.typeCell, []interface{}{icr})
		if err != nil {
			return err
		}
	}
	return nil
}

func promoteEnums(top *ascii.Section, section *ascii.Section, enumFields map[string]*sectionDataType) error {
	var dataTypesSection *ascii.Section
	for _, enum := range enumFields {
		if enum.section == nil {
			fmt.Printf("skipping non-section\n")
			continue
		}
		table := findFirstTable(enum.section)
		if table == nil {
			fmt.Printf("skipping no table\n")
			continue
		}
		fmt.Printf("data table %s row count %d\n", enum.name, len(table.Rows))
		_, columnMap, _ := findColumns(combineRows(table))
		for ct, c := range columnMap {
			fmt.Printf("%v -> %v\n", ct, c)
		}
		if valueIndex, ok := columnMap[matter.TableColumnValue]; !ok || valueIndex > 0 {
			fmt.Printf("missing value column\n")
			continue
		}

		title, err := types.NewStringElement(enum.name + "Enum")
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
		filter(enum.section, func(i interface{}) (remove bool, shortCircuit bool) {
			if t, ok := i.(*types.Table); ok && table == t {
				removedTable = true
				return true, true
			}
			return false, false
		})

		if !removedTable {
			return fmt.Errorf("unable to relocate enum value table")
		}

		enumSection, _ := types.NewSection(dataTypesSection.Base.Level+1, []interface{}{title})

		se, _ := types.NewStringElement(fmt.Sprintf("This data type is derived from %s", enum.dataType))
		p, _ := types.NewParagraph(nil, se)
		enumSection.AddElement(p)
		bl, _ := types.NewBlankLine()
		enumSection.AddElement(bl)
		enumSection.AddElement(table)
		enumSection.AddElement(bl)
		newAttr := make(types.Attributes)
		newAttr[types.AttrID] = "ref_" + title.Content + ", " + title.Content
		enumSection.AddAttributes(newAttr)

		s, err := ascii.NewSection(dataTypesSection, enumSection)

		if err != nil {
			return err
		}

		dataTypesSection.AppendSection(s)

		table.Attributes.Unset(types.AttrID)
		table.Attributes.Unset(types.AttrTitle)

		icr, _ := types.NewInternalCrossReference("ref_"+title.Content, "")
		err = setCellValue(enum.typeCell, []interface{}{icr})
		if err != nil {
			return err
		}
	}
	return nil
}

func promoteBitmaps(top *ascii.Section, section *ascii.Section, bitmapFields map[string]*sectionDataType) error {
	var dataTypesSection *ascii.Section
	for _, bitmap := range bitmapFields {
		if bitmap.section == nil {
			fmt.Printf("skipping non-section\n")
			continue
		}
		t := findFirstTable(bitmap.section)
		if t == nil {
			fmt.Printf("skipping no table\n")
			continue
		}
		fmt.Printf("data table row count %d\n", len(t.Rows))
		_, columnMap, _ := findColumns(combineRows(t))
		if bitIndex, ok := columnMap[matter.TableColumnBit]; !ok || bitIndex > 0 {
			fmt.Printf("missing bit column\n")
			continue
		}
		title, err := types.NewStringElement(bitmap.name + "Bitmap")
		if err != nil {
			return err
		}

		if dataTypesSection == nil {
			dataTypesSection, err = ensureDataTypesSection(top)
			if err != nil {
				return err
			}
		}

		bitmapSection, _ := types.NewSection(dataTypesSection.Base.Level+1, []interface{}{title})
		se, _ := types.NewStringElement(fmt.Sprintf("This data type is derived from %s", bitmap.dataType))
		p, _ := types.NewParagraph(nil, se)
		bitmapSection.AddElement(p)
		bl, _ := types.NewBlankLine()
		bitmapSection.AddElement(bl)
		bitmapSection.AddElement(t)
		bitmapSection.AddElement(bl)
		s, err := ascii.NewSection(dataTypesSection, bitmapSection)
		if err != nil {
			return err
		}

		dataTypesSection.AppendSection(s)
		filter(bitmap.section, func(i interface{}) (remove bool, shortCircuit bool) {
			if i == bitmap.section {
				fmt.Printf("removing existing section\n")
				return true, true
			}
			return false, false
		})
		t.Attributes.Set("tableSectionName", section.Name)
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
