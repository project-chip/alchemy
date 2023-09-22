package disco

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func organizeClassificationSection(doc *ascii.Doc, section *ascii.Section) {
	attributesTable := findFirstTable(section)
	if attributesTable != nil {
		organizeClassificationTable(doc, section, attributesTable)
	} else {
		fmt.Printf("No attributes table!")
	}
}

func organizeClassificationTable(doc *ascii.Doc, section *ascii.Section, attributesTable *types.Table) {
	rows := combineRows(attributesTable)

	headerRowIndex, columnMap, extraColumns := findColumns(rows, doc)

	if columnMap == nil {
		fmt.Println("can't rearrange classification table without header row")
		return
	}

	if len(columnMap) < 3 {
		fmt.Println("can't rearrange classification table with so few matches")
		return
	}

	renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.ClassificationTableColumnNames)

	reorderColumns(doc, section, rows, matter.ClassificationTableColumnOrder[:], columnMap, extraColumns)

}
