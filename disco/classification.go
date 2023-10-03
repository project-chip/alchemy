package disco

import (
	"fmt"
	"log/slog"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func (b *Ball) organizeClassificationSection(doc *ascii.Doc, section *ascii.Section) error {
	attributesTable := findFirstTable(section)
	if attributesTable == nil {
		return fmt.Errorf("no attributes table found")
	}
	b.organizeClassificationTable(doc, section, attributesTable)
	return nil
}

func (b *Ball) organizeClassificationTable(doc *ascii.Doc, section *ascii.Section, attributesTable *types.Table) error {
	rows := combineRows(attributesTable)

	headerRowIndex, columnMap, extraColumns, err := findColumns(rows)
	if err != nil {
		return err
	}

	if columnMap == nil {
		slog.Debug("can't rearrange classification table without header row")
		return nil
	}

	if len(columnMap) < 3 {
		slog.Debug("can't rearrange classification table with so few matches")
		return nil
	}

	err = renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.ClassificationTableColumnNames)
	if err != nil {
		return err
	}

	var order []matter.TableColumn

	docType, err := doc.DocType()
	if err != nil {
		return err
	}
	switch docType {
	case matter.DocTypeAppCluster:
		order = matter.AppClusterClassificationTableColumnOrder[:]
	case matter.DocTypeDeviceType:
		order = matter.DeviceTypeClassificationTableColumnOrder[:]
	}

	if len(order) > 0 {
		reorderColumns(doc, section, rows, order, columnMap, extraColumns)
	}
	return nil
}
