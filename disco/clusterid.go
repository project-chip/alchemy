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
		if len(clusterIDsTable.element.TableRows) > 2 {
			setSectionTitle(clusterIDs.section, matter.ClusterIDsSectionName)
		} else {
			setSectionTitle(clusterIDs.section, matter.ClusterIDSectionName)
		}

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

		err = b.reorderColumns(dp.doc, clusterIDs.section, &clusterIDsTable, matter.TableTypeClusterID)
		if err != nil {
			return err
		}
	}
	return
}
