package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func promoteDataTypes(top *ascii.Section, section *ascii.Section, sectionDataMap map[string]*sectionDataType) error {
	enumFields := make(map[string]*sectionDataType)
	bitmapFields := make(map[string]*sectionDataType)
	for field, info := range sectionDataMap {
		slog.Debug("\"field\": %s -> %s\n", field, info.dataType)
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
	dataTypesSection := findSectionByType(top, matter.SectionDataTypes)
	if dataTypesSection == nil {
		ts, _ := types.NewSection(top.Base.Level+1, []interface{}{"Data Types"})
		s, err := ascii.NewSection(ts)
		if err != nil {
			return err
		}
		dataTypesSection = s
		top.AppendSection(dataTypesSection)
	}
	for _, bitmap := range bitmapFields {
		if bitmap.section == nil {
			continue
		}
		t := findFirstTable(bitmap.section)
		if t == nil {
			continue
		}
		_, columnMap, _ := findColumns(t.Rows)
		_, ok := columnMap[matter.TableColumnBit]
		if !ok {
			continue
		}
		_, ok = columnMap[matter.TableColumnValue]
		if !ok {
			continue
		}
		bitmapSection, _ := types.NewSection(top.Base.Level+1, []interface{}{bitmap.name + "Bitmap"})
		bitmapSection.AddElement(t)
		s, err := ascii.NewSection(bitmapSection)
		if err != nil {
			return err
		}
		dataTypesSection.AppendSection(s)
		filter(bitmap.section, func(i interface{}) (remove bool, shortCircuit bool) {
			if i == bitmap.section {
				return true, true
			}
			return false, false
		})
		t.Attributes.Set("tableSectionName", section.Name)
	}
	return nil
}
