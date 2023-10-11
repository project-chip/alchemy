package disco

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func (b *Ball) organizeClusterIDSection(doc *ascii.Doc, section *ascii.Section) error {
	t := parse.FindFirstTable(section)
	if t == nil {
		return fmt.Errorf("no cluster ID section found")
	}
	return b.organizeClusterIDTable(doc, section, t)
}

func (b *Ball) organizeClusterIDTable(doc *ascii.Doc, section *ascii.Section, attributesTable *types.Table) error {

	setSectionTitle(section, matter.ClusterIDSectionName)

	rows := parse.TableRows(attributesTable)

	headerRowIndex, columnMap, extraColumns, err := parse.MapTableColumns(rows)
	if err != nil {
		return err
	}

	if columnMap == nil {
		return fmt.Errorf("can't rearrange cluster id table without header row")
	}

	if len(columnMap) < 2 {
		return fmt.Errorf("can't rearrange cluster id table with so few matches")
	}

	err = renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.ClusterIDTableColumnNames)
	if err != nil {
		return err
	}

	reorderColumns(doc, section, rows, matter.ClusterIDTableColumnOrder[:], columnMap, extraColumns)
	return nil
}
