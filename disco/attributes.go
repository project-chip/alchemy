package disco

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func organizeAttributesSection(doc *ascii.Doc, section *ascii.Section) {
	attributesTable := findFirstTable(section)
	if attributesTable != nil {
		organizeAttributesTable(doc, section, attributesTable)
	} else {
		fmt.Printf("No attributes table!")
	}
}

func organizeAttributesTable(doc *ascii.Doc, section *ascii.Section, attributesTable *types.Table) {
	rows := combineRows(attributesTable)

	_, columnMap, extraColumns := findColumns(rows, doc)

	if columnMap == nil {
		fmt.Println("can't rearrange attributes table without header row")
		return
	}

	if len(columnMap) < 5 {
		fmt.Println("can't rearrange attributes table with so few matches")
		return
	}

	fixAttributeAccess(doc, rows, columnMap)

	reorderColumns(doc, section, rows, matter.AttributesTableColumnOrder[:], columnMap, extraColumns)

}

func getAttributeNames(doc *ascii.Doc, rows []*types.TableRow, columnMap map[matter.TableColumn]int) (names []string) {
	nameMap := make(map[string]bool)
	if nameIndex, ok := columnMap[matter.TableColumnName]; ok {
		for _, row := range rows {
			val := strings.TrimSpace(getCellValue(row.Cells[nameIndex]))
			if _, ok := nameMap[val]; !ok {
				names = append(names, val)
			}
		}
	}
	return names
}

var accessPattern = regexp.MustCompile(`^(?:(?P<ColumnSpan>[0-9]+)?(?:\.(?P<RowSpan>[0-9]+))?\+)?(?:(?P<Duplication>[0-9]+)\*)?(?P<HorizontalAlignment>[<>^])?(?P<VerticalAlignment>\.[<>^])?(?P<Style>[adehlms])?$`)

func fixAttributeAccess(doc *ascii.Doc, rows []*types.TableRow, columnMap map[matter.TableColumn]int) {
	accessIndex, ok := columnMap[matter.TableColumnAccess]
	if !ok {
		return
	}
	for _, row := range rows {
		cell := row.Cells[accessIndex]
		if len(cell.Elements) == 0 {
			continue
		}
		p, ok := cell.Elements[0].(*types.Paragraph)
		if !ok || len(p.Elements) == 0 {
			continue
		}
		s, ok := p.Elements[0].(*types.StringElement)
		if !ok && s != nil {
			continue
		}
		//fmt.Printf("Access: %s\n", s.Content)
	}
}
