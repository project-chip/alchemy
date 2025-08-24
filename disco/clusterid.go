package disco

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
)

func (b *Baller) organizeClusterIDSection(cxt *discoContext) (err error) {
	for _, clusterIDs := range cxt.parsed.clusterIDs {
		clusterIDsTable := clusterIDs.table
		if clusterIDsTable == nil || clusterIDsTable.Element == nil {
			slog.Warn("Could not organize cluster ID section, as no table was found", log.Path("source", clusterIDs.section))
			return
		}
		if len(clusterIDsTable.Element.TableRows(asciidoc.RawReader)) > 2 {
			setSectionTitle(cxt.doc, clusterIDs.section, matter.ClusterIDsSectionName)
		} else {
			setSectionTitle(cxt.doc, clusterIDs.section, matter.ClusterIDSectionName)
		}

		if clusterIDsTable.ColumnMap == nil {
			return fmt.Errorf("can't rearrange cluster id table without header row in %s", cxt.doc.Path)
		}

		if len(clusterIDsTable.ColumnMap) < 2 {
			return fmt.Errorf("can't rearrange cluster id table with so few matches in %s", cxt.doc.Path)
		}

		err = b.renameTableHeaderCells(cxt, clusterIDs.section, clusterIDsTable, matter.Tables[matter.TableTypeClusterID].ColumnRenames)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in cluster ID table in %s: %w", cxt.doc.Path, err)
		}

		err = b.reorderColumns(cxt, clusterIDs.section, clusterIDsTable, matter.TableTypeClusterID)
		if err != nil {
			return err
		}
	}
	return
}
