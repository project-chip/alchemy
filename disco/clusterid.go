package disco

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

func (b *Ball) organizeClusterIDSection(cxt *discoContext, dp *docParse) (err error) {
	for _, clusterIDs := range dp.clusterIDs {
		clusterIDsTable := clusterIDs.table
		if clusterIDsTable.element == nil {
			return fmt.Errorf("no cluster ID section found")
		}
		setSectionTitle(clusterIDs.section, matter.ClusterIDSectionName)

		if clusterIDsTable.columnMap == nil {
			return fmt.Errorf("can't rearrange cluster id table without header row in %s", dp.doc.Path)
		}

		if len(clusterIDsTable.columnMap) < 2 {
			return fmt.Errorf("can't rearrange cluster id table with so few matches in %s", dp.doc.Path)
		}

		err = b.renameTableHeaderCells(clusterIDsTable.rows, clusterIDsTable.headerRow, clusterIDsTable.columnMap, nil)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in cluster ID table in %s: %w", dp.doc.Path, err)
		}

		b.reorderColumns(dp.doc, clusterIDs.section, clusterIDsTable.rows, matter.ClusterIDTableColumnOrder[:], clusterIDsTable.columnMap, clusterIDsTable.extraColumns)
	}
	return
}

/*
func (b *Ball) organizeClusterIDSection(doc *ascii.Doc, section *ascii.Section) error {
	t := ascii.FindFirstTable(section)
	if t == nil {
		return fmt.Errorf("no cluster ID section found")
	}
	return b.organizeClusterIDTable(doc, section, t)
}

func (b *Ball) organizeClusterIDTable(doc *ascii.Doc, section *ascii.Section, attributesTable *types.Table) error {

	setSectionTitle(section, matter.ClusterIDSectionName)

	rows := ascii.TableRows(attributesTable)

	headerRowIndex, columnMap, extraColumns, err := ascii.MapTableColumns(doc, rows)
	if err != nil {
		return fmt.Errorf("failed mapping table columns for cluster ID table in section %s in %s: %w", section.Name, doc.Path, err)
	}

	if columnMap == nil {
		return fmt.Errorf("can't rearrange cluster id table without header row in %s", doc.Path)
	}

	if len(columnMap) < 2 {
		return fmt.Errorf("can't rearrange cluster id table with so few matches in %s", doc.Path)
	}

	err = b.renameTableHeaderCells(rows, headerRowIndex, columnMap, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in cluster ID table in %s: %w", doc.Path, err)
	}

	b.reorderColumns(doc, section, rows, matter.ClusterIDTableColumnOrder[:], columnMap, extraColumns)
	return nil
}
*/
