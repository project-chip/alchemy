package disco

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func organizeClusterIDSection(doc *ascii.Doc, section *ascii.Section) {
	t := findFirstTable(section)
	if t != nil {
		organizeClusterIDTable(doc, section, t)
	}
}

func organizeClusterIDTable(doc *ascii.Doc, section *ascii.Section, attributesTable *types.Table) {

	setSectionTitle(section, matter.ClusterIDSectionName)

	rows := combineRows(attributesTable)

	headerRowIndex, columnMap, extraColumns := findColumns(rows)

	if columnMap == nil {
		fmt.Println("can't rearrange cluster id table without header row")
		return
	}

	if len(columnMap) < 2 {
		fmt.Println("can't rearrange cluster id table with so few matches")
		return
	}

	renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.ClusterIDTableColumnNames)

	reorderColumns(doc, section, rows, matter.ClusterIDTableColumnOrder[:], columnMap, extraColumns)
}
