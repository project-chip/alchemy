package disco

import (
	"fmt"
	"log/slog"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

func (b *Ball) organizeClassificationSection(doc *ascii.Doc, section *ascii.Section) error {
	attributesTable := ascii.FindFirstTable(section)
	if attributesTable == nil {
		return fmt.Errorf("no classification table found")
	}
	return b.organizeClassificationTable(doc, section, attributesTable)
}

func (b *Ball) organizeClassificationTable(doc *ascii.Doc, section *ascii.Section, attributesTable *types.Table) error {
	rows := ascii.TableRows(attributesTable)

	headerRowIndex, columnMap, extraColumns, err := ascii.MapTableColumns(doc, rows)
	if err != nil {
		return fmt.Errorf("failed mapping table columns for classification table in section %s: %w", section.Name, err)
	}

	if columnMap == nil {
		slog.Debug("can't rearrange classification table without header row")
		return nil
	}

	if len(columnMap) < 3 {
		slog.Debug("can't rearrange classification table with so few matches")
		return nil
	}

	err = b.renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.ClassificationTableColumnNames)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in section %s in %s: %w", section.Name, doc.Path, err)
	}

	var order []matter.TableColumn

	docType, err := doc.DocType()
	if err != nil {
		return fmt.Errorf("error getting doc type in section %s in %s: %w", section.Name, doc.Path, err)
	}
	switch docType {
	case matter.DocTypeCluster:
		order = matter.AppClusterClassificationTableColumnOrder[:]
	case matter.DocTypeDeviceType:
		order = matter.DeviceTypeClassificationTableColumnOrder[:]
	}

	if len(order) > 0 {
		b.reorderColumns(doc, section, rows, order, columnMap, extraColumns)
	}
	return nil
}
