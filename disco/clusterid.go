package disco

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func organizeClusterIDSection(doc *ascii.Doc, section *ascii.Section) error {
	t := findFirstTable(section)
	if t == nil {
		return fmt.Errorf("no cluster ID section found")
	}
	return organizeClusterIDTable(doc, section, t)
}

func organizeClusterIDTable(doc *ascii.Doc, section *ascii.Section, attributesTable *types.Table) error {

	setSectionTitle(section, matter.ClusterIDSectionName)

	rows := combineRows(attributesTable)

	headerRowIndex, columnMap, extraColumns := findColumns(rows)

	if columnMap == nil {
		return fmt.Errorf("can't rearrange cluster id table without header row")
	}

	if len(columnMap) < 2 {
		return fmt.Errorf("can't rearrange cluster id table with so few matches")
	}

	err := renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.ClusterIDTableColumnNames)
	if err != nil {
		return err
	}

	reorderColumns(doc, section, rows, matter.ClusterIDTableColumnOrder[:], columnMap, extraColumns)
	return nil
}
